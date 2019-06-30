#### 一、cache 的使用 
* 1.注意设置 ttl ,避免垃圾数据
* 2.zset 和set 根据不同场景使用，比如随机的获取 userID ,用set就可以，比如 zset 中score的使用用户上下线时间ms
* 3.pipeline 批处理写，减少 磁盘IO
* 4.transaction 事务的使用

#### 二、首页推荐列表
* 1.技术与业务的权衡，保护服务的考虑，可以接收有损方案
 比如，先来个组兜底的数据，再快速 goroutine 生成全量数据，然后下次获取到的就是全量数据，保证除第一次、超过1分钟cache 失效外，其他获取数据走cache。
* 或者考虑登录时回调，开启gorouine 快速生成数据cache


#### 三、context 在开启 goroutine 时，记得ctx 上下文传值别丢了
~~~
context.WithValue()
~~~
