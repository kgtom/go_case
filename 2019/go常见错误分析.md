### 一、interface

 #### 1.nil 
 * value 和 type 都是 nil
 
 #### 2.empty
 
 #### 3.断言 


### 二、error

#### 1.标准库(公共错误)错误
 * 公共错误码统一维护，例如：500，404
#### 2.业务错误
* 业务代码各自微服务使用各自代码起始段位。例如：微服务端口8031，则错误代码范围8031001~8031999
#### 3.微服务API，各个微服务内部消化错误，给网关一个明确的错误。
* 例如： A调B，B调C，每一层调错误的时候，到底怎么处理的？到底从C传给B，B再透传给A，A传给用户呢？这种方法也不太好，因为最终A暴露给客户端，需要我们微服务内部消化错误，并且转化一个明确的错误抛给上游


### 三、context
* 利用context上下文解决元数据传递，超时传递，
* 在启动新的goroutine时候，既要保证上下文传递到位，又要规避 Context Cancel这种问题，我们需要写个func生成一个新的ctx，同时把原来ctx的key/value复制出来。或者直接使用context.Background()即可。

~~~go
func CopyCtx(ctx context) context.Context {
    ret := context.Background() 
    ret = context.WithValue(ret, ctxKey1, ctx.Value(ctxKey1)
    ret = context.WithValue(ret, ctxKey2, ctx.Value(ctxKey2)  
    return ret
}

~~~




> reference
* [mp](https://mp.weixin.qq.com/s/PLzA22yfSV_byckTTezl5Q)
