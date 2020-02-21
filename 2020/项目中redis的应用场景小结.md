

## 本节大纲
* [一、redis五种对象应用场景](#1)
* [二、redis常用应用场景](#2)
## 一、redis五种对象应用场景

### string

* 热点数据缓存：首页列表序列化后存储
* 计算器限流功能：按照天、小时、分钟等维度计数器+TTL+INCRBY
* 限时业务：验证码过期

### hash
* 用户玩家信息、房间信息
* 用户图片ID关联，尽量底层的Encoding使用ziplist，节约内存
* 日活流量PV、UV统计：hset 、hlen,流量小的时候可用，简单高效，流量大的考虑zset或bitset位图

### list
* 轻量级消息队列:push 、pop

### set
* 匹配池：spop、scard
* 房间池(推荐池)：srandmember、smove

### zset
* 排行榜：Zrangebyscore
* 用户关系(好友、关注、关注)：各自关系及交集Zinterstore
* 延迟队列：购买钻石，5分钟内未支付的，订单自动取消。score是时间戳，使用 Zrangebyscore 指定区间的member
* 去重池:Zscore是否存在或者Zremrangebyrank缩减容量
* 匹配池(推荐池)：Zrangebyscore 、zrem
* 日活PV、UV统计：Zrangebyscore，流量小可用考虑使用hash，流量大考虑zset或bitset

## redis常用应用场景
### 分布式锁
* setNX:注意死锁、续命锁
### 延迟队列
* zset：Zrangebyscore指定score时间戳区间内数据执行
### 游戏场景
* 用户玩家、房间信息：hash
* 匹配池(推荐池)：set、zset
### 订阅推送服务
* 背景：订阅购买某种优惠劵，大约30w+人，要求在开始购买优惠券的前5分钟通知订阅用户
* 要求：实效性，及时推送给订阅用户；稳定性，用户量大的情况下，确保都被推送到
* 技术选型：
  - 1.传统定时扫描db：缺点 db磁盘io是瓶颈问题
  - 2.kafka延迟队列：会有订阅-取消订阅--再次订阅，需要单独去重，逻辑复杂
  - 3.redis的zset做延迟队列：参考分表思路，根据user_id%10(比如10个zset队列,把提醒用户时间做score)，后台10台server获取队列。10台server保证均衡的方式获取队列的数据，简单采用redis中incr自增id%10得到获取第几个队列的编号。

### 日活流量PV、UV统计
* 流量小用：hash
* 流量大用：zset、bitset，偏向bitset，二进制存储，占内存极少，1亿用户：1亿/1024/1024/8约为12M
* 流量大，概率统计，使用HyperLogLog概率算法：pfadd、pfcount 存在0.81%误差
