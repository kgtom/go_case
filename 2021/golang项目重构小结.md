## 微服务设计中痛点
* 1.边界拆分: 粒度先粗后细，太细不方便执行cicd;单向调用，不能循环调用；尽量并行调用
* 2.保障高可用高并发: 缓存、容断、弹性伸缩、监控(没有数据度量就没有优化)
* 3.工程效率: 业务优先，技术支撑；自动化(脚手架)


## 高可用

### 负载均衡
* nginx 、consul 、consul-template
### 限流、熔断降级、回滚、隔离
* 限流：总并发qps、瞬时并发nginx limit_conn、令牌桶(平均流入速率，并允许一定程度的突然性流量，最大速率取决于桶的容量和生产token的速率)、漏桶(固定的流量流出速率)
* 熔断降级：超时降级、超时次数、故障降级
* 回滚：代码库回滚、部署版本回滚、数据库回滚
* 隔离：进程隔离、读写隔离、热点隔离、集群隔离、机房隔离、资源隔离
### 超时及重试
* 超时：设置合理的超时时间，服务端、客户端设置
* 重试：不能无限超时，否则会导致依赖服务的压力，比如：一定时段请求超时超过20%比例或者请求时间只剩100ms，就关闭重试,因为要优先保障正常的请求
### 批处理：拼接batch处理
* pipeline、batch处理方式
### 高阶功能
* 请求认证
* 链路追踪
* 监控指标
## 高并发
### 缓存：
* 性能：充分利用缓存(内存)操作。获取数据耗时久的，先缓存在redis中，减少db的访问
* 并发：充分利用缓存的高并发，减少db的压力(mysql的最大并发1500-2000),
### 消息队列：
* 异步：将请求先放在mq或者redis中，然后再请求
* 消峰：并发量大的时候，不直接访问db
### 并行：充分利用cpu
* golang的channel+goroutine+sync包
* channel: 协调多个逻辑片段，解耦生产者和消费者
* sync:某个数据结构内部状态的原子操作
### 性能优化：磁盘、网络、代码
* df、free、netstat、ss 、pprof
* 例如：排查接口超时，思路：
 - 1.网络 
 - 2.top--load 
 - 3.代码是否有死锁堵塞 
 - 4.runtime gc
 - 5.性能分析pprof，发现P的数量大，对应M就打，造成线程数调度耗时较大，从而引发接口超时，使用 runtime.NumCPU()  && GOMAXPROC

 ~~~
  curl http://127.0.0.1:8001/debug/pprof/trace?seconds=300 > trace.out 
  go tool trace trace.out
 ~~~

## 微服务工程化实践

### 工程化组织：统一框架、脚手架效率
### API定义：字段、错误码
### 配置文件
### 单元测试
### k8s部署

## 微服务指标监控
### 背景
 在微服务架构下，服务越来越多，定位问题复杂，需要监控服务的运行状态以及及时发出告警。服务的监控包括：日志监控、调用链路监控、指标监控等，其中指标监控是整个微服务监控中最重要的。
### 指标监控分类
* 基础监控：对服务基础设施的监控，包括对容器、虚拟机、物理机的内存、cpu的使用率
* 运行时监控：主要是GC的监控，包括GC次数、GC耗时、线程数量的监控
* 通用监控：流量监控(流量高峰、流量的增长情况)，耗时监控（平均耗时参考意义不大，使用中位数P90\P95\P99）
* 错误监控：通过对服务错误率的观察可以了解到服务当前的健康状态，包括：5xx、4xx、熔断、限流等
### 方案
* 基于 prometheus+grafana+alertmanager,不同维度监控、告警，使用 Histogram 定义多个Buckets进行分位统计；使用Counter类型记录服务请求总量、错误总数等
* 提前预测问题发生的可能性，包括 压测、资源的评估以及定期对服务巡检
### prometheus 局限
* prometheus 基于metric的监控，不适用 日志、调用链监控
* prometheus 默认pull模型，尽量不用push
* prometheus 单机模式对cpu和内存要求大，需要做集群化和水平扩展
* 监控系统 可用性>一致性


## 小结
* 重大版，必须在线下跑一周，隐藏问题需要时间才能暴露
* 线上有问题，第一时间回滚版本，及时止损、复盘问题
* 测试case务必充分，尤其时对外openAPI对各种参数校验
* 使用go mod管理pkg时，有些第三方pkg不遵循开源协议，特别注意版，本使用go mod vendor或者使用锁定的版本(replace gopkg.in/xxx.v1 => github.com/xxx/xxx v1.10.0)


## reference
* [wx](https://mp.weixin.qq.com/s/x4EEXq6-6xv-lm-dAazcag)
* [wx](https://mp.weixin.qq.com/s/8vASJavOQrXw5bGEEMwd9Q)
* [wx](https://mp.weixin.qq.com/s/4SzZEUTmjwAAsH5Qd2GWLQ)
* [gocn](https://gocn.vip/topics/10983)
* [se](https://segmentfault.com/a/1190000037435267)
* [tx](https://cloud.tencent.com/developer/news/568962)


