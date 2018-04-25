---


---

<h2 id="api网关优点：">Api网关优点：</h2>
<ul>
<li>精细化api管理</li>
<li>所有api统一入口</li>
<li>权限验证</li>
<li>限流</li>
<li>负载均衡</li>
<li>数据聚合</li>
</ul>
<h2 id="api网关缺点：">Api网关缺点：</h2>
<ul>
<li>开发维护成本较高</li>
</ul>
<h2 id="推荐使用-gateway">推荐使用 gateway</h2>
<h3 id="深入了解-github--gateway">深入了解 <a href="https://github.com/fagongzi/gateway">github–gateway</a></h3>
<h2 id="环境搭建及使用">环境搭建及使用</h2>
<h3 id="docker-运行环境">1.docker 运行环境</h3>
<p>使用  <code>docker pull fagongzi/gateway</code>  命令下载Docker镜像, 使用  <code>docker run -d fagongzi/gateway</code>  运行镜像. 镜像启动后export 2个端口:</p>
<ul>
<li>
<p>80</p>
<p>Proxy的http端口，这个端口就是直接为终端用户服务的</p>
</li>
<li>
<p>9092</p>
<p>APIServer的对外GRPC的端口</p>
</li>
</ul>
<p><code>docker run -p 80:80 -p 9092:9092-d fagongzi/gateway</code></p>
<h2 id="举例：目前前端app-需要获取订单及库存数据。">2. 举例：目前前端app 需要获取订单及库存数据。</h2>
<p>订单api：192.168.0.10：8081</p>
<p>库存api：192.168.0.12：8082</p>
<p>docker 部署gateway:192.168.0.9</p>
<p>代码如下：</p>
<p>1.gateway有提供go的客户端，可以编程创建元数据.</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span>  <span class="token function">getClient</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>client<span class="token punctuation">.</span>Client<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span> client<span class="token punctuation">.</span><span class="token function">NewClient</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second<span class="token operator">*</span><span class="token number">10</span><span class="token punctuation">,</span>

<span class="token string">"192.168.0.9:9092"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre>
<p>2.创建两个cluster，一个给订单使用，另一个给库存使用。</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token comment">//创建 A订单，B库存 两个业务系统</span>

<span class="token keyword">func</span>  <span class="token function">createCluster</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>clusterAId<span class="token punctuation">,</span> clusterBId <span class="token builtin">uint64</span><span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

c<span class="token punctuation">,</span> err  <span class="token operator">:=</span>  <span class="token function">getClient</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span>  <span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">0</span><span class="token punctuation">,</span> err

<span class="token punctuation">}</span>

  

clusterAId<span class="token punctuation">,</span> err  <span class="token operator">=</span> c<span class="token punctuation">.</span><span class="token function">NewClusterBuilder</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Name</span><span class="token punctuation">(</span><span class="token string">"cluster-A"</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Loadbalance</span><span class="token punctuation">(</span>metapb<span class="token punctuation">.</span>RoundRobin<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Commit</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span>  <span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">0</span><span class="token punctuation">,</span> err

<span class="token punctuation">}</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"cluster-A id:"</span><span class="token punctuation">,</span> clusterAId<span class="token punctuation">)</span>

  

clusterBId<span class="token punctuation">,</span> err  <span class="token operator">=</span> c<span class="token punctuation">.</span><span class="token function">NewClusterBuilder</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Name</span><span class="token punctuation">(</span><span class="token string">"cluster-B"</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Loadbalance</span><span class="token punctuation">(</span>metapb<span class="token punctuation">.</span>RoundRobin<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Commit</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span>  <span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">0</span><span class="token punctuation">,</span> err

<span class="token punctuation">}</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"cluster-B id:"</span><span class="token punctuation">,</span> clusterBId<span class="token punctuation">)</span>

<span class="token keyword">return</span> clusterAId<span class="token punctuation">,</span> clusterBId<span class="token punctuation">,</span> <span class="token boolean">nil</span>

<span class="token punctuation">}</span>
</code></pre>
<p>2.创建两个server，真正部署 订单和库存api的服务器,<br>
需要绑定在clusterId.</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span>  <span class="token function">createServer</span><span class="token punctuation">(</span>clusterId <span class="token builtin">uint64</span><span class="token punctuation">,</span> addr<span class="token punctuation">,</span> checkPath <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>

c<span class="token punctuation">,</span> err  <span class="token operator">:=</span>  <span class="token function">getClient</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span> err

<span class="token punctuation">}</span>

  

sb  <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">NewServerBuilder</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token comment">// 必选项</span>

