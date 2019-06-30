### 背景
* 在生产环境使用keys、smembers 时会带来性能问题，如果处理较大的KV、set、hash、zset时会导致服务器堵塞，造成负载较高。

### 解决方案
* 使用SCAN、SSCAN、HSCAN、ZSCAN每次执行都只会返回少量元素

### 举例

* 禁用keys ,使用 scan 返回新的游标，直到游标为0时退出迭代
* 禁用smembers,使用sscan 
~~~
10.1.1.10:6379[20]> sadd myset f1 f2 f3 d1 d2
(integer) 5
10.1.1.10:6379[20]> smembers myset
1) "d1"
2) "f3"
3) "f2"
4) "f1"
5) "d2"
10.1.1.10:6379[20]> sscan myset 0 match f*
1) "0"
2) 1) "f1"
   2) "f3"
   3) "f2"
10.1.1.10:6379[20]> sscan myset 0 match f* count 2
1) "1"
2) 1) "f1"
10.1.1.10:6379[20]> sscan myset 1 match f* count 2
1) "3"
2) 1) "f3"
   2) "f2"
10.1.1.10:6379[20]> sscan myset 3 match f* count 2
1) "0"
2) (empty list or set)
10.1.1.10:6379[20]>

~~~
