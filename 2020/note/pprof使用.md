## 本节大纲
* [一、概括](#0)
* [二、排查profileCPU占用的采样情况](#1)
* [三、排查heap内存分配的采样情况](#2)
* [四、排查allocs频繁分配内存及回收GC情况](#3)
* [五、排查goroutine协程泄露采样情况](#4)
* [六、排查mutex锁争用的采样情况](#5)
* [七、排查blocks堵塞操作的采样情况](#6)

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
 * web :生产.svg文件，用浏览器打开查看
 
 ### 格式
 ~~~
 go tool pprof 最简单的使用方式为 go tool pprof [binary] [source]
 ~~~
 * binary 是应用程序的二进制文件，用来解析各种符号；
 * source 表示 profile 数据的来源，locahost本地的文件或者 http文件
 * 可以用本地二进制文件，解析远程profile文件

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


## <span id="2">三、排查内存使用情况</span>

### 格式

~~~
go tool pprof http://localhost:8001/debug/pprof/heap
~~~

### 运行

~~~

~~~
* 使用topN 、list 、web查看

## <span id="3">四、排查allocs频繁分配内存及回收GC情况</span>


## <span id="4">五、排查goroutine协程泄露采样情况</span>


## <span id="5">六、排查mutex锁争用的采样情况</span>


## <span id="6">七、排查blocks堵塞操作的采样情况</span>
