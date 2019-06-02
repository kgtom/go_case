### 问题 高峰时提示：Error allocating resoures for the client，不能为客户端分配链接

* 1.info clients查看当前的redis连接数

~~~
abc.aliyuncs.com:6379> info clients
# Clients
connected_clients:9640
client_longest_output_list:0
client_biggest_input_buf:27
blocked_clients:8

~~~

* 2.config get maxclients 可以查询redis允许的最大连接数。

~~~

abc.aliyuncs.com:6379> config get maxclients
1) "maxclients"
2) "10000"
~~~


* 3.client list 查看目前连接详情

~~~
abc.aliyuncs.com:6379> client list
d=3755278 addr=10.1.0.12:42434 fd=93 name= age=131466 idle=207 flags=N db=140 sub=0 psub=0 multi=-1 qbuf=0 qbuf-free=0 obl=0 oll=0 omem=0 events=r traffic-control=NULL cmd=sadd type=user next_opid=-1

......
~~~

发现客户端idle 空闲时间太长，连接池维护太多连接没有及时释放。

* 4. CONFIG SET timeout 30 手动释放

~~~
abc.aliyuncs.com:6379> CONFIG SET timeout 30
OK
~~~
* 5.再次 client list 查看，idle 空闲时长30一下了。

### 总结
 未设置合理的idle 超时时长，导致连接未被及时释放，客户端分配不到资源。
