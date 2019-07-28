

## 一、理解概念
   
 ### 整个系统由生产者、Broker Server和消费者三部分组成，生产者和消费者由开发人员编写，通过API连接到Broker Server进行数据操作。

 * 1.Topic 是Kafka下消息的类别，这是逻辑上的概念，用来区分、隔离不同的消息数据，开发中只需要关注写入到哪个topic及从哪个topic消费即可。

 * 2.Partition 是Kafka下数据存储的基本单元，这个是物理上的概念。同一个topic的数据，会被分散的存储到多个partition中，这些partition可以在同一台机器上，也可以是在多台机器上。这种分布式存储在mongodb、es 都可以见到，优势明显：利于水平扩展，避免单台机器磁盘空间和性能的问题；数据冗余利于提高容灾能力。

 * 3.Consumer Group 同样是逻辑上的概念，是Kafka实现单播和广播两种消息模型的手段。
同一个topic的数据，会广播给不同的group；同一个group中的 consumer，只有一个consumer  能拿到这个数据。换句话说，对于同一个topic，每个group都可以拿到同样的所有数据，但是数据进入group后只能被其中的一个 consumer 消费。
consumer 的数量通常不超过partition的数量，且二者最好保持整数倍关系，因为Kafka在设计时假定了一个partition只能被一个worker消费（同一group内）。

* 为了便于实现MQ中的多播，重复消费等引入的概念。如果ConsumerA以及ConsumerB同在一个UserGroup，那么ConsumerA消费的数据ConsumerB就无法消费了。
因为所有user group中的consumer共使用一套offset偏移量。

* 单播：一个topic 只有一个group 订阅

* 4.Broker
物理概念，指服务于Kafka的一个node。



### 一个典型的消息队列 Kafka 集群包含：
* 		Producer：通过 push 模式向消息队列 Kafka Broker 发送消息，可以是网站的页面访问、服务器日志等，也可以是 CPU 和内存相关的系统资源信息；
* 		Kafka Broker：消息队列 Kafka 的服务器，用于存储消息；支持水平扩展，一般 Broker 节点数量越多，集群吞吐率越高；
* 		Consumer Group：通过 pull 模式从消息队列 Kafka Broker 订阅并消费消息;
* 		Zookeeper：管理集群的配置、选举 leader，以及在 Consumer Group 发生变化时进行负载均衡。


## 二、理解 发布—订阅 和生产—消费

消息队列 Kafka 的 Pub/Sub 模型
消息队列 Kafka 采用 Pub/Sub（发布/订阅）模型，其中：
* 		Consumer Group 和 Topic 的关系是 N:N。 同一个 Consumer Group 可以订阅多个 Topic，同一个 Topic 也可以同时被多个 Consumer Group 订阅。
* 		同一 Topic 的一条消息只能被同一个 Consumer Group 内的任意一个 Consumer 消费，但多个 Consumer Group 可同时消费这一消息。
* 


## 三、应用场景

* 收集用户行为，不同行为创建不同topic,进行异步处理
* 日志聚合
* 数据统计，将数据从kafka导入到maxcompute


## 四、总结
生产者(producer)将消息记录(record)发送到kafka中的主题中(topic), 一个主题可以有多个分区(partition), 消息最终存储在分区中，topic 发布到group中，group中的消费者(consumer)最终从主题的分区中获取消息。

> reference

* [aliyun](https://help.aliyun.com/document_detail/68152.html?spm=a2c4g.11186623.2.12.54a472e4JsbKBI)
* [csdn](https://blog.csdn.net/kuluzs/article/details/71171537)
