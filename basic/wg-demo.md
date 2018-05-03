---


---

<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span>  <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

start  <span class="token operator">:=</span> time<span class="token punctuation">.</span><span class="token function">Now</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">UnixNano</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

<span class="token keyword">for</span>  i  <span class="token operator">:=</span>  <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span>  <span class="token number">3</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>

wg<span class="token punctuation">.</span><span class="token function">Add</span><span class="token punctuation">(</span><span class="token number">1</span><span class="token punctuation">)</span>

<span class="token comment">//调用cal()之前add,如果放在cal()后面，将"panic: sync: negative WaitGroup counter"</span>

<span class="token function">cal</span><span class="token punctuation">(</span>i<span class="token punctuation">)</span>

  

<span class="token punctuation">}</span>

wg<span class="token punctuation">.</span><span class="token function">Wait</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

end  <span class="token operator">:=</span> time<span class="token punctuation">.</span><span class="token function">Now</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">UnixNano</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

  

fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"end : const %d ms\n"</span><span class="token punctuation">,</span> end<span class="token operator">-</span>start<span class="token punctuation">)</span>

<span class="token punctuation">}</span>

<span class="token keyword">func</span>  <span class="token function">cal</span><span class="token punctuation">(</span>i <span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>i <span class="token operator">*</span>  <span class="token number">2</span><span class="token punctuation">)</span>

wg<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre>

