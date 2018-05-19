---


---

<h3 id="重点强调一下-性能测试结果：">重点强调一下 性能测试结果：</h3>
<blockquote>
<p>注意: 稍后，当我们的服务运行到Docker Swarm模式的Docker容器中时， 我们会在那里做所有基准测试并捕获度量。</p>
</blockquote>
<p>在开始负载测试之前，我们的基于Go的accountservice内存消耗可以从macbook的任务管理器中查看到，大概如下:</p>
<p><img src="https://segmentfault.com/img/bV921r?w=1472&amp;h=158" alt="clipboard.png" title="clipboard.png"></p>
<p>1.8MB， 不是特别坏。让我们使用Gatling测试，运行每秒1000个请求。需要记住一点，我们使用了非常幼稚的实现，我们仅仅响应一个硬编码的JSON响应。</p>
<p><img src="https://segmentfault.com/img/bV921t?w=1586&amp;h=156" alt="clipboard.png" title="clipboard.png"></p>
<p>服务每秒1000个请求，占用的内存也只是增加到28MB。 依然是Spring Boot应用程序启动时候使用内存的1/10. 当我们给它添加一些真正的功能时，看这些数字变化会更加有意思。</p>
<h3 id="性能和cpu使用率">性能和CPU使用率</h3>
<p><img src="https://segmentfault.com/img/bV921u?w=1590&amp;h=206" alt="clipboard.png" title="clipboard.png"></p>
<p>提供每秒1000个请求，每个核大概使用8%。</p>
<p><img src="https://segmentfault.com/img/bV921y?w=1994&amp;h=348" alt="clipboard.png" title="clipboard.png"></p>
<p>注意，Gatling一回合子微秒延迟如何， 但是平均延迟报告值为每个请求0ms, 花费庞大的11毫秒。 在这点上来看，我们的accountservice执行还是表现出色的，在子毫秒范围内大概每秒服务745个请求。</p>
<blockquote>
<p>Written with <a href="https://stackedit.io/">StackEdit</a>.</p>
</blockquote>

