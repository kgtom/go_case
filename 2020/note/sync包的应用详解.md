
### sync.Metux
 对共享资源的互斥访问，保证多goroutine时，有且只有一个G访问资源
 
 ~~~
 mux := &sync.Mutex{}

mux.Lock()
// Add or Update共享资源变量 (例如切片，结构体指针等)
mux.Unlock()
 ~~~

### sync.RWMetux

* 并发读G没有锁，但读的同时，如果有写，则写需要等待读锁释放
* 写锁互斥
* 使用场景：读多，写少

### sync.WaitGroup
同步元语，等待一组G的完成，内部计数器实现

~~~
wg := &sync.WaitGroup{}

for i := 0; i < 5; i++ {
  wg.Add(1)
  go func(n int) {
    // do Something
    wg.Done()
  }(i)
}

wg.Wait()
// go on do something
~~~

### sync.Map
* 一个并发安全的Map
* 使用场景：1.多读少写；2.写和读不同分区的key
* 官方对普通map性能优化了不少，大多数场景使用map+metux代替sync.Map


### 总结
* 使用指针类型，不能复制，使用指针传递！！！

