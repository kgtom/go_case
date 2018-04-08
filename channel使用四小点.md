---


---

<h3 id="第一--从close的unbuffered-channel上执行读操作，返回channel对应类型的零值，但写入会panic。读写都不会堵塞。">第一  从close的unbuffered channel上执行读操作，返回channel对应类型的零值，但写入会panic。读写都不会堵塞。</h3>
<pre class=" language-go"><code class="prism  language-go">ch  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">bool</span><span class="token punctuation">)</span>
<span class="token function">close</span><span class="token punctuation">(</span>ch<span class="token punctuation">)</span>
v  <span class="token operator">:=</span>  <span class="token operator">&lt;-</span>ch
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span> <span class="token comment">//false</span>
v<span class="token punctuation">,</span> ok  <span class="token operator">:=</span>  <span class="token operator">&lt;-</span>ch
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> v<span class="token punctuation">,</span> <span class="token string">"ok:"</span><span class="token punctuation">,</span> ok<span class="token punctuation">)</span> <span class="token comment">//false,false</span>
ci  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span>
<span class="token function">close</span><span class="token punctuation">(</span>ci<span class="token punctuation">)</span>
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch int:"</span><span class="token punctuation">,</span> <span class="token operator">&lt;-</span>ci<span class="token punctuation">)</span> <span class="token comment">//0 channel对应类型的零值</span>
ci <span class="token operator">&lt;-</span>  <span class="token number">2</span>  <span class="token comment">//send on closed channel</span>
</code></pre>
<h3 id="第二-从close带buffered-channel可以读取，返回对应类型的零值，但写入会panic。读写都不会堵塞。">第二 从close带buffered channel,可以读取，返回对应类型的零值，但写入会panic。读写都不会堵塞。</h3>
<pre class=" language-go"><code class="prism  language-go">ch  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">,</span> <span class="token number">2</span><span class="token punctuation">)</span>
ch <span class="token operator">&lt;-</span>  <span class="token number">1</span>
ch <span class="token operator">&lt;-</span>  <span class="token number">2</span>
<span class="token function">close</span><span class="token punctuation">(</span>ch<span class="token punctuation">)</span>
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch:"</span><span class="token punctuation">,</span> <span class="token operator">&lt;-</span>ch<span class="token punctuation">)</span> <span class="token comment">//1</span>
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch:"</span><span class="token punctuation">,</span> <span class="token operator">&lt;-</span>ch<span class="token punctuation">)</span> <span class="token comment">//2</span>
v<span class="token punctuation">,</span> ok  <span class="token operator">:=</span>  <span class="token operator">&lt;-</span>ch
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> v<span class="token punctuation">,</span> <span class="token string">"ok:"</span><span class="token punctuation">,</span> ok<span class="token punctuation">)</span> <span class="token comment">//0,false channel对应类型的零值</span>
ch <span class="token operator">&lt;-</span>  <span class="token number">2</span>  <span class="token comment">//send on closed channel</span>
</code></pre>
<h3 id="第三-没有初始化channelnil读写都会堵塞">第三 没有初始化channel(nil)读写都会堵塞</h3>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">var</span>  ch  <span class="token keyword">chan</span>  <span class="token builtin">int</span>
ch <span class="token operator">&lt;-</span>  <span class="token number">1</span>  <span class="token comment">//fatal error: all goroutines are asleep </span>
– deadlock<span class="token operator">!</span>
</code></pre>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">var</span>  ch  <span class="token keyword">chan</span>  <span class="token builtin">int</span>
<span class="token operator">&lt;-</span>ch <span class="token comment">//fatal error: all goroutines are asleep </span>
– deadlock<span class="token operator">!</span>
</code></pre>
<h3 id="第四-nil-channel妙用">第四 nil channel妙用</h3>
<p>先看一个死循环的例子：本想依次输出 1 ，3 数字。</p>
<pre class=" language-go"><code class="prism  language-go">ch1  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span>
ch2  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span>

<span class="token keyword">go</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second <span class="token operator">*</span>  <span class="token number">1</span><span class="token punctuation">)</span>
ch1 <span class="token operator">&lt;-</span>  <span class="token number">1</span>
<span class="token function">close</span><span class="token punctuation">(</span>ch1<span class="token punctuation">)</span>
<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">go</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second <span class="token operator">*</span>  <span class="token number">2</span><span class="token punctuation">)</span>
ch2 <span class="token operator">&lt;-</span>  <span class="token number">3</span>
<span class="token function">close</span><span class="token punctuation">(</span>ch2<span class="token punctuation">)</span>
<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">for</span> <span class="token punctuation">{</span>

<span class="token keyword">select</span> <span class="token punctuation">{</span>
<span class="token keyword">case</span>  v  <span class="token operator">:=</span>  <span class="token operator">&lt;-</span>ch1<span class="token punctuation">:</span>
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch1:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>
<span class="token keyword">case</span>  v  <span class="token operator">:=</span>  <span class="token operator">&lt;-</span>ch2<span class="token punctuation">:</span>
fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch2:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre>
<p><code>输出结果： 1，0，0，0(无数0) 原因： 1.close(ch1)后获取到都是 int类型的零值 0； 2.case顺序执行</code></p>
<p>改正方案：</p>
<pre class=" language-go"><code class="prism  language-go">ch1 <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span> <span class="token builtin">int</span><span class="token punctuation">)</span>

ch2 <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span> <span class="token builtin">int</span><span class="token punctuation">)</span>

<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second <span class="token operator">*</span> <span class="token number">1</span><span class="token punctuation">)</span>

ch1 <span class="token operator">&lt;-</span> <span class="token number">1</span>

<span class="token function">close</span><span class="token punctuation">(</span>ch1<span class="token punctuation">)</span>

<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
  time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span>Second <span class="token operator">*</span> <span class="token number">2</span><span class="token punctuation">)</span>
  ch2 <span class="token operator">&lt;-</span> <span class="token number">3</span>
  <span class="token function">close</span><span class="token punctuation">(</span>ch2<span class="token punctuation">)</span>
 <span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">for</span> <span class="token punctuation">{</span>

<span class="token keyword">select</span> <span class="token punctuation">{</span>
<span class="token keyword">case</span> v<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>ch1<span class="token punctuation">:</span>
<span class="token keyword">if</span> ok <span class="token punctuation">{</span>
  fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch1:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>
<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
  ch1 <span class="token operator">=</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>
<span class="token keyword">case</span> v<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>ch2<span class="token punctuation">:</span>
<span class="token keyword">if</span> ok <span class="token punctuation">{</span>
 fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"ch2:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>
<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
 ch2 <span class="token operator">=</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">if</span> ch1 <span class="token operator">==</span> <span class="token boolean">nil</span> <span class="token operator">&amp;&amp;</span> ch2 <span class="token operator">==</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
   <span class="token keyword">break</span>
<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre>
<p><code>ch1=nil ,堵塞在case这个分支，然后ch2才有机会运行。</code></p>

