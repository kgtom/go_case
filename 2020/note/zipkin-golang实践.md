
一、接入代码库参考


grpc:
使用的grpc.StatsHandler中间件用于跟踪gRPC服务器和客户端请求。

服务端：
~~~
import (
    "google.golang.org/grpc"
    zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
)
 
 
tracer := GetTracer("demoService", "127.0.01")
 
 //默认span
server = grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))
 
 
 // 自定义span :tracer can now be used to create spans.
   span, ctx := zkTracer.StartSpanFromContext(ctx, "server ping")
   // ... do some work ...
   span.Finish()
~~~

客户端：
~~~
import (
    "google.golang.org/grpc"
    zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
)
 
 
//默认span
conn, err = grpc.Dial(addr, grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)))
//自定义span
 
ctx := context.Background()
span, ctx := zkTracer.StartSpanFromContext(ctx, "client ping")
// ... do some work ...
span.Finish()
~~~

http：
提供了一种易于使用的http.Handler中间件来跟踪服务器端请求



参考类库：https://zipkin.io/pages/extensions_choices.html

