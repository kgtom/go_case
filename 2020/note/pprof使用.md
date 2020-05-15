## 本节大纲
* [一、概括](#0)
* [二、排查profileCPU占用的采样情况](#1)
* [三、排查heap内存分配的采样情况](#2)
* [四、排查allocs频繁分配内存及回收GC情况](#3)
* [五、排查goroutine协程泄露采样情况](#4)
* [六、排查mutex锁争用的采样情况](#5)
* [七、排查blocks堵塞操作的采样情况](#6)
* [八、trace的使用](#7)

## <span id="0">一、概括</span>
* 常用命令：topN 、list 、web、pdf 等
* 使用goroutine很简单，重要合理管理G的生命周期，避免无法结束G,资源无法回收，造成内存泄露
* 内存泄露：debug/pprof/heap可以查看内存分配情况、哪行代码占用多少内存及其代码位置，但不能直接说明有内存泄露，但如果一个代码位置占用内存持续增长，则基本可以确定存在内存泄露。查找内存问题：可以使用 两次 heap的diff
* top命令，重点关注 RES占用物理内存大小，可以比较不同时间点的内存占用比
* goroutine泄露：可以采用两次diff比较或者发现一段时间G数量只增不减，需要重点关注

## <span id="1">二、排查CPU占用情况</span>

### 背景
通过top命令，发现某个进程占用cpu过高，一般情况下死循环导致。
### 常用命令
 * go tool pprof http://xxx/debug/pprof/profile?second=30: 等待30s后，查看或导出结果(默认30s,可以自定义时长)
 * top：查看cpu较高的调用
 * text、pdf ：导出可视化格式 text或者pdf，前提需要安装 graphviz(brew install graphviz)
 * list+func：查看在代码的哪个位置及每行代码的耗时
 * web :生产.svg文件，用浏览器打开查看,方框特别大，箭头特别粗，占cpu大，重点看
 
 ### 格式
 ~~~
 go tool pprof 最简单的使用方式为 go tool pprof [binary] [source]
 ~~~
 * binary 是应用程序的二进制文件，用来解析各种符号；
 * source 表示 profile 数据的来源，locahost本地的文件或者 http文件
 * 可以用本地二进制文件，解析远程profile文件
 ~~~
 (base) DESKTOP-HBQDAKA :: ~ » go tool pprof -http=":8005" devspace/pontus/src/apps/front/front  /Users/tom/pprof/pprof.front.samples.cpu.002.pb.gz
 ~~~

### 运行
~~~

go tool pprof http://localhost:8001/debug/pprof/profile\?seconds\=30                                                                                             1 ↵
Fetching profile over HTTP from http://localhost:8001/debug/pprof/profile?seconds=30
Saved profile in /Users/tom/pprof/pprof.front.samples.cpu.001.pb.gz
File: front
Type: cpu
Time: May 12, 2020 at 5:59pm (CST)
Duration: 30s, Total samples = 20ms (0.067%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) text
Showing nodes accounting for 20ms, 100% of 20ms total
Showing top 10 nodes out of 32
      flat  flat%   sum%        cum   cum%
      10ms 50.00% 50.00%       10ms 50.00%  encoding/json.indirect
      10ms 50.00%   100%       10ms 50.00%  runtime.memclrNoHeapPointers
         0     0%   100%       10ms 50.00%  clients/modelproxy.(*cliImpl).CommonClassify
         0     0%   100%       10ms 50.00%  common/proto.RegisterPontusHandler.func2
         0     0%   100%       10ms 50.00%  common/proto.request_Pontus_InvokeClassifyModel_0
         0     0%   100%       10ms 50.00%  encoding/json.(*decodeState).literalStore
         0     0%   100%       10ms 50.00%  encoding/json.(*decodeState).object
         0     0%   100%       10ms 50.00%  encoding/json.(*decodeState).unmarshal
         0     0%   100%       10ms 50.00%  encoding/json.(*decodeState).value
         0     0%   100%       10ms 50.00%  encoding/json.Unmarshal
(pprof) pdf
Generating report in profile001.pdf
(pprof) list
command list requires an argument
(pprof) list CommonClassify
Total: 20ms
ROUTINE ======================== clients/modelproxy.(*cliImpl).CommonClassify in /dworkspace/src/clients/modelproxy/modelproxy.go
         0       10ms (flat, cum) 50.00% of Total
 Error: open /dworkspace/src/clients/modelproxy/modelproxy.go: no such file or directory
~~~

### 可视化火焰图
#### 命令格式：
~~~
$ go tool pprof -http=":8005" [binary] [profile]
~~~
* binary：可执行应用程序的二进制文件，由go build生成
* profile：pb格式的文件，由go tool pprof 采样生成

~~~
go tool pprof http://localhost:8001/debug/pprof/profile\?seconds\=30                                                                                             1 ↵
Fetching profile over HTTP from http://localhost:8001/debug/pprof/profile?seconds=30
Saved profile in /Users/tom/pprof/pprof.front.samples.cpu.001.pb.gz
~~~

#### 运行可视化火焰图

* 使用8005端口开启，其中front是二进制文件
~~~

(base) DESKTOP-HBQDAKA :: src/apps/front ‹master*› » go tool pprof -http=":8005" front /Users/tom/pprof/pprof.samples.cpu.001.pb.gz
~~~

* 自动打开浏览器：localhost:8005/ui
* 查看火焰图 ：http://localhost:8005/ui/flamegraph
* 火焰图的x轴占用 CPU 使用的长短，y轴代表调用链，顶层正在执行的，下方都是它的父函数。观察业务函数在火焰图中的长(宽)度，如果是占据较长(宽)，说明可能存在性能问题,也要留意顶层占较宽的(顶层多数是库函数)另外火焰图的配色没有含义，像火一样。

## <span id="2">三、排查内存使用情况</span>
~~~
(base) tomdeMacBook-Pro :: ~ » go tool pprof http://10.1.1.10:8001/debug/pprof/heap
Fetching profile over HTTP from http://10.1.1.10:8001/debug/pprof/heap
Saved profile in /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: front
Type: inuse_space
Time: May 15, 2020 at 6:03pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1044.70kB, 100% of 1044.70kB total
      flat  flat%   sum%        cum   cum%
  532.26kB 50.95% 50.95%   532.26kB 50.95%  vendor/github.com/gogo/protobuf/proto.RegisterType
  512.44kB 49.05%   100%   512.44kB 49.05%  vendor/github.com/go-redis/redis/internal/pool.NewConnPool
         0     0%   100%   512.44kB 49.05%  dao/lock.init.0
         0     0%   100%  1044.70kB   100%  runtime.main
         0     0%   100%   512.44kB 49.05%  vendor/github.com/go-redis/redis.NewClient
         0     0%   100%   512.44kB 49.05%  vendor/github.com/go-redis/redis.newConnPool
         0     0%   100%   532.26kB 50.95%  vendor/k8s.io/api/batch/v1beta1.init.0
(pprof) web
(pprof) exit
(base) tomdeMacBook-Pro :: ~ » go tool pprof http://10.1.1.10:8001/debug/pprof/heap
Fetching profile over HTTP from http://10.1.1.10:8001/debug/pprof/heap
Saved profile in /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
File: front
Type: inuse_space
Time: May 15, 2020 at 6:04pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 6295.79kB, 100% of 6295.79kB total
Showing top 10 nodes out of 36
      flat  flat%   sum%        cum   cum%
 4737.09kB 75.24% 75.24%  4737.09kB 75.24%  vendor/xxx.op.xxx.com/xxx-rd/rollingwriter.glob..func1
  532.26kB  8.45% 83.70%   532.26kB  8.45%  vendor/github.com/gogo/protobuf/proto.RegisterType
     514kB  8.16% 91.86%      514kB  8.16%  bufio.NewReaderSize
  512.44kB  8.14%   100%   512.44kB  8.14%  vendor/github.com/go-redis/redis/internal/pool.NewConnPool
         0     0%   100%      514kB  8.16%  bufio.NewReader
         0     0%   100%  1184.27kB 18.81%  clients/kre.(*cliImpl).InferImgAudit
         0     0%   100%  1184.27kB 18.81%  clients/modelproxy.(*cliImpl).CommonClassify
         0     0%   100%  1184.27kB 18.81%  clients/modelproxy.(*cliImpl).CommonClassify.func1
         0     0%   100%  1184.27kB 18.81%  common/proto._Pontus_InvokeClassifyModel_Handler
         0     0%   100%   512.44kB  8.14%  dao/lock.init.0
(pprof) web
(pprof) exit
(base) tomdeMacBook-Pro :: ~ » go tool pprof -base /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
File: front
Type: inuse_space
Time: May 15, 2020 at 6:03pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 5251.10kB, 100% of 5251.10kB total
Showing top 10 nodes out of 29
      flat  flat%   sum%        cum   cum%
 4737.09kB 90.21% 90.21%  4737.09kB 90.21%  vendor/xxx.op.xxx.com/xxx-rd/rollingwriter.glob..func1
     514kB  9.79%   100%      514kB  9.79%  bufio.NewReaderSize
         0     0%   100%      514kB  9.79%  bufio.NewReader
         0     0%   100%  1184.27kB 22.55%  clients/kre.(*cliImpl).InferImgAudit
         0     0%   100%  1184.27kB 22.55%  clients/modelproxy.(*cliImpl).CommonClassify
         0     0%   100%  1184.27kB 22.55%  clients/modelproxy.(*cliImpl).CommonClassify.func1
         0     0%   100%  1184.27kB 22.55%  common/proto._Pontus_InvokeClassifyModel_Handler
         0     0%   100%      514kB  9.79%  net/http.(*conn).serve
         0     0%   100%      514kB  9.79%  net/http.newBufioReader
         0     0%   100%  2368.55kB 45.11%  service/classifymodel.(*classifyModel).InvokeClassifyModel.func1
(pprof) list .InvokeClassifyModel.func1
Total: 5.13MB
ROUTINE ======================== service/classifymodel.(*classifyModel).InvokeClassifyModel.func1 in /dworkspace/src/service/classifymodel/facade.go
         0     2.31MB (flat, cum) 45.11% of Total
 Error: open /dworkspace/src/service/classifymodel/facade.go: no such file or directory
(pprof)
(pprof) exit
(base) DESKTOP-HBQDAKA :: ~ » go tool pprof -http=":8005" devspace/pontus/src/apps/front/front  /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz

~~~

另外：list+func ，可以查看本地代码，但查看不到远程代码，两种方式查看：
* 1.没有浏览器的，使用wget http://10.1.1.10:8001/debug/pprof/heap?debug=1 ,下载下来，根据方法名称查询
* 2.有浏览器的，直接查看 http://10.1.1.10:8001/debug/pprof/heap?debug=1 页面


## <span id="3">四、排查allocs频繁分配内存及回收GC情况</span>


## <span id="4">五、排查goroutine协程泄露采样情况</span>
* 查看goroutine的数量：打开http://10.1.1.10:8001/debug/pprof
* 查看当前所有运行的 goroutines 堆栈跟踪,格式：go tool pprof http://10.1.1.10:8001/debug/pprof/goroutine
* 使用 web、list 等查看
* 查看G总数量，及堵塞在某行代码的G的数量

 http://10.1.1.10:8001/debug/pprof/goroutine?debug=1
 ~~~
 goroutine profile: total 45  
6 @ 0x4316bf 0x441198 0x87c4ad 0x45eb91
#	0x87c4ac	vendor/google.golang.org/grpc.(*ccBalancerWrapper).watcher+0x12c	/dworkspace/src/vendor/google.golang.org/grpc/balancer_conn_wrappers.go:115

 ~~~
 - 45 G的总数量
 - 6个G 在代码115行堵塞
 
* 查看G的详细信息，包括G的编号、G的状态、持续时间
~~~
goroutine 1 [chan receive, 229 minutes]:
vendor/xxx.op.xxx.com/xxx-rd/doraemon/server.(*Server).Run(0xc000030000)
	/dworkspace/src/vendor/xxx.op.xxx.com/xxx-rd/xxx/server/server.go:170 +0xbb
common/server.Run(0x1f40)
	/dworkspace/src/common/server/server.go:32 +0x9f
main.main()
	/dworkspace/src/apps/front/main.go:66 +0xc9
~~~
- 编号1的G 具体在代码main.go 66行，堵塞了229分钟了，因为这个是监听客户端，所以从上线后一直堵塞，如果不堵塞了，说明服务挂了。
#### 格式
~~~
(base) tomdeMacBook-Pro :: ~ »  go tool pprof http://10.1.1.10:8001/debug/pprof/goroutine
Fetching profile over HTTP from http://10.1.1.10:8001/debug/pprof/goroutine
Saved profile in /Users/tom/pprof/pprof.front.goroutine.007.pb.gz
File: front
Type: goroutine
Time: May 14, 2020 at 8:01pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 44, 100% of 44 total
Showing top 10 nodes out of 74
      flat  flat%   sum%        cum   cum%
        42 95.45% 95.45%         42 95.45%  runtime.gopark
         1  2.27% 97.73%          1  2.27%  runtime.notetsleepg
         1  2.27%   100%          1  2.27%  runtime/pprof.writeRuntimeProfile
         0     0%   100%          7 15.91%  bufio.(*Reader).Read
         0     0%   100%          1  2.27%  bytes.(*Buffer).ReadFrom
         0     0%   100%          1  2.27%  common/proto.RegisterPontusHandlerFromEndpoint.func1.1
         0     0%   100%          1  2.27%  common/server.Run
         0     0%   100%          1  2.27%  crypto/tls.(*Conn).Read
         0     0%   100%          1  2.27%  crypto/tls.(*Conn).readFromUntil
         0     0%   100%          1  2.27%  crypto/tls.(*Conn).readRecord
(pprof)

~~~



## <span id="5">六、排查mutex锁争用的采样情况</span>
* 获取导致 mutex 争用的 goroutine 堆栈，使用前需要先设置采样大小 runtime.SetMutexProfileFraction(1)
~~~
(base) tomdeMacBook-Pro :: ~ » go tool pprof http://10.1.1.10:8001/debug/pprof/mutex
~~~

## <span id="6">七、排查blocks堵塞操作的采样情况</span>
* 获取导致阻塞的 goroutine 堆栈(如 channel, mutex 等)，使用前需要先设置采样大小 runtime.SetBlockProfileRate(1)
~~~
go tool pprof http://10.1.1.10:8001/debug/pprof/block

~~~
## <span id="7">八、trace的使用</span>

### 格式

* 生产trace.out 文件
* 打开trace.out 文件，默认启动随机的http端口
* 查看浏览器打开的页面
~~~
(base) DESKTOP-HBQDAKA :: ~ » curl http://10.1.1.10:8001/debug/pprof/trace?seconds=10 > trace.out
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 24.9M    0 24.9M    0     0  2532k      0 --:--:--  0:00:10 --:--:-- 3069k
(base) DESKTOP-HBQDAKA :: ~ »
(base) DESKTOP-HBQDAKA :: ~ »
(base) DESKTOP-HBQDAKA :: ~ » ls
Applications    Documents       Library         Music           Postman         bak             flamegraph.html miniconda3      pprof           projects        trace.out
Desktop         Downloads       Movies          Pictures        Public          devspace        go              pip.conf        profile001.pdf  tools           zdevspace_bak
(base) DESKTOP-HBQDAKA :: ~ » go tool trace trace.out
2020/05/13 11:17:11 Parsing trace...
2020/05/13 11:17:20 Splitting trace...
2020/05/13 11:17:33 Opening browser. Trace viewer is listening on http://127.0.0.1:52116

~~~

* 浏览器打开页面(ps chrome 80版本以下可以正常打开view trace)
~~~
View trace
Goroutine analysis
Network blocking profile (⬇)
Synchronization blocking profile (⬇)
Syscall blocking profile (⬇)
Scheduler latency profile (⬇)
User-defined tasks
User-defined regions
Minimum mutator utilization
~~~



