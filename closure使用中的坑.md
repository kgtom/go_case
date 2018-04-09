---


---

<h3 id="closure-的介绍">closure 的介绍</h3>
<hr>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"fmt"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	<span class="token keyword">var</span> i <span class="token builtin">int</span>

	<span class="token comment">//closure1</span>

	f <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">int</span> <span class="token punctuation">{</span>

		i<span class="token operator">++</span>
		x <span class="token operator">:=</span> i
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"x:"</span><span class="token punctuation">,</span> x<span class="token punctuation">)</span> <span class="token comment">//1</span>
		<span class="token keyword">return</span> x

	<span class="token punctuation">}</span>

	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"i:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span> <span class="token comment">//0 未调用f()</span>

	<span class="token comment">//fmt.Println("x2:", x) //closure内变量，访问不到。</span>

	<span class="token function">f</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token comment">//closure2</span>

	<span class="token punctuation">{</span>

		z <span class="token operator">:=</span> i
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"z:"</span><span class="token punctuation">,</span> z<span class="token punctuation">)</span> <span class="token comment">//1</span>

	<span class="token punctuation">}</span>

	<span class="token comment">//fmt.Println("z2:", z)//closure内变量，访问不到。</span>

	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"i2:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span> <span class="token comment">//1 调用f()后i++</span>

	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"end"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>


</code></pre>
<h3 id="closure-坑">closure 坑</h3>
<hr>
<pre class=" language-go"><code class="prism  language-go">  <span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	a <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span><span class="token string">"tom"</span><span class="token punctuation">,</span> <span class="token string">"lili"</span><span class="token punctuation">,</span> <span class="token string">"lucy"</span><span class="token punctuation">}</span>

	<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> v <span class="token operator">:=</span> <span class="token keyword">range</span> a <span class="token punctuation">{</span>
		<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>
		<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">1</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"end"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

</code></pre>
<p><code>输出 三个 lucy,原因：执行循环体，获取最后一次循环的值</code></p>
<p>改进方案一：</p>
<hr>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	a <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span><span class="token string">"tom"</span><span class="token punctuation">,</span> <span class="token string">"lili"</span><span class="token punctuation">,</span> <span class="token string">"lucy"</span><span class="token punctuation">}</span>

	<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> v <span class="token operator">:=</span> <span class="token keyword">range</span> a <span class="token punctuation">{</span>
		vv <span class="token operator">:=</span> v<span class="token comment">//拷贝一份新的</span>
		<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> vv<span class="token punctuation">)</span>
		<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">1</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"end"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre>
<p>改进方案二：</p>
<hr>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	a <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span><span class="token string">"tom"</span><span class="token punctuation">,</span> <span class="token string">"lili"</span><span class="token punctuation">,</span> <span class="token string">"lucy"</span><span class="token punctuation">}</span>

	<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> v <span class="token operator">:=</span> <span class="token keyword">range</span> a <span class="token punctuation">{</span>

		<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span>vv <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> vv<span class="token punctuation">)</span>
		<span class="token punctuation">}</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">1</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>
	fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"end"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre>

