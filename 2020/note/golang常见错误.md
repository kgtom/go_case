
### 阻塞的Goroutine和资源泄露
Rob Pike在2012年的Google I/O大会上所做的“Go Concurrency Patterns”的演讲上，说道过几种基础的并发模式。从一组目标中获取第一个结果就是其中之一。

~~~
func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
~~~
这个函数在每次搜索重复时都会起一个goroutine。每个goroutine把它的搜索结果发送到结果的channel中。结果channel的第一个值被返回。

那其他goroutine的结果会怎样呢？还有那些goroutine自身呢？

在 First() 函数中的结果channel是没缓存的。这意味着只有第一个goroutine返回。其他的goroutine会困在尝试发送结果的过程中。这意味着，如果你有不止一个的重复时，每个调用将会泄露资源。

为了避免泄露，你需要确保所有的goroutine退出。一个不错的方法是使用一个有足够保存所有缓存结果的channel。

~~~
func First(query string, replicas ...Search) Result {  
    c := make(chan Result,len(replicas))
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
~~~
另一个不错的解决方法是使用一个有default情况的select语句和一个保存一个缓存结果的channel。default情况保证了即使当结果channel无法收到消息的情况下，goroutine也不会堵塞。

~~~
func First(query string, replicas ...Search) Result {  
    c := make(chan Result,1)
    searchReplica := func(i int) { 
        select {
        case c <- replicas[i](query):
        default:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
~~~
你也可以使用特殊的取消channel来终止workers。

~~~
func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    done := make(chan struct{})
    defer close(done)
    searchReplica := func(i int) { 
        select {
        case c <- replicas[i](query):
        case <- done:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }

    return <-c
}

~~~

> reference
[cnblogs](https://www.cnblogs.com/jhhe66/articles/9232691.html)
