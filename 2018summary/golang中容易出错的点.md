---


---

<h2 id="golang中容易出错的点">golang中容易出错的点</h2>
<h3 id="参数覆盖了函数功能试图在一个字符串上调用println函数">1.参数覆盖了函数功能,试图在一个字符串上调用Println函数</h3>
<pre class=" language-go"><code class="prism  language-go">
<span class="token keyword">func</span>  <span class="token function">foo</span><span class="token punctuation">(</span>v <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span> <span class="token punctuation">{</span>

fmt <span class="token operator">:=</span> <span class="token string">"bar"</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token comment">//编译报错</span>

<span class="token keyword">break</span>

<span class="token punctuation">}</span>

<span class="token punctuation">}</span>

</code></pre>
<h3 id="参数覆盖了struct结构体类型，-在go中没有类型可以作为值分配给变量。">2.参数覆盖了struct结构体类型， 在go中没有类型可以作为值分配给变量。</h3>
<pre class=" language-go"><code class="prism  language-go">
<span class="token keyword">type</span>  foo  <span class="token keyword">struct</span> <span class="token punctuation">{</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

foo <span class="token operator">:=</span> <span class="token operator">&amp;</span>foo<span class="token punctuation">{</span><span class="token punctuation">}</span>

foo <span class="token operator">=</span> <span class="token operator">&amp;</span>foo<span class="token punctuation">{</span><span class="token punctuation">}</span> <span class="token comment">//显示foo不是类型的错误消息</span>

<span class="token punctuation">}</span>

</code></pre>
<h3 id="myint-是int-类型别名">3. myint 是int 类型别名</h3>
<pre class=" language-go"><code class="prism  language-go">
<span class="token keyword">type</span>  myint  <span class="token builtin">int</span>

  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">var</span>  i myint <span class="token operator">=</span> <span class="token number">1</span>

<span class="token keyword">var</span>  j <span class="token builtin">int</span> <span class="token operator">=</span> <span class="token number">2</span>

<span class="token function">foo</span><span class="token punctuation">(</span>i<span class="token punctuation">)</span>

<span class="token function">foo</span><span class="token punctuation">(</span>j<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">foo</span><span class="token punctuation">(</span>i <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">switch</span>  v <span class="token operator">:=</span> i<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">case</span>  <span class="token builtin">int</span><span class="token punctuation">:</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"type of int :"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>

<span class="token keyword">case</span> myint<span class="token punctuation">:</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"type of myint:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token punctuation">}</span>

</code></pre>
<pre class=" language-go"><code class="prism  language-go">
<span class="token keyword">type</span>  myint <span class="token operator">=</span> <span class="token builtin">int</span>  <span class="token comment">//int和 myint 是完全相同的类型</span>

  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">var</span>  i myint <span class="token operator">=</span> <span class="token number">1</span>

<span class="token keyword">var</span>  j <span class="token builtin">int</span> <span class="token operator">=</span> <span class="token number">2</span>

<span class="token function">foo</span><span class="token punctuation">(</span>i<span class="token punctuation">)</span>

<span class="token function">foo</span><span class="token punctuation">(</span>j<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">foo</span><span class="token punctuation">(</span>i <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">switch</span>  v <span class="token operator">:=</span> i<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">case</span>  <span class="token builtin">int</span><span class="token punctuation">:</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"type of int :"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token keyword">switch</span>  v <span class="token operator">:=</span> i<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">case</span> myint<span class="token punctuation">:</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"type of myint:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token punctuation">}</span>

</code></pre>
<h3 id="通过类型别名，修改-no-type的bug.">4.通过类型别名，修改 "no type"的bug.</h3>
<pre class=" language-go"><code class="prism  language-go">
  

<span class="token keyword">type</span>  Foo  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

foo <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"foo"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token keyword">var</span>  foo2 Foo

foo2 <span class="token operator">=</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"foo2"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token function">test</span><span class="token punctuation">(</span>foo<span class="token punctuation">)</span>

<span class="token function">test</span><span class="token punctuation">(</span>foo2<span class="token punctuation">)</span>

<span class="token function">test2</span><span class="token punctuation">(</span>foo<span class="token punctuation">)</span>

<span class="token function">test2</span><span class="token punctuation">(</span>foo2<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">test</span><span class="token punctuation">(</span>foo Foo<span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token function">foo</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">test2</span><span class="token punctuation">(</span>i <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

  

<span class="token keyword">switch</span>  v <span class="token operator">:=</span> i<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">case</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>

<span class="token function">v</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token keyword">default</span><span class="token punctuation">:</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"no type"</span><span class="token punctuation">)</span> <span class="token comment">//foo2 明确是Foo类型，而非func()。这个bug，通过Go1.9 添加类型别名来处理。</span>

<span class="token punctuation">}</span>

<span class="token punctuation">}</span>

</code></pre>
<p>将 type Foo func() 修成type Foo = func() 就可以解决。</p>
<h3 id="闭包">闭包</h3>
<pre class=" language-go"><code class="prism  language-go">
  

f <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token number">3</span><span class="token punctuation">)</span>

<span class="token keyword">for</span>  i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token number">3</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

f<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"i:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token punctuation">}</span>

f<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

</code></pre>
<p>输出：3 3 3</p>
<pre class=" language-go"><code class="prism  language-go">
  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

  

f <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token number">3</span><span class="token punctuation">)</span>

<span class="token keyword">for</span>  i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token number">3</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

<span class="token comment">//创建一个临时函数，它将变量作为参数并返回闭包。我们立即调用它</span>

f<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>j <span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">return</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"i:"</span><span class="token punctuation">,</span> j<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token punctuation">}</span><span class="token punctuation">(</span>i<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

f<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token punctuation">}</span>

</code></pre>
<p>输出 0 1 2</p>
<pre class=" language-go"><code class="prism  language-go">
<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

  

f <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token number">3</span><span class="token punctuation">)</span>

<span class="token keyword">for</span>  i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token number">3</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

i <span class="token operator">:=</span> i <span class="token comment">//循环的每个过程中都会创建一个新变量（新指针）。</span>

f<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"i:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token punctuation">}</span>

f<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

f<span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token punctuation">}</span>

</code></pre>
<p>输出： 0 1 2</p>
<p>reference:</p>
<p><a href="https://jaxenter.com/5-things-you-hate-about-go-143422.html">https://jaxenter.com/5-things-you-hate-about-go-143422.html</a></p>

