

## 本节大纲
* [一、案例一分页子查询](#1)
* [二、案例二join问题](#2)
* [三、总结](#3)


## <span id="1">一、案例一分页子查询</span>
  之前做2B项目的时候，订单业务数据每天通过MQ异步发送到财务系统，进行财务结算，财务系统订单分页列表及各种纬度条件查询，目前该表t未分库分表，大约1000万条记录，sql如下格式：
~~~
select xxx,xxx,xxx,xxx from t where xxx=xxx and xx=xx limit 0 100
~~~
* 场景：比如查询到offset=50000条时，limit 100条，查询巨慢，大于10s，优化前sql：
~~~
select xxx,xxx,xxx,xxx from t where xxx=xxx and xx=xx limit 50000 100
~~~
* 优化后sql，使用主键条件子查询模式
~~~
select xxx,xxx,xxx,xxx from t 
 right join (
 select id from t where xxx=xxx and xx=xx limit 50000 100
)temp
on t.id=temp.id
~~~
* 比较优化前后sql


sql | 索引 | 回表查询 | buffer pool |耗时|
:-: | :-: | :-: | :-: |:-:
优化前 | 非聚集索引 | 需要回表查询，io多一倍 | 占用缓冲池空间大 |>10s
优化后 | 子查询非聚集索引+聚集索引| 不需要回表查询 | 占用缓冲池空间小 |<500ms

* 优化sql原理：减少回表二次查询，使用子查询查询主键，然后根据主键join，同时减少buffer pool的空间


## <span id="2">二、案例二join问题</span>
* 场景：做后台erp系统经常遇到多表查询，原则上 我们join表个数限制在3张之内，因为我们测试过，超过3张表，如果数据量超过千万级，查询速度巨慢，甚至超时，如果表数量少，join多个表问题也不大。
* 解决：
  - 1.3张表join 必须关联字段建立索引；
  - 2.如果不使用join的话，在业务逻辑层写简单sql，聚合数据
## <span id="3">三、总结</span>
* 案例一需要理解原理：1.Innodb索引底层实现B+tree 2.Innodb存储体系缓存池buffer pool使用
* 案例二超过三个表不允许使用join或者在业务逻辑层简单sql做数据聚合
