
### 背景
两个定时任务(console-A、console-B 两个互为备份)同时处理订单状态，为了避免两个任务同时抢占同一笔订单，使用分布式锁，谁先处理orderID=1001的订单，则优先处理状态。

### lock代码

~~~ go
package lock

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/hkspirt/redislock"
)

const (
	DefaultInterval       = 100 * time.Millisecond //每隔100ms尝试
	orderKeyRetry      = 2                      //尝试2次
	orderKeyExpire     = 1 * time.Minute        //锁定1min
	orderLockKeyPrefix = "order|lk|order"
)

// RedisLock redis锁
type RedisLock struct {
	source  string
	rlock   *redislock.RedisLock
	rclient *redis.Client
}

// NewOrderLock 生成 提现锁
func NewOrderLock(orderID string) *RedisLock {
	source := fmt.Sprintf("%s:%s", orderLockKeyPrefix, orderID)
	rl := &RedisLock{
		source: source,
		rlock: redislock.NewRedisLockWithParam(
			source,
			orderKeyExpire,
			orderKeyRetry,
			DefaultInterval,
			false, //此处为false, 解锁和上锁不一定是同一个进程
		),
		rclient: redisCli,
	}
	return rl
}

// Lock 上锁
func (rl *RedisLock) Lock() error {
	return rl.rlock.Lock(rl.rclient)
}

// UnLock 解锁
func (rl *RedisLock) UnLock() error {
	return rl.rlock.UnLock(rl.rclient)
}

// Refresh 刷新过期时间
func (rl *RedisLock) Refresh() error {
	return rl.rlock.Refresh(rl.rclient)
}

~~~

### 封装lock
~~~go
//全局订单 行级锁
func AcquireOrderLock(ctx context.Context, orderID string) (*lock.RedisLock, bool) {
	rl := lock.NewOrderLock(orderID)
	err := rl.Lock()
	if err != nil {
		return nil, false
	}
	return rl, true
}

~~~

### 调用

~~~
for _,v:=range orderList{

AcquireOrderLock(v.orderID)

//此处实现自己的业务逻辑
}
~~~
