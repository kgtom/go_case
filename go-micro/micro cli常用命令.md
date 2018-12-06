### services list 
~~~
# tom @ tom-pc in ~/goprojects/src/github.com/micro/example [22:11:43] 
$ micro list services                                            
consul
go.micro.srv.example
topic:go.micro.srv.example

~~~

### 客户端调用
~~~
# tom @ tom-pc in ~/goprojects/src/github.com/micro/example [22:07:17] C:1
$ micro call  go.micro.srv.example  Example.Call '{"name":"tom"}'
{
        "msg": "Hello tom"
}

~~~


注意：Example.Call '{"name":"tom"}'，Call后面的空格。

### 健康检查
~~~
# tom @ tom-pc in ~/goprojects/src/github.com/micro/example [22:13:50] C:1
$ micro health go.micro.srv.example
service  go.micro.srv.example

version latest

node            address:port            status
go.micro.srv.example-752827bb-5ccf-4d32-bf85-bd0a945f47d6               192.168.1.94:52038              ok
~~~

> Reference:
* [mycodingnow](http://mycodingnow.com/blog/2018/07/2018-7-5-micro-cli.html)
