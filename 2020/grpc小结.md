## 概括

 * 包括 grpc 重试、负载均衡(代理、客户端)、超时控制、健康检查
## 重试(包括熔断降级)
 

### 重试解决的问题？
* 短时间不可用，eg 网络抖动，硬件故障(路由和负载均衡等)、资源没有隔离好，总之网络是不可靠的：重试
* 长时间不可用，光纤北被挖：不使用重试，应该指数退避重试BACKOFF：第一次重试然后等待间隔

### 短时间故障处理方案
* 感知，感知不同错误类型，grpc 有自己一套 status code 错误码，例如 UNAVAILABLE、ABORTED、DEADLINE_EXCEEDED
* 处理：不用类型选择不同重试处理策略，例如：网络抖动，重试，网络不可用，指数退避(固定、立刻、随机)重试

### 重试策略分类
* 重试策略，失败后立即重试
* 对冲策略，只用于幂等的操作，多次请求，有一个成功返回即可


#### 重试时间策略
* 客户端client 默认关闭了重试的，需要在环境变量中设置 GRPC_GO_RETRY=on 来开启重试,例如：GRPC_GO_RETRY=on go run client.go
* MaxAttempts = 2，即最多尝试 2 次，即也就是最多重试 1 次
* RetryableStatusCodes 设置了 UNAVAILABLE，可以解决：rpc err: code = Unavailable desc = transport is closing
* RetryableStatusCodes 中设置 DeadlineExceeded 或者Canceled 是无用的，因为在重试逻辑的代码里发现 ctx超时或取消就会立即退出重试逻辑,不尽兴重试

~~~

func main() {

	retryPolicy := `{
        "methodConfig": [{
          "name": [{"service": "sample.Sample"}],
          "waitForReady": true,
          "retryPolicy": {
                  "MaxAttempts": 3,
                  "InitialBackoff": "0.1s",
                  "MaxBackoff": "1s",
                  "BackoffMultiplier": 2.0,
                  "RetryableStatusCodes": [ "UNAVAILABLE" ]
          }
        }]}`

	
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy), grpc.WithBalancerName(roundrobin.Name)) //k8s\dns都使用)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := client.NewSampleClient(conn)

	// Contact the server and print out its response.
	name := "hello world222"

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &client.STSamplePingReq{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s,err:%v", r.Message, err)

}
~~~


#### 对冲策略
不建议采用


#### 限流策略
防止服务器过载，采用限流配置，现在 重试和对冲策略
~~~
"retryThrottling": {
  "maxTokens": 6,
  "tokenRatio": 0.1
}
~~~


## 负载均衡

* 1.客户端负载均衡： 优点：k8s api watch ep,低延迟， 缺点：客户端实现成本高

~~~
//method one: k8s api ,watch ep(ps:headless srv or normal srv is ok,do not forget k8s cfg serviceAccount)
kuberesolver.RegisterInCluster()
conn, err := grpc.Dial(
	"k8s-customer:///svcName.ns:8001",
	grpc.WithInsecure(),
	grpc.WithBalancerName(roundrobin.Name),
)
~~~
* 2.负载均衡代理，基于k8s svc(有头服务) 做负载均衡优点：对于客户端简单，适用于短连接，缺点：不均衡
~~~
//method two: k8s dns,find dns A record，then find pod ip(headless）

conn, err := grpc.Dial(
	svrName.ns:8001,
	grpc.WithInsecure(),
	grpc.WithTimeout(defaultDialTimeout*time.Millisecond),
	grpc.WithBalancerName(roundrobin.Name),
)
                                        //k8s svc 做负载均衡
~~~~

* 3.负载均衡代理，基于k8s svc(无头服务，headless,clusterIP: None)的dns,优点：简单，缺点：dns默认刷新30min，有点长，可调节时长

~~~
method three: k8s serviceName,find vip(clusterIP)
ctx, _ := context.WithTimeout(context.Background(), defaultDialTimeout*time.Millisecond)
conn, err := grpc.DialContext(ctx, "dns:///svrName.ns:8001", grpc.WithInsecure())

~~~


## 超时控制

### 针对每次方法调用

~~~
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &client.STSamplePingReq{Name: name})
~~~

### 针对建立连接后，每次ctx 
~~~
defaultDialTimeout:=3000ms
conn, err := grpc.Dial(
	svrName.ns:8001,
	grpc.WithInsecure(),
	grpc.WithTimeout(defaultDialTimeout*time.Millisecond),
	grpc.WithBalancerName(roundrobin.Name),
~~~



> reference
 * [grpc-retry](https://github.com/grpc/grpc-go/tree/master/examples/features/retry)

 * [nobugware-k8s](https://blog.nobugware.com/post/2019/kubernetes_mesh_network_load_balancing_grpc_services/
)
