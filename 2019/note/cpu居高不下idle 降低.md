## 场景 cpu 占用升高，idle 降低

### 1.先理解 G的来源
   在main函数执行期间，运行时系统会根据go语句，复用或者新建goroutine来执行go函数，这些goroutine放在p的可运行的G队列。
   如果因为调用超时，导致创建了更多的goroutine，这部分goroutine在GC 时不会被回收，虚增了系统的并发数。

### 2.GC的回收
 GC 不会回收的goroutine,简单粗暴的方式，重启服务。

### 3.服务抖动或者超时，导致没有可以G，需要new G，导致cpu升高，idel 降低
