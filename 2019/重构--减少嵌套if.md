
### 1.优化if逻辑,最小化判断的条件（最有可能的放在前面）
~~~
if (i <= 10) {
    //TODO
 } else if (i > 10 && i < 100) {
   //TODO
 } else {
   //TODO
 }
~~~


### 2.switch/case 判断条件多的，例如经典返回值处理，有时候也可以使用hash/map，性能高O(1)
~~~
switch (rsp.state) {
        case 'SUCCESS':
            //TODO
            break;
        case 'FAIL':
            //TODO
            break;
        default :
            //TODO
    }
~~~

### 2.重构，用 OO 的继承或者组合(结合设计模式)
~~~
例子：
1.如果是qq，则qq登录
2.如果是wx，则wx登录
3.如果是phone，则phone登录
~~~
解决：
定义一个接口 loginer
~~~
type interface loginer{
    Login(ctx context.context)error
}
~~~
qq、wx、phone 分别在各自子类中去实现Login(),实质 设计模式中的模板方法





关于好代码：
代码短、可读性好、可维护行好
