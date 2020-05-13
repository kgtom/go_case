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
常用命令：topN 、list 、web、pdf 等

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
* 火焰图的x轴占用 CPU 使用的长短，y轴代表调用链，顶层正在执行的，下方都是它的父函数，顶层如果是占据较宽，说明可能存在性能问题。另外火焰图的配色没有含义，像火一样。


## <span id="2">三、排查内存使用情况</span>

* 使用命令topN 、list 、web查看
~~~
---1.采样远程服务，生产pb：(base) DESKTOP-HBQDAKA :: ~ » go tool pprof http://10.1.1.10:8001/debug/pprof/heap\?seconds\=60
---2.使用top 、list 、web 查看
---3.使用本地二进制文件，运行远程pb,生成火焰图：(base) DESKTOP-HBQDAKA :: ~ » go tool pprof -http=":8005" devspace/pontus/src/apps/front/front  /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz

(base) DESKTOP-HBQDAKA :: ~ » go tool pprof http://10.1.1.10:8001/debug/pprof/heap\?seconds\=60
Fetching profile over HTTP from http://10.1.1.10:8001/debug/pprof/heap?seconds=60
Saved profile in /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
File: front
Type: inuse_space
Time: May 13, 2020 at 9:21pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
(pprof)
(pprof) top
Showing nodes accounting for 18343.99kB, 100% of 18343.99kB total
Showing top 10 nodes out of 68
      flat  flat%   sum%        cum   cum%
14211.29kB 77.47% 77.47% 14211.29kB 77.47%  
vendor/git.op.xxx.com/xxx-rd/rollingwriter.glob..func1
 1024.38kB  5.58% 83.06%  1024.38kB  5.58%  runtime.malg
  544.67kB  2.97% 86.02%   544.67kB  2.97%  vendor/google.golang.org/grpc/internal/transport.newBufWriter
     514kB  2.80% 88.83%      514kB  2.80%  bufio.NewWriterSize
  513.31kB  2.80% 91.62%   513.31kB  2.80%  vendor/golang.org/x/net/http2/hpack.(*headerFieldTable).addEntry
  512.25kB  2.79% 94.42%   512.25kB  2.79%  vendor/github.com/golang/protobuf/proto.(*tagMap).put
  512.05kB  2.79% 97.21%  1536.42kB  8.38%  runtime.systemstack
  512.05kB  2.79%   100%   512.05kB  2.79%  vendor/google.golang.org/grpc/internal/transport.(*Stream).waitOnHeader
         0     0%   100%   512.25kB  2.79%  clients/kre.(*cliImpl).InferImgAudit
         0     0%   100%   512.25kB  2.79%  clients/kre/proto.(*InferImgAuditReq).String

(pprof) list .InferImgAudit
Total: 3.03MB
ROUTINE ======================== clients/kre.(*cliImpl).InferImgAudit in /dworkspace/src/clients/kre/kre.go
         0   512.25kB (flat, cum) 16.49% of Total
 Error: open /dworkspace/src/clients/kre/kre.go: no such file or directory
ROUTINE ======================== clients/kre/proto.(*InferImgAuditReq).String in /dworkspace/src/clients/kre/proto/kre.pb.go
         0   512.25kB (flat, cum) 16.49% of Total
 Error: open /dworkspace/src/clients/kre/proto/kre.pb.go: no such file or directory
(pprof) web
(pprof) exit
(base) DESKTOP-HBQDAKA :: ~ » go tool pprof -http=":8005" devspace/pontus/src/apps/front/front  /Users/tom/pprof/pprof.front.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz

~~~


## <span id="3">四、排查allocs频繁分配内存及回收GC情况</span>


## <span id="4">五、排查goroutine协程泄露采样情况</span>


## <span id="5">六、排查mutex锁争用的采样情况</span>


## <span id="6">七、排查blocks堵塞操作的采样情况</span>

## <span id="7">八、trace的使用</span>

### 格式

* 生产trace.out 文件
* 打开trace.out 文件，默认启动随机的http端口
* 查看浏览器打开的页面
~~~
(base) DESKTOP-HBQDAKA :: ~ » curl localhost:8001/debug/pprof/trace?seconds=10 > trace.out
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

* 浏览器打开页面
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