sb<span class="token punctuation">.</span><span class="token function">Addr</span><span class="token punctuation">(</span>addr<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">HTTPBackend</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">MaxQPS</span><span class="token punctuation">(</span><span class="token number">100</span><span class="token punctuation">)</span>

  

<span class="token comment">// 健康检查，可选项</span>

<span class="token comment">// 每个10秒钟检查一次，每次检查的超时时间30秒，即30秒后端Server没有返回认为后端不健康</span>

sb<span class="token punctuation">.</span><span class="token function">CheckHTTPCode</span><span class="token punctuation">(</span>checkPath<span class="token punctuation">,</span> time<span class="token punctuation">.</span>Second<span class="token operator">*</span><span class="token number">10</span><span class="token punctuation">,</span> time<span class="token punctuation">.</span>Second<span class="token operator">*</span><span class="token number">30</span><span class="token punctuation">)</span>

  

<span class="token comment">// 熔断器，可选项</span>

<span class="token comment">// 统计周期1秒钟</span>

sb<span class="token punctuation">.</span><span class="token function">CircuitBreakerCheckPeriod</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>

<span class="token comment">// 在Close状态60秒后自动转到Half状态</span>

sb<span class="token punctuation">.</span><span class="token function">CircuitBreakerCloseToHalfTimeout</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second <span class="token operator">*</span>  <span class="token number">60</span><span class="token punctuation">)</span>

<span class="token comment">// Half状态下，允许10%的流量流入后端</span>

sb<span class="token punctuation">.</span><span class="token function">CircuitBreakerHalfTrafficRate</span><span class="token punctuation">(</span><span class="token number">10</span><span class="token punctuation">)</span>

<span class="token comment">// 在Half状态，1秒内有2%的请求失败了，转换到Close状态</span>

sb<span class="token punctuation">.</span><span class="token function">CircuitBreakerHalfToCloseCondition</span><span class="token punctuation">(</span><span class="token number">2</span><span class="token punctuation">)</span>

<span class="token comment">// 在Half状态，1秒内有90%的请求成功了，转换到Open状态</span>

sb<span class="token punctuation">.</span><span class="token function">CircuitBreakerHalfToOpenCondition</span><span class="token punctuation">(</span><span class="token number">90</span><span class="token punctuation">)</span>

  

id<span class="token punctuation">,</span> err  <span class="token operator">:=</span> sb<span class="token punctuation">.</span><span class="token function">Commit</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span> err

<span class="token punctuation">}</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"server_A："</span><span class="token punctuation">,</span> id<span class="token punctuation">,</span> <span class="token string">"clusterA id:"</span><span class="token punctuation">,</span> clusterId<span class="token punctuation">,</span> <span class="token string">"checkPath:"</span><span class="token punctuation">,</span> checkPath<span class="token punctuation">)</span>

<span class="token comment">// 把这个server加入到cluster A、B</span>

c<span class="token punctuation">.</span><span class="token function">AddBind</span><span class="token punctuation">(</span>clusterId<span class="token punctuation">,</span> id<span class="token punctuation">)</span>

<span class="token keyword">return</span>  <span class="token boolean">nil</span>

<span class="token punctuation">}</span>
</code></pre>
<p>3.为 serverA及订单Api创建api规则，serverB雷同。</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span>  <span class="token function">createAPIForServerA</span><span class="token punctuation">(</span>cid <span class="token builtin">uint64</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>

c<span class="token punctuation">,</span> err  <span class="token operator">:=</span>  <span class="token function">getClient</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span> err

<span class="token punctuation">}</span>



sb  <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">NewAPIBuilder</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token comment">// 必选项</span>

sb<span class="token punctuation">.</span><span class="token function">Name</span><span class="token punctuation">(</span><span class="token string">"测试apiA"</span><span class="token punctuation">)</span>



<span class="token comment">// 设置URL规则，匹配所有开头为/api/user的请求</span>

sb<span class="token punctuation">.</span><span class="token function">MatchURLPattern</span><span class="token punctuation">(</span><span class="token string">"/api/v1/(.+)"</span><span class="token punctuation">)</span>

<span class="token comment">// 匹配GET请求</span>

sb<span class="token punctuation">.</span><span class="token function">MatchMethod</span><span class="token punctuation">(</span><span class="token string">"GET"</span><span class="token punctuation">)</span>

<span class="token comment">// 匹配所有请求</span>

sb<span class="token punctuation">.</span><span class="token function">MatchMethod</span><span class="token punctuation">(</span><span class="token string">"*"</span><span class="token punctuation">)</span>

<span class="token comment">// 不启动</span>

