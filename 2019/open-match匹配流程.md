# open-match匹配流程

(金庆的专栏 2019.1)

https://github.com/GoogleCloudPlatform/open-match

open-match 是一个通用的游戏匹配框架。
由游戏提供自定义的匹配算法（以docker镜像的方式提供）。

分为多个进程，各进程之间共享一个 redis.

* 前端, 接收玩家加入 redis，成功后通知玩家房间服地址
* 后端，设置一局游戏的匹配规则，设置房间服地址
* MMFOrc，启动匹配算法(MMF)
* MMF, 自定义匹配算法，读取 redis 获取玩家，匹配成功就将结果写入 redis. 仅匹配一局就退出。

游戏服中连接 open-match 的前端与后端的进程，分别称为 frontendclient 和 Director。
输入分2部份，一是玩家信息，二是对局信息。
Director 向后端输入对局信息，就会收到一个接一个的对局人员列表.
Director 需要为每个对局开房间，然后通知后端房间地址。
后端将房间地址写入 redis, 然后前端读取到房间地址，就通知 frontendclient，让玩家进入房间。

## test/cmd/frontendclient

模拟大厅服或组队服，连接前端API, 请求匹配玩家/队伍。成功后将收到房间服(DGS)的地址(Assignment)。

Player 实际上是一个队伍，其中ID字段是用空格分隔的多个ID. 
虽然参数类型都是 Player, CreatePlayer() 参数为整个队伍，而 GetUpdates() 参数是单个玩家。

main() 中创建多个玩家，每个玩家调用 GetUpdates() 以获取结果，go waitForResults() 中处理结果。
waitForResult() 读取流中的匹配结果，压入 resultsChan（但好像 resultsChan 仅用于打印）。
所有玩家合并到 g 实例中，然后调用 CreatePlayer() 请求匹配。

cleanup() 调用 DeletePlayer() 来删除匹配请求，不仅需删除整个队伍，也需要删除单个玩家。

好像最后取结果没取对地方，应该从 resultChan 中获取 Assignment, 并用该地址 udpClient().

看了该示例就可以理解 frontend.proto

## examples/backendclient

MatchObject.Properties 是从 testprofile.json 读取的，应该改名为 Profile 是否更好点？
pbProfile 是 MatchObject，Profile 等同于 MatchObject?
Profile 的定义是 MMF 所需的所有参数。
`pbProfile.Properties = jsonProfile` 重复了2遍。

ListMatches()列出这个Profile的所有匹配。
收到一个匹配后，须用CreateAssignments()将房间服地址, 称为 Assignment, 发送到所有游戏客户端。

## cmd/frontendapi

CreatePlayer() 将 Player 对象写入 redis, 键值为 Player.Id, 类型为 HSET。
对 Player 的每个 attribute，添加到 ZSET 中去。
此处 Player 是一组玩家。

GetUpdates() 每隔2s读取redis, Player数据有变化时就发送。此处 Player 是单个玩家。

如果CreatePlayer()中队伍只有一个玩家，
则写入的Player与GetUpdates()中读取的玩家是同一个redis键。

## cmd/backendapi

CreateMatch() 中 profile 类型为 MatchObject, 是一个比赛的限制条件。
profile 先写入 redis, 键为 profile.Id.
`requestKey := xid() + "." + profile.Id`,
并将 requestKey 加入 redis 集合 "profileq"。
然后每2s查询 redis, 看是否有 requestKey 键出现，并返回该值。

ListMatch() 每2s调用一次 CreateMatch().

DeleteMatch() 仅仅删除 Id 这个键。

CreateAssignments() 为多个队伍设置Assignment, 即房间地址。
遍历所有Roster中的Player对象，在redis中设置Assignment.
(Assignment 更改后，会触发前端更新。)
将所有 Player.Id 从 "proposed" 移到 "deindexed"，这两个是 ZSET, 分值为加入时间。
Roster 应该是比赛中的阵营，如红方，蓝方，每个阵营中可有多个队伍。

DeleteAssignments() 仅仅遍历所有 Player 对象来删除 Assignment 字段。

## cmd/mmforc

匹配流程是由 mmforc (matchmaking function orchestrator) 控制的。

mmforc 每秒从 redis 的 profileq 中取出 100 个成员, 其中 profileq 是个set类型，
使用命令为`SPOP profileq 100`.

对每个 profile, 创建一个 k8s 任务：

```
    // Kick off the job asynchrnously
    go mmfunc(ctx, profile, cfg, defaultMmfImages, clientset, &pool)
```

每隔10s, 还有所有匹配任务都完成后，需要 `checkProposals`, 即创建 evaluator 任务。

profileq 中的元素 profile 为字符串，matchObjectID.profileID。
以 profileID 为键，可以从 redis 读取 profile 的内容, profile 是个 MatchObject 对象。

profile 的内容为 json 串，其中 "jsonkeys.mmfImages" 为 mmf (matchmaking function) 镜像。

如果profile读取失败，或者 mmfImages 为空，则使用默认的镜像。mmfImages 未来会支持多个镜像。

通过 MMF_* 环境变量传入各种参数.

## mmf

示例：examples\functions\golang\manual-simple

从环境变量 "MMF_PROFILE_ID" 解析出 profileID, 并向 redis 查询(HGETALL) profile，HSET 类型。

从 profile 中取 pools 字段，即匹配条件。
pools 分为多个 pool, 每个 pool 中有多个 filter, 每个 filter 向 redis 取符合的 Player.

profile 用到以下字段：

* "properties.playerPool"
  json串，是一些过滤条件，如“mmr: 100-999”
* "properties.roster"
  json串, 是多个队伍大小，如 “red: 4”

示例见：`examples\backendclient\profiles\testprofile.json`

### 简单匹配过程

simple mmf 的匹配过程如下：

1. 从 redis 查询 profile，获取过滤条件和各队伍大小
1. 每个过滤条件向 redis 查询，所有结果的交集为可选成员
1. 去除 ignoreList, 即最近 800s 内已匹配成功的成员，即 proposal 和 deindexed ZSET 列表。
1. 如果可选成员个数太小，则 insufficient_players 并退出
1. 分配各个队伍成员
1. 向 redis 记录结果

### 结果

profile 中添加 roster，即各阵营成员名单，存入 prososalKey.
保存不分队伍的成员名单。
然后向 "proposalq" 添加 prososalKey

### 细节

poolRosters 以 (pool名, filter attribute) 为键，值为 Player ID 列表. 
保存从 redis 查询的符合条件的 Player ID.

overlaps 以 pool 名为键，保存符合该pool中所有filter的 Player ID 列表，去除 ignore list.

rosters 是 profile 中的 "properties.rosters" 字段。不知何用？
遍历 rosters, 为每个阵营的每个player找到对应pool的PlayerID, 保存到 mo.Rosters.
其中 profileRosters 好像没用。



>reference
[cppblog](http://www.cppblog.com/jinq0123/archive/2019/01/31/216228.html)
