## 概括

 * 包括 grpc 重试、负载均衡(代理、客户端)、超时控制、健康检查
### 重试(包括熔断降级)
 TODO

### 负载均衡

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


### 超时控制
TODO


> reference
 * [grpc-retry](https://github.com/grpc/grpc-go/tree/master/examples/features/retry)

 * [nobugware-k8s](https://blog.nobugware.com/post/2019/kubernetes_mesh_network_load_balancing_grpc_services/
)
