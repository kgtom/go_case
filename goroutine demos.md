---


---

<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">package</span> main

  

<span class="token keyword">import</span> <span class="token punctuation">(</span>

<span class="token string">"fmt"</span>

<span class="token string">"math/rand"</span>

<span class="token string">"sync"</span>

<span class="token string">"sync/atomic"</span>

<span class="token string">"time"</span>

<span class="token punctuation">)</span>

  

<span class="token keyword">var</span>  wg sync<span class="token punctuation">.</span>WaitGroup

<span class="token keyword">var</span>  counter  <span class="token builtin">int64</span>

<span class="token keyword">var</span>  mux sync<span class="token punctuation">.</span>Mutex

  

<span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

  

ch  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">bool</span><span class="token punctuation">)</span>

a  <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span><span class="token string">"tom"</span><span class="token punctuation">,</span> <span class="token string">"lili"</span><span class="token punctuation">,</span> <span class="token string">"lucy"</span><span class="token punctuation">}</span>

<span class="token comment">// for _, v := range a {</span>

<span class="token comment">// go func(p string) {</span>

<span class="token comment">// fmt.Println("v:", p)</span>

<span class="token comment">// ch &lt;- true</span>

<span class="token comment">// }(v)</span>

  

<span class="token comment">// }</span>

  

<span class="token keyword">for</span>  <span class="token boolean">_</span><span class="token punctuation">,</span> v  <span class="token operator">:=</span>  <span class="token keyword">range</span> a <span class="token punctuation">{</span>

v  <span class="token operator">:=</span> v

<span class="token keyword">go</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"v:"</span><span class="token punctuation">,</span> v<span class="token punctuation">)</span>

ch <span class="token operator">&lt;-</span>  <span class="token boolean">true</span>

<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token punctuation">}</span>

<span class="token keyword">for</span>  <span class="token boolean">_</span>  <span class="token operator">=</span>  <span class="token keyword">range</span> a <span class="token punctuation">{</span>

<span class="token comment">//fmt.Println("ch:", &lt;-ch)</span>

<span class="token operator">&lt;-</span>ch

<span class="token punctuation">}</span>

  

<span class="token comment">// r := add()</span>

<span class="token comment">// for n := range sum(r) {</span>

<span class="token comment">// fmt.Println("ret:", n)</span>

<span class="token comment">// }</span>

  

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"end"</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token keyword">func</span>  <span class="token function">add</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token operator">&lt;-</span><span class="token keyword">chan</span>  <span class="token builtin">int</span> <span class="token punctuation">{</span>

out  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span>

<span class="token keyword">go</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">5</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

out <span class="token operator">&lt;-</span> i

<span class="token punctuation">}</span>

<span class="token function">close</span><span class="token punctuation">(</span>out<span class="token punctuation">)</span>

<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token keyword">return</span> out

<span class="token punctuation">}</span>

<span class="token keyword">func</span>  <span class="token function">sum</span><span class="token punctuation">(</span>i <span class="token operator">&lt;-</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token operator">&lt;-</span><span class="token keyword">chan</span>  <span class="token builtin">int</span> <span class="token punctuation">{</span>

s  <span class="token operator">:=</span>  <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span>  <span class="token builtin">int</span><span class="token punctuation">)</span>

<span class="token keyword">var</span>  sum  <span class="token builtin">int</span>

<span class="token keyword">go</span>  <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span>  n  <span class="token operator">:=</span>  <span class="token keyword">range</span> i <span class="token punctuation">{</span>

sum <span class="token operator">+=</span> n

<span class="token punctuation">}</span>

s <span class="token operator">&lt;-</span> sum

<span class="token function">close</span><span class="token punctuation">(</span>s<span class="token punctuation">)</span>

<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token keyword">return</span> s

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">GoTest</span><span class="token punctuation">(</span>ch <span class="token keyword">chan</span>  <span class="token builtin">bool</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"working"</span><span class="token punctuation">)</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">2</span>  <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"worked"</span><span class="token punctuation">)</span>

ch <span class="token operator">&lt;-</span>  <span class="token boolean">true</span>

<span class="token punctuation">}</span>

  

<span class="token keyword">func</span>  <span class="token function">atomicTest</span><span class="token punctuation">(</span>s <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span><span class="token function">Duration</span><span class="token punctuation">(</span>rand<span class="token punctuation">.</span><span class="token function">Intn</span><span class="token punctuation">(</span><span class="token number">3</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Millisecond<span class="token punctuation">)</span>

atomic<span class="token punctuation">.</span><span class="token function">AddInt64</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>counter<span class="token punctuation">,</span> <span class="token number">1</span><span class="token punctuation">)</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>s<span class="token punctuation">,</span> i<span class="token punctuation">,</span> <span class="token string">"counter:"</span><span class="token punctuation">,</span> atomic<span class="token punctuation">.</span><span class="token function">LoadInt64</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>counter<span class="token punctuation">)</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

wg<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token comment">//mutex</span>

<span class="token keyword">func</span>  <span class="token function">increater</span><span class="token punctuation">(</span>s <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

rand<span class="token punctuation">.</span><span class="token function">Seed</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span><span class="token function">Now</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">UnixNano</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

<span class="token comment">//x := counter</span>

mux<span class="token punctuation">.</span><span class="token function">Lock</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token comment">//x++</span>

counter<span class="token operator">++</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span>time<span class="token punctuation">.</span><span class="token function">Duration</span><span class="token punctuation">(</span>rand<span class="token punctuation">.</span><span class="token function">Intn</span><span class="token punctuation">(</span><span class="token number">3</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Millisecond<span class="token punctuation">)</span>

<span class="token comment">//counter = x</span>

mux<span class="token punctuation">.</span><span class="token function">Unlock</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>s<span class="token punctuation">,</span> i<span class="token punctuation">,</span> <span class="token string">"counter:"</span><span class="token punctuation">,</span> counter<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

wg<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token comment">//wg</span>

<span class="token keyword">func</span>  <span class="token function">foo</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"foo:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">3</span>  <span class="token operator">*</span> time<span class="token punctuation">.</span>Millisecond<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

wg<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>

  

<span class="token comment">//wg</span>

<span class="token keyword">func</span>  <span class="token function">bar</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">"bar:"</span><span class="token punctuation">,</span> i<span class="token punctuation">)</span>

time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">20</span>  <span class="token operator">*</span> time<span class="token punctuation">.</span>Millisecond<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

wg<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre>