sb<span class="token punctuation">.</span><span class="token function">Down</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token comment">// 启用</span>

sb<span class="token punctuation">.</span><span class="token function">UP</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token comment">// 分发到Cluster A</span>

sb<span class="token punctuation">.</span><span class="token function">AddDispatchNode</span><span class="token punctuation">(</span>cid<span class="token punctuation">)</span>



<span class="token comment">// 可选项</span>

<span class="token comment">// 匹配所有host，和MatchMethod、MatchURLPattern互斥</span>

<span class="token comment">//sb.MatchDomain("user.xxx.com")</span>

<span class="token comment">// 增加访问黑名单</span>

<span class="token comment">//sb.AddBlacklist("192.168.0.1", "192.168.1.*", "192.168.*")</span>

<span class="token comment">// 增加访问报名单</span>

<span class="token comment">//sb.AddWhitelist("192.168.3.1", "192.168.3.*", "192.168.*")</span>

<span class="token comment">// 移除黑白名单</span>

<span class="token comment">//sb.RemoveBlacklist("192.168.0.1") // 剩余："192.168.1.*", "192.168.*"</span>

<span class="token comment">//sb.RemoveWhitelist("192.168.3.1") // 剩余："192.168.3.*", "192.168.*"</span>



<span class="token comment">// 增加默认值</span>

sb<span class="token punctuation">.</span><span class="token function">DefaultValue</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span><span class="token string">"{\"value\", \"default-a-test1\"}"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>

<span class="token comment">// 为默认值增加header</span>

sb<span class="token punctuation">.</span><span class="token function">AddDefaultValueHeader</span><span class="token punctuation">(</span><span class="token string">"token"</span><span class="token punctuation">,</span> <span class="token string">"token_xxxxx"</span><span class="token punctuation">)</span>

<span class="token comment">// 为默认值增加Cookie</span>

sb<span class="token punctuation">.</span><span class="token function">AddDefaultValueCookie</span><span class="token punctuation">(</span><span class="token string">"sid"</span><span class="token punctuation">,</span> <span class="token string">"xxxxx"</span><span class="token punctuation">)</span>



id<span class="token punctuation">,</span> err  <span class="token operator">:=</span> sb<span class="token punctuation">.</span><span class="token function">Commit</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"api commit err:"</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>

<span class="token keyword">return</span> err

<span class="token punctuation">}</span>



fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"api id :"</span><span class="token punctuation">,</span> id<span class="token punctuation">,</span> <span class="token string">"clusterA id:"</span><span class="token punctuation">,</span> cid<span class="token punctuation">)</span>

<span class="token keyword">return</span>  <span class="token boolean">nil</span>

<span class="token punctuation">}</span>
</code></pre>
<p>4.启动网卡</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token comment">// //创建Cluster A、B</span>

idA<span class="token punctuation">,</span> idB<span class="token punctuation">,</span> <span class="token boolean">_</span>  <span class="token operator">:=</span>  <span class="token function">createCluster</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token comment">//创建Server,</span>

<span class="token comment">//把 serverA:8081 和 Cluster-A 绑定</span>

<span class="token comment">//把serverB:8083 和 Cluster-B 绑定</span>

<span class="token function">createServer</span><span class="token punctuation">(</span>idA<span class="token punctuation">,</span> <span class="token string">"192.168.0.10:8081"</span><span class="token punctuation">,</span> <span class="token string">"/api/v1/check"</span><span class="token punctuation">)</span>

<span class="token function">createServer</span><span class="token punctuation">(</span>idB<span class="token punctuation">,</span> <span class="token string">"192.168.0.12:8082"</span><span class="token punctuation">,</span> <span class="token string">"/api/v2/check"</span><span class="token punctuation">)</span>

  

<span class="token comment">//创建API 转发到 Cluster-A</span>

err  <span class="token operator">:=</span>  <span class="token function">createAPIForServerA</span><span class="token punctuation">(</span>idA<span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"api a err:"</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token comment">//创建API 转发到 Cluster-B</span>

err  <span class="token operator">=</span>  <span class="token function">createAPIForServerB</span><span class="token punctuation">(</span>idB<span class="token punctuation">)</span>

<span class="token keyword">if</span> err <span class="token operator">!=</span>  <span class="token boolean">nil</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"api b err:"</span><span class="token punctuation">,</span> err<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  <span class="token punctuation">}</span>
</code></pre>
<p>5.测试</p>
<p>访问<br>
订单api: 192.168.0.9/api/v1/test1<br>
库存api: 192.168.0.9/api/v2/test1</p>
<p>6.总结 最重要精细化管理每一个对外的api。</p>

