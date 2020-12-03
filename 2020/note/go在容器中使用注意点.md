
### 一门新语言会关注三方面：
* 语言特性：简洁、快速、高并发
* 应用场景，解决啥问题：云原生容器开发、服务端高并发场景
* 坑，需要注意地方： gorutine并行、gc内存管理、容器化使用(DNS在Docker集成nscd、CGO与镜像、Docker集成监控Prometheus、高性能内核参数somaxconn)

### DNS

配置容器中的 DNS 选项
您可以在容器服务的编排模板中通过指定 dns 和 dns_options 来指定容器的 DNS 服务器和 DNS 的选项。

比如:

~~~
testdns:
  image: nginx
  dns:
    - '8.8.8.8'
    - '8.8.4.4'
  dns_options:
    - use-vc
    - no-tld-query
 
 ~~~
    
上面的示例配置了服务的容器的 DNS 服务器和 DNS 的查询选项。

说明 为了进行服务发现，Docker 在每个容器中内置了 DNS 服务器。您在容器中看到的 /etc/resolv.conf 文件中的 DNS 服务器是 Docker 的内置 DNS 服务器 127.0.0.11。Docker 会监听内置服务器的 DNS 请求，并将 DNS 请求转发到通过 dns 配置的服务器上。
优化 DNS 解析
在请求域名时，DNS 解析可能会超时或者失败导致网站无法访问。操作系统上一般会启用 nscd 服务用于做 DNS 的缓存以便避免 DNS 解析失败。但容器的镜像中一般不会配置 nscd 服务，您可以在经常做 DNS 解析的容器上安装 nscd 服务来优化容器中的 DNS 解析。

您需要首先安装 nscd 软件包，然后在容器启动的时候首先启动 nscd 服务，再启动自己的进程。

~~~
FROM registry.aliyuncs.com/acs/ubuntu
RUN apt-get update && apt-get install -y nscd && rm -rf /var/lib/apt/lists/*
CMD service nscd start; bash
~~~

### 修改内核参数

somaxconn是限制了接收新 TCP 连接侦听队列的大小，它的默认值是128，高并发场景下需要修改该值，提供一种yml修改方式，如下：
~~~
apiVersion: v1
kind: Pod
metadata:
  name: test-sysctl
  annotations:
    security.alpha.kubernetes.io/unsafe-sysctls: net.core.somaxconn=65535
spec:
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 80
      protocol: TCP
  nodeSelector:
    kubernetes.io/hostname: cn-xxx.i-xxxxx    
~~~

> reference:
   * [js](https://www.jianshu.com/p/3214db41ce48)
   * [aly](https://help.aliyun.com/knowledge_detail/54697.html)
   * [aly](https://developer.aliyun.com/article/603745)
   
