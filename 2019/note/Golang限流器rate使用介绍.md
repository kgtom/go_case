

golang标准库中就自带了限流算法的实现，即golang.org/x/time/rate。

该限流器是基于Token Bucket(令牌桶)实现的。


### 构造一个限流器
我们可以使用以下方法构造一个限流器对象：

~~~
limiter := NewLimiter(10, 1);
~~~

这里有两个参数：

* 第一个参数是r Limit。代表每秒可以向Token桶中产生多少token。Limit实际上是float64的别名。
* 第二个参数是b int。b代表Token桶的容量大小。
那么，对于以上例子来说，其构造出的限流器含义为，其令牌桶大小为1, 以每秒10个Token的速率向桶中放置Token。

除了直接指定每秒产生的Token个数外，还可以用Every方法来指定向Token桶中放置Token的间隔，例如：

~~~
limit := Every(100 * time.Millisecond);
limiter := NewLimiter(limit, 1);
~~~

以上就表示每100ms往桶中放一个Token。本质上也就是一秒钟产生10个。

Limiter提供了三类方法供用户消费Token，用户可以每次消费一个Token，也可以一次性消费多个Token。
而每种方法代表了当Token不足时，各自不同的对应手段。

>reference
[cyhone](https://www.cyhone.com/articles/usage-of-golang-rate/)
