性能优化(数据库、cache、ulimit 、tcp)

测试环境：
centos7.5 4核4G ssd硬盘 mysql 5.7，最大连接数500，reids 2G
要求：数据库负载cpu <50% ,服务器负载cpu <70% ,p99 100ms,错误率 <1%
压测：jemeter、k6 

### mysql(Too many connections)
并发数理论值：2000左右

~~~
show variables like '%max_connections%'; 
~~~
设置最大连接数：通常最大连接数是并发的80%左右，1500,实际
innodb_buffer_pool_size：索引和数据缓冲区大小，一般设置物理内存的60%-70%
innodb_log_buffer_size = 8M   一般不用超过16M

### cache
配置信息，先从db 获取，缓存cache,定期同步

### ulimit(too many files open)
修改系统最大连接数
~~~
# vi /etc/security/limits.conf  #修改配置，*代表所有用户，也可指定用户设置
* soft nofile 65535
* hard nofile 65535
# ulimit -SHn 65535   #立刻生效
~~~

### tcp系统内核参数(timewait、没有可用端口)

查看time_wait 数量
~~~

(base) pc-Pro :: ~ » netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
FIN_WAIT_2 1
CLOSE_WAIT 5
TIME_WAIT 3
ESTABLISHED 175
~~~
net.ipv4.tcp_keepalive_time = 1200 当keepalive起用的时候，TCP发送keepalive消息的频度。缺省是2小时，改为20分钟
net.ipv4.tcp_fin_timeout = 30 #TIME_WAIT超时时间，默认是60s 
net.ipv4.tcp_tw_reuse = 1     表示开启重用。允许将TIME-WAIT sockets重新用于新的TCP连接，默认为0，表示关闭
net.ipv4.tcp_tw_recycle = 1  表示开启TCP连接中TIME-WAIT sockets的快速回收，默认为0，表示关闭；
net.ipv4.ip_local_port_range = 1024 65000  表示用于向外连接的端口范围。缺省情况下很小：32768到61000，改为10000到65000

net.ipv4.tcp_max_tw_buckets = 6000 表示系统同时保持TIME_WAIT的最大数量，如果超过这个数字，TIME_WAIT将立刻被清除并打印警告信息。默 认为180000，改为6000
  
net.ipv4.tcp_max_syn_backlog = 8192 表示SYN队列的长度，默认为1024，加大队列长度为8192，可以容纳更多等待连接的网络连接数



### 硬盘
ssd 代替sata; raid 1-5,减少数据库io压力


