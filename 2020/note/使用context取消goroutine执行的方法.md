
### 为什么需要取消功能
* 举个场景，http 请求，上游已经取消了请求，下游涉及查询db或者redis的自研，应该立刻取消，减少资源浪费
### 取消功能场景

* 1.基于context请求下游服务

~~~
quit := make(chan struct{}, 1)
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go func() {
		// do something 
		quit <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		//fmt.Println("done....",ctx.Err())
	case <-quit:
		//fmt.Println("quit.....")
		break
	}
~~~
 或者
~~~
  ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	ret, err := client.CommonClassify(ctx, req)

~~~
* 2.基于http请求，当上游请求取消后，应该取消http请求，避免资源浪费

~~~
  
  req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		fmt.Println("err....",err)
		return false, ""
	}

	ctx, _ = context.WithTimeout(ctx, 1000*time.Millisecond)

	// 使用短连接，参考：https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi
	req.Close = true

	client := &http.Client{
		Transport: &http.Transport{DisableKeepAlives: true},
		Timeout:   3 * time.Second, //超时保护
	}
	//带有超时时间ctx与req的关联
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err....",err)
		return false, ""
	}
~~~


>reference:
