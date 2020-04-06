
## 本篇主要介绍如何实现超时控制，主要有三种
* 一、context.WithTimeout/context.WithDeadline + time.After
* 二、context.WithTimeout/context.WithDeadline + time.NewTimer
* 三、channel + time.After/time.NewTimer


## 一、context.WithTimeout/context.WithDeadline + time.After

~~~
func AsyncInvoke() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
	defer cancel()
	go func(ctx context.Context) {
		// 发送HTTP请求
	}()

	select {
	case <-ctx.Done():
		fmt.Println("invoke success")
		return
	case <-time.After(time.Duration(time.Millisecond * 600)):
		fmt.Println("timeout")
		return
	}
}
~~~
小结：
1、通过context的WithTimeout设置一个有效时间为500毫秒的context。
2、该context会在耗尽500毫秒后或者方法执行完成后结束，结束的时候会向通道ctx.Done发送信号。
3、已经设置了context的有效时间，为什么还要加上这个time.After？
  这是因为该方法内的context是自己申明的，可以手动设置对应的超时时间，但是在大多数场景，这里的ctx是从上游一直传递过来的，对于上游传递过来的context还剩多少时间，我们是不知道的，所以这时候通过time.After设置一个自己预期的超时时间就很有必要了。
4、注意，这里要记得调用cancel()，不然即使提前执行完了，还需要等到800毫秒后context才会被释放。



## 二、context.WithTimeout/context.WithDeadline + time.NewTimer

~~~
func AsyncInvoke() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond * 500))
	defer cancel()
	timer := time.NewTimer(time.Duration(time.Millisecond * 600))

	go func(ctx context.Context) {
		// 发送HTTP请求
	}()

	select {
	case <-ctx.Done():
		timer.Stop()
		timer.Reset(time.Second)
		fmt.Println("invoke success")
		return
	case <-timer.C:
		fmt.Println("timeout")
		return
	}
}
~~~
小结：
如果接口调用提前完成，则监听到Done信号，然后关闭定时器。否则，会在指定的timer即600毫秒后执行超时后的业务逻辑。

## 三、channel + time.After/time.NewTimer

~~~
func AsyncInvoke() {
  ctx := context.Background()
	done := make(chan struct{}, 1)

	go func(ctx context.Context) {
		// 发送HTTP请求
    
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("invoke success")
		return
	case <-time.After(time.Duration(600 * time.Millisecond)):
		fmt.Println("timeout")
		return
	}
}
~~~
小结：
1.使用channel处理goroutine之间通信，处理结束后，发送done
2.监听done信号，如果在time.After超时时间之前接收到，则正常返回，否则走向time.After的超时逻辑，执行超时逻辑代码。

>reference
 * [juejin](https://juejin.im/post/5e774a73e51d4526c70fd0a4)
