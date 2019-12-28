## 一、背景
 生产环境上用户图片的微服务，redis内存占用一直持续增长，存储是用户上传图片picID和userID的映射。通过用户分享的图片picID找到对应userID,目前存储KV 字符串对象存储。
 大约有200万左右用户选择分享图片，人均分享10张，大约2000万条记录。
 
 * 使用 redis string对象，key=md5(picID) value=userID，大约2.2G
![2.2g](https://github.com/kgtom/back-end/blob/master/pic/2.2g.png)
 
 * 更换redis hash对象存储,直接看图，节省了接近85%内存
![315m](https://github.com/kgtom/back-end/blob/master/pic/315m.png) 
 
 ## 二、hash对象底层实现
  hash对象的编码分ziplist和hashTable，当存储 field-value数量不超512并且field 、value字符串长度都不超64字节，使用ziplist，反之使用hashTable。
 ### ziplist
  redis节约内存而开发的顺序型数据结构,每个节点可以保存一个字节数组或者整数值。新增、获取单个值O(1),获取所有、删除需要遍历O(N)
  当新的field-value加入到hash时，先将保存field的ziplist节点放在表尾，再将value放在表尾，两者紧挨着，再来一个新的则放在之前的后面。即先来的放表头方向，后来的放表尾方向。
~~~ 
10.1.183.247:6379[29]> hset user name tom
(integer) 1
10.1.183.247:6379[29]> hset user age 18
(integer) 1
10.1.183.247:6379[29]> object encoding user
"ziplist"
~~~
ziplist压缩列表底层实现如下：
 ~~~
 zlbytes zltail zllen "name" "tom" "age" "18" zlend
 ~~~
 
 ### hashTable
 hashTable编码的hash对象使用字典map作为底层实现。每一个键值对使用字典的键值对来保存。
 字典map的底层是哈希表dictht实现，一个哈希表里面可以有多个哈希节点dictEntry，每个哈希节点保存了字典中的键值对和一个单向链表(解决哈希冲突，将相同值的键值对连在一起)。实际数组(多个哈希节点)+链表。新增、删除、获取单个O(1)，获取所有O(N)。
 ~~~
 -- value 超过64字节
 10.1.183.247:6379[29]> hset user addr aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
(integer) 0
10.1.183.247:6379[29]> object encoding user
"hashtable"
10.1.183.247:6379[29]> type user
hash
 ~~~
 
  ## 三、具体实现
  
  了解hash对象的底层实现后，我们应该知道，应该使用ziplist来存放kv信息，因为节省内存啊。
  但ziplist 需要满足两个条件，键值对field-value 不超过512并且键值对不超过64字节。
  虽然这两个值是可修改的，通过hash-max-ziplist-vlaue/entries。做法如下：
  * 1.不修改配置的情况下，计算出需要多少个hash对象，2k万/512=39062.5,为了尽可能平均分配到不同hash,我们分配50000个hash。
  * 2.将键值对 filed 不在MD5,直接使用picID/50000,决定分配到哪个hash。
  * 3.hset pic_user0 100 100000001。
  
  ## 四、总结
  * ziplist存储，因为连续内存区间，减少内存，缺点不能单独设置某个键值对过期时间，需要console定期在非高峰时段扫描清理
  * 字符串KV存储优点存储、获取、ttl都方便，只是内存占用较大，因为底层是动态字符串SDS，空间预分配，分配额外空间。
