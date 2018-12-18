## 本节知识
* [一、consul功能介绍](#1)
* [二、服务注册](#2)
     * [1.json配置文件方式](#21)
     * [2.http api方式](#22)
* [三、服务发现](#3)
     * [1.命令行](#31)
     * [2.代码](#32)
* [四、kv使用](#4)

### <span id="1">一、consul功能介绍</span>
一、概述

consul是google开源的一个使用go语言开发的服务发现、配置管理中心服务。
包括服务注册与发现框架、分布一致性协议实现、健康检查、Key/Value存储、多数据中心方案。
每个节点都需要运行agent，他有两种运行模式server和client。每个数据中心官方建议需要3或5个server节点以保证数据安全，同时保证server-leader的选举能够正确的进行。

**client**
是consul节点的一种client模式，注册到当前节点的服务会被转发到SERVER。

**server**

是consul的server模式。功能和CLIENT都一样。

**server-leader**

它需要负责同步注册的信息给其它的SERVER，同时也要负责各个节点的健康监测。

**raft**

server节点之间的数据一致性保证，一致性协议使用的是raft，而 zk 用的paxos，etcd用raft。

**服务发现协议**

consul采用http和dns协议，etcd只支持http

**服务注册**

consul支持两种方式实现服务注册。
* 一种服务自己调用API实现注册，
* 另一种方式是通过json个是的配置文件实现注册，consul官方建议使用第二种方式。




**服务发现**

consul支持两种方式实现服务发现。
* 一种是通过http API来查询有哪些服务，
* 另外一种是通过consul agent 自带的DNS（8600端口），域名是以NAME.service.consul的形式给出，NAME是服务的名称。

**服务间的通信协议**

Consul使用gossip协议(是一种去中心化、容错并保证最终一致性的协议)。包括(lan/wan gossip)


**安装成功后**
访问ui：

http://192.168.1.198:8500/ui

常用端口：
~~~
8300：consul agent服务

8301：lan gossip

8302：wan gossip

8500：http api端口

8600：DNS服务端口
~~~

**为什么选择consul**
* 简单易用，不需要集成sdk
* 自带UI及健康检查
* 支持多数据中心

### <span id="2">二、服务注册</span>

#### <span id="21"> 第一种：配置文件服务注册</span>
~~~
# tom @ tom-pc in / [19:15:32] C:1
$ sudo mkdir consul.conf
# tom @ tom-pc in /consul.conf [19:21:27] C:1
$ sudo vi web_8082.json
~~~

内容如下：
~~~
 {
     "service": {
       "name": "web",
       "tags": ["master"],
       "address": "127.0.0.1",
       "port": 8082,
       "checks": [
         {
           "http": "http://localhost:8082/health",
          "interval": "10s"
        }
      ]
    }
  }
~~~

web服务代码如下：
~~~go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello consul--web")
	fmt.Fprintf(w, "hello consul--web")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello health check from consul--web!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8082", nil)
}

~~~

* 运行 8082  
* 运行consul，带上-config-dir命令行选项

这样在./consul.conf下的所有文件扩展为.json的文件都会被Consul agent作为配置文件读取。
~~~
# tom @ tom-pc in ~ [19:29:32] C:1
$ consul agent -dev -config-dir=/consul.conf
==> Starting Consul agent...
==> Consul agent running!
           Version: 'v1.4.0'
           Node ID: '0a534ae7-ed52-03a2-3b81-79bb9c36032c'
         Node name: 'tom-pc'
        Datacenter: 'dc1' (Segment: '<all>')
            Server: true (Bootstrap: false)
       Client Addr: [127.0.0.1] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
      Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
           Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false

~~~

* 如果 web8082服务未运行，则consul 提示
~~~
2018/12/17 19:40:47 [WARN] agent: Check "service:web" HTTP request failed: Get http://localhost:8082/health: dial tcp [::1]:8082: connect: connection refused
~~~
* 如果web 服务运行正常，则：
~~~

    2018/12/17 19:48:17 [DEBUG] agent: Check "service:web" is passing
    2018/12/17 19:48:17 [DEBUG] agent: Service "web" in sync
    2018/12/17 19:48:17 [INFO] agent: Synced check "service:web"

~~~
* 也可以访问 UI http://127.0.0.1:8500/ui/dc1/services 查看。

#### <span id="22">第二种 服务启动时自己注册到 consul</span>

* web8083服务及自动注册到consul
~~~go
package main

import (
	"fmt"

	"time"

	"net/http"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func main() {

	//注册服务
	c := NewConsulRegister()
	err := c.AgentServiceRegistration()
	if err != nil {
		fmt.Println("AgentServiceRegistration err:", err)
	}
	fmt.Println("AgentServiceRegistration success")

	//启动8083
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8083", nil)

}

var cnt = 0

func handler(w http.ResponseWriter, r *http.Request) {
	res := "hello consul--web2:" + strconv.Itoa(cnt)
	fmt.Println(res)
	cnt++
	fmt.Fprintf(w, res)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	res := "health check from consul--web2:" + strconv.Itoa(cnt)
	fmt.Println(res)
	cnt++

}

// NewConsulRegister create a new consul register
func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Address: "127.0.0.1:8500",
		Service: "web2",
		Health:  "health",
		Tag:     []string{"8083_test"},
		Port:    8083,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

// consul service register
type ConsulRegister struct {
	Address                        string
	Service                        string
	Health                         string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

//AgentServiceRegistration
func (r *ConsulRegister) AgentServiceRegistration() error {

	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	agent := client.Agent()
	ip := getLocalIP()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v", r.Service),
		Name:    fmt.Sprintf("%v", r.Service),
		Tags:    r.Tag, // tag，可以为空
		Port:    r.Port,
		Address: ip,
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval: r.Interval.String(), // 健康检查间隔
			HTTP:     fmt.Sprintf("http://%s:%d/%s", ip, r.Port, r.Health),

			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(), //过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}
//本地 测试直接返回127.0.0.1，线上则需去掉注释
func getLocalIP() string {
	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//	return ""
	//}
	//for _, address := range addrs {
	//	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			return ipnet.IP.String()
	//		}
	//	}
	//}
	return "127.0.0.1"
}

~~~



* 重启consul,8082和8083两个服务运行正常，同样也可以在 http://127.0.0.1:8500/ui/dc1/services 查看。

![consul](https://github.com/kgtom/back-end/blob/master/pic/consul.png)

~~~
# tom @ tom-pc in ~ [23:01:27]
$ consul agent -dev -config-dir=/consul.conf
 2018/12/17 23:01:31 [DEBUG] agent: Check "service:web" is passing
 2018/12/17 23:01:39 [DEBUG] agent: Check "service:web2" is passing
~~~

### <span id="3">三、服务发现</span>

#### <span id="31">第一种通过命令</span>
* web.service.consul：通过json配置文件注册服务，其中web:是json配置中的服务名称
~~~

# tom @ tom-pc in ~ [23:23:29]
$ dig @127.0.0.1 -p 8600 web.service.consul SRV

; <<>> DiG 9.10.6 <<>> @127.0.0.1 -p 8600 web.service.consul SRV
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 25904
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 3
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;web.service.consul.		IN	SRV

;; ANSWER SECTION:
web.service.consul.	0	IN	SRV	1 1 8082 tom-pc.node.dc1.consul.

;; ADDITIONAL SECTION:
tom-pc.node.dc1.consul.	0	IN	A	127.0.0.1
tom-pc.node.dc1.consul.	0	IN	TXT	"consul-network-segment="

;; Query time: 0 msec
;; SERVER: 127.0.0.1#8600(127.0.0.1)
;; WHEN: Mon Dec 17 23:23:46 CST 2018
;; MSG SIZE  rcvd: 141

~~~


* web2.service.consul SRV:通过http服务注册，web2为服务名称
~~~

# tom @ tom-pc in ~ [23:23:17]
$ dig @127.0.0.1 -p 8600 web2.service.consul SRV

; <<>> DiG 9.10.6 <<>> @127.0.0.1 -p 8600 web2.service.consul SRV
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 58388
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 3
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;web2.service.consul.		IN	SRV

;; ANSWER SECTION:
web2.service.consul.	0	IN	SRV	1 1 8083 tom-pc.node.dc1.consul.

;; ADDITIONAL SECTION:
tom-pc.node.dc1.consul.	0	IN	A	127.0.0.1
tom-pc.node.dc1.consul.	0	IN	TXT	"consul-network-segment="

;; Query time: 1 msec
;; SERVER: 127.0.0.1#8600(127.0.0.1)
;; WHEN: Mon Dec 17 23:23:28 CST 2018
;; MSG SIZE  rcvd: 142

~~~

####  <span id="32">第二种 通过代码</span>
~~~
package main

import (
	"fmt"
	"net"
	"strconv"

	"bufio"

	"github.com/hashicorp/consul/api"
)

func main() {

	//非默认情况下 需要设置实际的参数，例如Address
	//config := api.DefaultConfig()
	////config.Address = ""
	//client, err := api.NewClient(config)
	//if err != nil {
	//	fmt.Println("client err:", err)
	//}
	//默认情况下
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println("client err:", err)
	}
	services, err := client.Agent().Services()

	addrs := map[string]string{}
	for _, service := range services {
		addrs[net.JoinHostPort(service.Address, strconv.Itoa(service.Port))] = service.ID
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))
		if err != nil {
			fmt.Println("conn err:", err)
		}
		fmt.Println("remoteAddr conn:", conn.RemoteAddr())
		//test
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("get status：", status)
	}
	fmt.Println("------")
	for k, v := range addrs {
		fmt.Println("addr:", k, "id:", v)
	}
	//output:
	//addr: 127.0.0.1:8082 id: web
	//addr: 127.0.0.1:8083 id: web2

}
~~~


###  <span id="4">kv使用</span>

#### 第一种方式 通过代码
~~~
package main

import "github.com/hashicorp/consul/api"
import "fmt"

func main() {
	//  new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	// put
	p := &api.KVPair{Key: "k3", Value: []byte("333")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// lookup
	pair, _, err := kv.Get("k3", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("kv: %v %s\n", pair.Key, pair.Value)
}

~~~

#### 第二种 命令行方式
~~~
# tom @ tom-pc in ~ [21:45:19] 
$ curl -X PUT -d "bar" http://localhost:8500/v1/kv/foo
true

# tom @ tom-pc in ~ [21:45:25]
$ curl http://localhost:8500/v1/kv/foo
[
    {
        "LockIndex": 0,
        "Key": "foo",
        "Flags": 0,
        "Value": "YmFy",
        "CreateIndex": 19,
        "ModifyIndex": 19
    }
]

~~~



