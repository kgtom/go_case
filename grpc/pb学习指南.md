---


---

<h1 id="前言">前言</h1>
<ul>
<li>习惯用 <code>Json、XML</code> 数据存储格式的你们，相信大多都没听过<code>Protocol Buffer</code></li>
<li><code>Protocol Buffer</code> 其实 是 <code>Google</code>出品的一种轻量 &amp; 高效的结构化数据存储格式，性能比 <code>Json、XML</code> 真的强！太！多！</li>
</ul>
<blockquote>
<p>由于 <code>Google</code>出品，我相信<code>Protocol Buffer</code>已经具备足够的吸引力</p>
</blockquote>
<ul>
<li>今天，我将献上一份全面 &amp; 详细的 <code>Protocol Buffer</code>攻略，含介绍、特点、具体使用、源码分析、序列化原理等等，希望您们会喜欢。</li>
</ul>
<hr>
<h1 id="目录">目录</h1>
<p><img src="https://user-gold-cdn.xitu.io/2018/5/14/1635c4ee3ab468d6?imageView2/0/w/1280/h/960/format/webp/ignore-error/1" alt="示意图"></p>
<hr>
<h1 id="定义">1. 定义</h1>
<p>一种 结构化数据 的数据存储格式（类似于 <code>XML、Json</code> ）</p>
<blockquote>
<ol>
<li><code>Google</code> 出品 （开源）</li>
<li><code>Protocol Buffer</code> 目前有两个版本：<code>proto2</code> 和 <code>proto3</code></li>
<li>因为<code>proto3</code> 还是beta 版，所以本次讲解是 <code>proto2</code></li>
</ol>
</blockquote>
<hr>
<h1 id="作用">2. 作用</h1>
<p>通过将 结构化的数据 进行 串行化（<strong>序列化</strong>），从而实现 <strong>数据存储 / RPC 数据交换</strong>的功能</p>
<blockquote>
<ol>
<li>序列化： 将 数据结构或对象 转换成 二进制串 的过程</li>
<li>反序列化：将在序列化过程中所生成的二进制串 转换成 数据结构或者对象 的过程</li>
</ol>
</blockquote>
<hr>
<h1 id="特点">3. 特点</h1>
<ul>
<li>对比于 常见的 <code>XML、Json</code> 数据存储格式，<code>Protocol Buffer</code>有如下特点：</li>
</ul>
<p><img src="https://user-gold-cdn.xitu.io/2018/5/14/1635c4ee3aa3b3cc?imageView2/0/w/1280/h/960/format/webp/ignore-error/1" alt="Protocol Buffer 特点"></p>
<hr>
<h1 id="应用场景">4. 应用场景</h1>
<p>传输数据量大 &amp; 网络环境不稳定 的<strong>数据存储、RPC 数据交换</strong> 的需求场景</p>
<blockquote>
<p>如 即时IM （QQ、微信）的需求场景</p>
</blockquote>
<hr>
<h1 id="总结">总结</h1>
<p>在 <strong>传输数据量较大</strong>的需求场景下，<code>Protocol Buffer</code>比<code>XML、Json</code> 更小、更快、使用 &amp; 维护更简单！</p>
<hr>
<h1 id="序列化原理解析">5. 序列化原理解析</h1>
<ul>
<li>序列化的本质：对数据进行编码 + 存储</li>
<li><code>Protocol Buffer</code>的性能好：传输效率快，主要原因 = <strong>序列化速度快 &amp; 序列化后的数据体积小</strong>，其原因如下：</li>
</ul>
<ol>
<li>
<p>序列化速度快的原因： a. 编码 / 解码 方式简单（只需要简单的数学运算 = 位移等等） b. 采用 <strong><code>PB</code> 自身的框架代码 和 编译器</strong> 共同完成</p>
</li>
<li>
<p>序列化后的数据量体积小（即数据压缩效果好）的原因： a. 采用了独特的编码方式，如<code>Varint</code>、<code>Zigzag</code>编码方式等等 b. 采用<code>T - L - V</code> 的数据存储方式：减少了分隔符的使用 &amp; 数据存储得紧凑</p>
</li>
</ol>
<p>更加详细的介绍，请看文章：<a href="https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fcarson_ho%2Farticle%2Fdetails%2F70568606">Protocol Buffer 序列化原理大揭秘 - 为什么Protocol Buffer性能这么好？</a></p>
<p>至此， 关于<code>Protocol Buffer</code>的序列化原理讲解完毕。下面将继续讲解如何具体使用<code>Protocol Buffer</code></p>
<hr>
<h1 id="使用步骤--实例讲解">6. 使用步骤 &amp; 实例讲解</h1>
<p>使用 <code>Protocol Buffer</code> 的流程如下：</p>
<p><img src="https://user-gold-cdn.xitu.io/2018/5/14/1635c4ee3aaaa87c?imageView2/0/w/1280/h/960/format/webp/ignore-error/1" alt="Protocol Buffer使用流程"></p>
<p>下面，我将对流程中的每个流程进行详细讲解。</p>
<h3 id="环境配置">6.1 环境配置</h3>
<ul>
<li>要使用<code>Protocol Buffer</code> ，需要先在电脑上安装<code>Protocol Buffer</code></li>
<li>具体请看文章：<a href="https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fcarson_ho%2Farticle%2Fdetails%2F70208514">手把手教你如何安装Protocol Buffer</a></li>
</ul>
<p>至此， <code>Protocol Buffer</code>已经安装完成。下面将讲解如何具体使用<code>Protocol Buffer</code></p>
<hr>
<h3 id="构建-protocol-buffer-消息对象模型">6.2 构建 <code>Protocol Buffer</code> 消息对象模型</h3>
<ul>
<li>构建步骤具体如下：</li>
</ul>
<p><img src="https://user-gold-cdn.xitu.io/2018/5/14/1635c4ee3a94d7c0?imageView2/0/w/1280/h/960/format/webp/ignore-error/1" alt="构建步骤"></p>
<ul>
<li>下面将通过一个实例（<code>Android（Java）</code> 平台为例）详细介绍每个步骤。</li>
<li>具体请看文章：<a href="https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fcarson_ho%2Farticle%2Fdetails%2F70267574">这是一份很有诚意的 Protocol Buffer 语法详解</a></li>
</ul>
<p>至此， 关于<code>Protocol Buffer</code>的语法 &amp; 如何构建<code>Protocol Buffer</code> 消息对象模型讲解完毕。下面将继续讲解如何具体使用<code>Protocol Buffer</code></p>
<blockquote>
<p>reference:<a href="https://juejin.im/post/5af8e9316fb9a07aab29f46d">addr:</a></p>
</blockquote>

