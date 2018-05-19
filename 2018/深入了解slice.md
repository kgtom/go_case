---


---

<p>切片是 Go 中的一种基本的数据结构，使用这种结构可以用来管理数据集合。切片的设计想法是由动态数组概念而来，为了开发者可以更加方便的使一个数据结构可以自动增加和减少。但是切片本身并不是动态数据或者数组指针。切片常见的操作有 reslice、append、copy。与此同时，切片还具有可索引，可迭代的优秀特性。</p>
<h2 id="一.-切片和数组">一. 切片和数组</h2>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_1.png" alt=""></p>
<p>关于切片和数组怎么选择？接下来好好讨论讨论这个问题。</p>
<p>在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    arrayA <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">100</span><span class="token punctuation">,</span> <span class="token number">200</span><span class="token punctuation">}</span>
    <span class="token keyword">var</span> arrayB <span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token builtin">int</span>

    arrayB <span class="token operator">=</span> arrayA

    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"arrayA : %p , %v\n"</span><span class="token punctuation">,</span> <span class="token operator">&amp;</span>arrayA<span class="token punctuation">,</span> arrayA<span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"arrayB : %p , %v\n"</span><span class="token punctuation">,</span> <span class="token operator">&amp;</span>arrayB<span class="token punctuation">,</span> arrayB<span class="token punctuation">)</span>

    <span class="token function">testArray</span><span class="token punctuation">(</span>arrayA<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">testArray</span><span class="token punctuation">(</span>x <span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"func Array : %p , %v\n"</span><span class="token punctuation">,</span> <span class="token operator">&amp;</span>x<span class="token punctuation">,</span> x<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>打印结果：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">arrayA <span class="token punctuation">:</span> <span class="token number">0xc4200bebf0</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">200</span><span class="token punctuation">]</span>  
arrayB <span class="token punctuation">:</span> <span class="token number">0xc4200bec00</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">200</span><span class="token punctuation">]</span>  
<span class="token keyword">func</span> Array <span class="token punctuation">:</span> <span class="token number">0xc4200bec30</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">200</span><span class="token punctuation">]</span>

</code></pre>
<p>可以看到，三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的。那这会导致什么问题呢？</p>
<p>假想每次传参都用数组，那么每次数组都要被复制一遍。如果数组大小有 100万，在64位机器上就需要花费大约 800W 字节，即 8MB 内存。这样会消耗掉大量的内存。于是乎有人想到，函数传参用数组的指针。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    arrayA <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token number">2</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">100</span><span class="token punctuation">,</span> <span class="token number">200</span><span class="token punctuation">}</span>
    <span class="token function">testArrayPoint</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>arrayA<span class="token punctuation">)</span>   <span class="token comment">// 1.传数组指针</span>
    arrayB <span class="token operator">:=</span> arrayA<span class="token punctuation">[</span><span class="token punctuation">:</span><span class="token punctuation">]</span>
    <span class="token function">testArrayPoint</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>arrayB<span class="token punctuation">)</span>   <span class="token comment">// 2.传切片</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"arrayA : %p , %v\n"</span><span class="token punctuation">,</span> <span class="token operator">&amp;</span>arrayA<span class="token punctuation">,</span> arrayA<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">testArrayPoint</span><span class="token punctuation">(</span>x <span class="token operator">*</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"func Array : %p , %v\n"</span><span class="token punctuation">,</span> x<span class="token punctuation">,</span> <span class="token operator">*</span>x<span class="token punctuation">)</span>
    <span class="token punctuation">(</span><span class="token operator">*</span>x<span class="token punctuation">)</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token operator">+=</span> <span class="token number">100</span>
<span class="token punctuation">}</span>

</code></pre>
<p>打印结果：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> Array <span class="token punctuation">:</span> <span class="token number">0xc4200b0140</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">200</span><span class="token punctuation">]</span>  
<span class="token keyword">func</span> Array <span class="token punctuation">:</span> <span class="token number">0xc4200b0180</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">300</span><span class="token punctuation">]</span>  
arrayA <span class="token punctuation">:</span> <span class="token number">0xc4200b0140</span> <span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token number">100</span> <span class="token number">400</span><span class="token punctuation">]</span>

</code></pre>
<p>这也就证明了数组指针确实到达了我们想要的效果。现在就算是传入10亿的数组，也只需要再栈上分配一个8个字节的内存给指针就可以了。这样更加高效的利用内存，性能也比之前的好。</p>
<p>不过传指针会有一个弊端，从打印结果可以看到，第一行和第三行指针地址都是同一个，万一原数组的指针指向更改了，那么函数里面的指针指向都会跟着更改。</p>
<p>切片的优势也就表现出来了。用切片传数组参数，既可以达到节约内存的目的，也可以达到合理处理好共享内存的问题。打印结果第二行就是切片，切片的指针和原来数组的指针是不同的。</p>
<p>由此我们可以得出结论：</p>
<p>把第一个大数组传递给函数会消耗很多内存，采用切片的方式传参可以避免上述问题。切片是引用传递，所以它们不需要使用额外的内存并且比使用数组更有效率。</p>
<p>但是，依旧有反例。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token string">"testing"</span>

<span class="token keyword">func</span> <span class="token function">array</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token number">1024</span><span class="token punctuation">]</span><span class="token builtin">int</span> <span class="token punctuation">{</span>  
    <span class="token keyword">var</span> x <span class="token punctuation">[</span><span class="token number">1024</span><span class="token punctuation">]</span><span class="token builtin">int</span>
    <span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token function">len</span><span class="token punctuation">(</span>x<span class="token punctuation">)</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
        x<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> i
    <span class="token punctuation">}</span>
    <span class="token keyword">return</span> x
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">slice</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span> <span class="token punctuation">{</span>  
    x <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">,</span> <span class="token number">1024</span><span class="token punctuation">)</span>
    <span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token function">len</span><span class="token punctuation">(</span>x<span class="token punctuation">)</span><span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
        x<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> i
    <span class="token punctuation">}</span>
    <span class="token keyword">return</span> x
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">BenchmarkArray</span><span class="token punctuation">(</span>b <span class="token operator">*</span>testing<span class="token punctuation">.</span>B<span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    <span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> b<span class="token punctuation">.</span>N<span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
        <span class="token function">array</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">BenchmarkSlice</span><span class="token punctuation">(</span>b <span class="token operator">*</span>testing<span class="token punctuation">.</span>B<span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    <span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> b<span class="token punctuation">.</span>N<span class="token punctuation">;</span> i<span class="token operator">++</span> <span class="token punctuation">{</span>
        <span class="token function">slice</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span>

</code></pre>
<p>我们做一次性能测试，并且禁用内联和优化，来观察切片的堆上内存分配的情况。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">  <span class="token keyword">go</span> test <span class="token operator">-</span>bench <span class="token punctuation">.</span> <span class="token operator">-</span>benchmem <span class="token operator">-</span>gcflags <span class="token string">"-N -l"</span>

</code></pre>
<p>输出结果比较“令人意外”：</p>
<p>vim</p>
<pre class=" language-vim"><code class="prism  language-vim">BenchmarkArray<span class="token operator">-</span><span class="token number">4</span>          <span class="token number">500000</span>              <span class="token number">3637</span> ns<span class="token operator">/</span>op               <span class="token number">0</span> B<span class="token operator">/</span>op          <span class="token number">0</span> alloc s<span class="token operator">/</span>op  
BenchmarkSlice<span class="token operator">-</span><span class="token number">4</span>          <span class="token number">300000</span>              <span class="token number">4055</span> ns<span class="token operator">/</span>op            <span class="token number">8192</span> B<span class="token operator">/</span>op          <span class="token number">1</span> alloc s<span class="token operator">/</span>op

</code></pre>
<p>解释一下上述结果，在测试 Array 的时候，用的是4核，循环次数是500000，平均每次执行时间是3637 ns，每次执行堆上分配内存总量是0，分配次数也是0 。</p>
<p>而切片的结果就“差”一点，同样也是用的是4核，循环次数是300000，平均每次执行时间是4055 ns，但是每次执行一次，堆上分配内存总量是8192，分配次数也是1 。</p>
<p>这样对比看来，并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，而且小数组在栈上拷贝的消耗也未必比 make 消耗大。</p>
<h2 id="二.-切片的数据结构">二. 切片的数据结构</h2>
<p>切片本身并不是动态数组或者数组指针。它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。<strong>切片本身是一个只读对象，其工作机制类似数组指针的一种封装</strong>。</p>
<p>切片（slice）是对数组一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。切片提供了一个与指向数组的动态窗口。</p>
<p>给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个长度可变的数组。</p>
<p>Slice 的数据结构定义如下:</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">type</span> slice <span class="token keyword">struct</span> <span class="token punctuation">{</span>  
    array unsafe<span class="token punctuation">.</span>Pointer
    <span class="token builtin">len</span>   <span class="token builtin">int</span>
    <span class="token builtin">cap</span>   <span class="token builtin">int</span>
<span class="token punctuation">}</span>

</code></pre>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_2.png" alt=""></p>
<p>切片的结构体由3部分构成，Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。cap 总是大于等于 len 的。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_3.png" alt=""></p>
<p>如果想从 slice 中得到一块内存地址，可以这样做：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">s <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token number">200</span><span class="token punctuation">)</span>  
ptr <span class="token operator">:=</span> unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>s<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">)</span>

</code></pre>
<p>如果反过来呢？从 Go 的内存地址中构造一个 slice。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">var</span> ptr unsafe<span class="token punctuation">.</span>Pointer  
<span class="token keyword">var</span> s1 <span class="token operator">=</span> <span class="token keyword">struct</span> <span class="token punctuation">{</span>  
    addr <span class="token builtin">uintptr</span>
    <span class="token builtin">len</span> <span class="token builtin">int</span>
    <span class="token builtin">cap</span> <span class="token builtin">int</span>
<span class="token punctuation">}</span><span class="token punctuation">{</span>ptr<span class="token punctuation">,</span> length<span class="token punctuation">,</span> length<span class="token punctuation">}</span>
s <span class="token operator">:=</span> <span class="token operator">*</span><span class="token punctuation">(</span><span class="token operator">*</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>s1<span class="token punctuation">)</span><span class="token punctuation">)</span>

</code></pre>
<p>构造一个虚拟的结构体，把 slice 的数据结构拼出来。</p>
<p>当然还有更加直接的方法，在 Go 的反射中就存在一个与之对应的数据结构 SliceHeader，我们可以用它来构造一个 slice</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">var</span> o <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span>  
sliceHeader <span class="token operator">:=</span> <span class="token punctuation">(</span><span class="token operator">*</span>reflect<span class="token punctuation">.</span>SliceHeader<span class="token punctuation">)</span><span class="token punctuation">(</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>o<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>  
sliceHeader<span class="token punctuation">.</span>Cap <span class="token operator">=</span> length  
sliceHeader<span class="token punctuation">.</span>Len <span class="token operator">=</span> length  
sliceHeader<span class="token punctuation">.</span>Data <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>ptr<span class="token punctuation">)</span>

</code></pre>
<h2 id="三.-创建切片">三. 创建切片</h2>
<p>make 函数允许在运行期动态指定数组长度，绕开了数组类型必须使用编译期常量的限制。</p>
<p>创建切片有两种形式，make 创建切片，空切片。</p>
<h3 id="make-和切片字面量">1. make 和切片字面量</h3>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">makeslice</span><span class="token punctuation">(</span>et <span class="token operator">*</span>_type<span class="token punctuation">,</span> <span class="token builtin">len</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token builtin">int</span><span class="token punctuation">)</span> slice <span class="token punctuation">{</span>  
    <span class="token comment">// 根据切片的数据类型，获取切片的最大容量</span>
    maxElements <span class="token operator">:=</span> <span class="token function">maxSliceCap</span><span class="token punctuation">(</span>et<span class="token punctuation">.</span>size<span class="token punctuation">)</span>
    <span class="token comment">// 比较切片的长度，长度值域应该在[0,maxElements]之间</span>
    <span class="token keyword">if</span> <span class="token builtin">len</span> <span class="token operator">&lt;</span> <span class="token number">0</span> <span class="token operator">||</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">len</span><span class="token punctuation">)</span> <span class="token operator">&gt;</span> maxElements <span class="token punctuation">{</span>
        <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"makeslice: len out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 比较切片的容量，容量值域应该在[len,maxElements]之间</span>
    <span class="token keyword">if</span> <span class="token builtin">cap</span> <span class="token operator">&lt;</span> <span class="token builtin">len</span> <span class="token operator">||</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span> <span class="token operator">&gt;</span> maxElements <span class="token punctuation">{</span>
        <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"makeslice: cap out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 根据切片的容量申请内存</span>
    p <span class="token operator">:=</span> <span class="token function">mallocgc</span><span class="token punctuation">(</span>et<span class="token punctuation">.</span>size<span class="token operator">*</span><span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span><span class="token punctuation">,</span> et<span class="token punctuation">,</span> <span class="token boolean">true</span><span class="token punctuation">)</span>
    <span class="token comment">// 返回申请好内存的切片的首地址</span>
    <span class="token keyword">return</span> slice<span class="token punctuation">{</span>p<span class="token punctuation">,</span> <span class="token builtin">len</span><span class="token punctuation">,</span> <span class="token builtin">cap</span><span class="token punctuation">}</span>
<span class="token punctuation">}</span>

</code></pre>
<p>还有一个 int64 的版本：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">makeslice64</span><span class="token punctuation">(</span>et <span class="token operator">*</span>_type<span class="token punctuation">,</span> len64<span class="token punctuation">,</span> cap64 <span class="token builtin">int64</span><span class="token punctuation">)</span> slice <span class="token punctuation">{</span>  
    <span class="token builtin">len</span> <span class="token operator">:=</span> <span class="token function">int</span><span class="token punctuation">(</span>len64<span class="token punctuation">)</span>
    <span class="token keyword">if</span> <span class="token function">int64</span><span class="token punctuation">(</span><span class="token builtin">len</span><span class="token punctuation">)</span> <span class="token operator">!=</span> len64 <span class="token punctuation">{</span>
        <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"makeslice: len out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token builtin">cap</span> <span class="token operator">:=</span> <span class="token function">int</span><span class="token punctuation">(</span>cap64<span class="token punctuation">)</span>
    <span class="token keyword">if</span> <span class="token function">int64</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span> <span class="token operator">!=</span> cap64 <span class="token punctuation">{</span>
        <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"makeslice: cap out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token keyword">return</span> <span class="token function">makeslice</span><span class="token punctuation">(</span>et<span class="token punctuation">,</span> <span class="token builtin">len</span><span class="token punctuation">,</span> <span class="token builtin">cap</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>实现原理和上面的是一样的，只不过多了把 int64 转换成 int 这一步罢了。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_4.png" alt=""></p>
<p>上图是用 make 函数创建的一个 len = 4， cap = 6 的切片。内存空间申请了6个 int 类型的内存大小。由于 len = 4，所以后面2个暂时访问不到，但是容量还是在的。这时候数组里面每个变量都是0 。</p>
<p>除了 make 函数可以创建切片以外，字面量也可以创建切片。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_5.png" alt=""></p>
<p>这里是用字面量创建的一个 len = 6，cap = 6 的切片，这时候数组里面每个元素的值都初始化完成了。<strong>需要注意的是 [ ] 里面不要写数组的容量，因为如果写了个数以后就是数组了，而不是切片了。</strong></p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_6.png" alt=""></p>
<p>还有一种简单的字面量创建切片的方法。如上图。上图就 Slice A 创建出了一个 len = 3，cap = 3 的切片。从原数组的第二位元素(0是第一位)开始切，一直切到第四位为止(不包括第五位)。同理，Slice B 创建出了一个 len = 2，cap = 4 的切片。</p>
<h3 id="nil-和空切片">2. nil 和空切片</h3>
<p>nil 切片和空切片也是常用的。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">var</span> slice <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span>

</code></pre>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_7.png" alt=""></p>
<p>nil 切片被用在很多标准库和内置函数中，描述一个不存在的切片的时候，就需要用到 nil 切片。比如函数在发生异常的时候，返回的切片就是 nil 切片。nil 切片的指针指向 nil。</p>
<p>空切片一般会用来表示一个空的集合。比如数据库查询，一条结果也没有查到，那么就可以返回一个空切片。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">silce <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span> <span class="token punctuation">,</span> <span class="token number">0</span> <span class="token punctuation">)</span>  
slice <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span> <span class="token punctuation">}</span>

</code></pre>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_8.png" alt=""></p>
<p>空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。</p>
<p>最后需要说明的一点是。不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的。</p>
<h2 id="四.-切片扩容">四. 切片扩容</h2>
<p>当一个切片的容量满了，就需要扩容了。怎么扩，策略是什么？</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">growslice</span><span class="token punctuation">(</span>et <span class="token operator">*</span>_type<span class="token punctuation">,</span> old slice<span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token builtin">int</span><span class="token punctuation">)</span> slice <span class="token punctuation">{</span>  
    <span class="token keyword">if</span> raceenabled <span class="token punctuation">{</span>
        callerpc <span class="token operator">:=</span> <span class="token function">getcallerpc</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>et<span class="token punctuation">)</span><span class="token punctuation">)</span>
        <span class="token function">racereadrangepc</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>et<span class="token punctuation">.</span>size<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> callerpc<span class="token punctuation">,</span> <span class="token function">funcPC</span><span class="token punctuation">(</span>growslice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token keyword">if</span> msanenabled <span class="token punctuation">{</span>
        <span class="token function">msanread</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>et<span class="token punctuation">.</span>size<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token keyword">if</span> et<span class="token punctuation">.</span>size <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
        <span class="token comment">// 如果新要扩容的容量比原来的容量还要小，这代表要缩容了，那么可以直接报panic了。</span>
        <span class="token keyword">if</span> <span class="token builtin">cap</span> <span class="token operator">&lt;</span> old<span class="token punctuation">.</span><span class="token builtin">cap</span> <span class="token punctuation">{</span>
            <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"growslice: cap out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
        <span class="token punctuation">}</span>

        <span class="token comment">// 如果当前切片的大小为0，还调用了扩容方法，那么就新生成一个新的容量的切片返回。</span>
        <span class="token keyword">return</span> slice<span class="token punctuation">{</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>zerobase<span class="token punctuation">)</span><span class="token punctuation">,</span> old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token punctuation">,</span> <span class="token builtin">cap</span><span class="token punctuation">}</span>
    <span class="token punctuation">}</span>

  <span class="token comment">// 这里就是扩容的策略</span>
    newcap <span class="token operator">:=</span> old<span class="token punctuation">.</span><span class="token builtin">cap</span>
    doublecap <span class="token operator">:=</span> newcap <span class="token operator">+</span> newcap
    <span class="token keyword">if</span> <span class="token builtin">cap</span> <span class="token operator">&gt;</span> doublecap <span class="token punctuation">{</span>
        newcap <span class="token operator">=</span> <span class="token builtin">cap</span>
    <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
        <span class="token keyword">if</span> old<span class="token punctuation">.</span><span class="token builtin">len</span> <span class="token operator">&lt;</span> <span class="token number">1024</span> <span class="token punctuation">{</span>
            newcap <span class="token operator">=</span> doublecap
        <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
            <span class="token keyword">for</span> newcap <span class="token operator">&lt;</span> <span class="token builtin">cap</span> <span class="token punctuation">{</span>
                newcap <span class="token operator">+=</span> newcap <span class="token operator">/</span> <span class="token number">4</span>
            <span class="token punctuation">}</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>

    <span class="token comment">// 计算新的切片的容量，长度。</span>
    <span class="token keyword">var</span> lenmem<span class="token punctuation">,</span> newlenmem<span class="token punctuation">,</span> capmem <span class="token builtin">uintptr</span>
    <span class="token keyword">const</span> ptrSize <span class="token operator">=</span> unsafe<span class="token punctuation">.</span><span class="token function">Sizeof</span><span class="token punctuation">(</span><span class="token punctuation">(</span><span class="token operator">*</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">(</span><span class="token boolean">nil</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token keyword">switch</span> et<span class="token punctuation">.</span>size <span class="token punctuation">{</span>
    <span class="token keyword">case</span> <span class="token number">1</span><span class="token punctuation">:</span>
        lenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token punctuation">)</span>
        newlenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span>
        capmem <span class="token operator">=</span> <span class="token function">roundupsize</span><span class="token punctuation">(</span><span class="token function">uintptr</span><span class="token punctuation">(</span>newcap<span class="token punctuation">)</span><span class="token punctuation">)</span>
        newcap <span class="token operator">=</span> <span class="token function">int</span><span class="token punctuation">(</span>capmem<span class="token punctuation">)</span>
    <span class="token keyword">case</span> ptrSize<span class="token punctuation">:</span>
        lenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token punctuation">)</span> <span class="token operator">*</span> ptrSize
        newlenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span> <span class="token operator">*</span> ptrSize
        capmem <span class="token operator">=</span> <span class="token function">roundupsize</span><span class="token punctuation">(</span><span class="token function">uintptr</span><span class="token punctuation">(</span>newcap<span class="token punctuation">)</span> <span class="token operator">*</span> ptrSize<span class="token punctuation">)</span>
        newcap <span class="token operator">=</span> <span class="token function">int</span><span class="token punctuation">(</span>capmem <span class="token operator">/</span> ptrSize<span class="token punctuation">)</span>
    <span class="token keyword">default</span><span class="token punctuation">:</span>
        lenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token punctuation">)</span> <span class="token operator">*</span> et<span class="token punctuation">.</span>size
        newlenmem <span class="token operator">=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token builtin">cap</span><span class="token punctuation">)</span> <span class="token operator">*</span> et<span class="token punctuation">.</span>size
        capmem <span class="token operator">=</span> <span class="token function">roundupsize</span><span class="token punctuation">(</span><span class="token function">uintptr</span><span class="token punctuation">(</span>newcap<span class="token punctuation">)</span> <span class="token operator">*</span> et<span class="token punctuation">.</span>size<span class="token punctuation">)</span>
        newcap <span class="token operator">=</span> <span class="token function">int</span><span class="token punctuation">(</span>capmem <span class="token operator">/</span> et<span class="token punctuation">.</span>size<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token comment">// 判断非法的值，保证容量是在增加，并且容量不超过最大容量</span>
    <span class="token keyword">if</span> <span class="token builtin">cap</span> <span class="token operator">&lt;</span> old<span class="token punctuation">.</span><span class="token builtin">cap</span> <span class="token operator">||</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>newcap<span class="token punctuation">)</span> <span class="token operator">&gt;</span> <span class="token function">maxSliceCap</span><span class="token punctuation">(</span>et<span class="token punctuation">.</span>size<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        <span class="token function">panic</span><span class="token punctuation">(</span><span class="token function">errorString</span><span class="token punctuation">(</span><span class="token string">"growslice: cap out of range"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token keyword">var</span> p unsafe<span class="token punctuation">.</span>Pointer
    <span class="token keyword">if</span> et<span class="token punctuation">.</span>kind<span class="token operator">&amp;</span>kindNoPointers <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
        <span class="token comment">// 在老的切片后面继续扩充容量</span>
        p <span class="token operator">=</span> <span class="token function">mallocgc</span><span class="token punctuation">(</span>capmem<span class="token punctuation">,</span> <span class="token boolean">nil</span><span class="token punctuation">,</span> <span class="token boolean">false</span><span class="token punctuation">)</span>
        <span class="token comment">// 将 lenmem 这个多个 bytes 从 old.array地址 拷贝到 p 的地址处</span>
        <span class="token function">memmove</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> old<span class="token punctuation">.</span>array<span class="token punctuation">,</span> lenmem<span class="token punctuation">)</span>
        <span class="token comment">// 先将 P 地址加上新的容量得到新切片容量的地址，然后将新切片容量地址后面的 capmem-newlenmem 个 bytes 这块内存初始化。为之后继续 append() 操作腾出空间。</span>
        <span class="token function">memclrNoHeapPointers</span><span class="token punctuation">(</span><span class="token function">add</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> newlenmem<span class="token punctuation">)</span><span class="token punctuation">,</span> capmem<span class="token operator">-</span>newlenmem<span class="token punctuation">)</span>
    <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
        <span class="token comment">// 重新申请新的数组给新切片</span>
        <span class="token comment">// 重新申请 capmen 这个大的内存地址，并且初始化为0值</span>
        p <span class="token operator">=</span> <span class="token function">mallocgc</span><span class="token punctuation">(</span>capmem<span class="token punctuation">,</span> et<span class="token punctuation">,</span> <span class="token boolean">true</span><span class="token punctuation">)</span>
        <span class="token keyword">if</span> <span class="token operator">!</span>writeBarrier<span class="token punctuation">.</span>enabled <span class="token punctuation">{</span>
            <span class="token comment">// 如果还不能打开写锁，那么只能把 lenmem 大小的 bytes 字节从 old.array 拷贝到 p 的地址处</span>
            <span class="token function">memmove</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> old<span class="token punctuation">.</span>array<span class="token punctuation">,</span> lenmem<span class="token punctuation">)</span>
        <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
            <span class="token comment">// 循环拷贝老的切片的值</span>
            <span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span><span class="token number">0</span><span class="token punctuation">)</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> lenmem<span class="token punctuation">;</span> i <span class="token operator">+=</span> et<span class="token punctuation">.</span>size <span class="token punctuation">{</span>
                <span class="token function">typedmemmove</span><span class="token punctuation">(</span>et<span class="token punctuation">,</span> <span class="token function">add</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> i<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">add</span><span class="token punctuation">(</span>old<span class="token punctuation">.</span>array<span class="token punctuation">,</span> i<span class="token punctuation">)</span><span class="token punctuation">)</span>
            <span class="token punctuation">}</span>
        <span class="token punctuation">}</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 返回最终新切片，容量更新为最新扩容之后的容量</span>
    <span class="token keyword">return</span> slice<span class="token punctuation">{</span>p<span class="token punctuation">,</span> old<span class="token punctuation">.</span><span class="token builtin">len</span><span class="token punctuation">,</span> newcap<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

</code></pre>
<p>上述就是扩容的实现。主要需要关注的有两点，一个是扩容时候的策略，还有一个就是扩容是生成全新的内存地址还是在原来的地址后追加。</p>
<h4 id="扩容策略">1. 扩容策略</h4>
<p>先看看扩容策略。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    slice <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">10</span><span class="token punctuation">,</span> <span class="token number">20</span><span class="token punctuation">,</span> <span class="token number">30</span><span class="token punctuation">,</span> <span class="token number">40</span><span class="token punctuation">}</span>
    newSlice <span class="token operator">:=</span> <span class="token function">append</span><span class="token punctuation">(</span>slice<span class="token punctuation">,</span> <span class="token number">50</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"Before slice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> slice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>slice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> newSlice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>newSlice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    newSlice<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token operator">+=</span> <span class="token number">10</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"After slice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> slice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>slice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"After newSlice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> newSlice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>newSlice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>输出结果：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">Before slice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">20</span> <span class="token number">30</span> <span class="token number">40</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200b0140</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">4</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
Before newSlice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">20</span> <span class="token number">30</span> <span class="token number">40</span> <span class="token number">50</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200b0180</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">5</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">8</span>  
After slice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">20</span> <span class="token number">30</span> <span class="token number">40</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200b0140</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">4</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
After newSlice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">30</span> <span class="token number">30</span> <span class="token number">40</span> <span class="token number">50</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200b0180</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">5</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">8</span>

</code></pre>
<p>用图表示出上述过程。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_9.png" alt=""></p>
<p>从图上我们可以很容易的看出，新的切片和之前的切片已经不同了，因为新的切片更改了一个值，并没有影响到原来的数组，新切片指向的数组是一个全新的数组。并且 cap 容量也发生了变化。这之间究竟发生了什么呢？</p>
<p>Go 中切片扩容的策略是这样的：</p>
<p>如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。</p>
<p>一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。</p>
<p><strong>注意：扩容扩大的容量都是针对原来的容量而言的，而不是针对原来数组的长度而言的。</strong></p>
<h4 id="新数组-or-老数组-？">2. 新数组 or 老数组 ？</h4>
<p>再谈谈扩容之后的数组一定是新的么？这个不一定，分两种情况。</p>
<p>情况一：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    array <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token number">4</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">10</span><span class="token punctuation">,</span> <span class="token number">20</span><span class="token punctuation">,</span> <span class="token number">30</span><span class="token punctuation">,</span> <span class="token number">40</span><span class="token punctuation">}</span>
    slice <span class="token operator">:=</span> array<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">:</span><span class="token number">2</span><span class="token punctuation">]</span>
    newSlice <span class="token operator">:=</span> <span class="token function">append</span><span class="token punctuation">(</span>slice<span class="token punctuation">,</span> <span class="token number">50</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"Before slice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> slice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>slice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> newSlice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>newSlice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    newSlice<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token operator">+=</span> <span class="token number">10</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"After slice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> slice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>slice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>slice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"After newSlice = %v, Pointer = %p, len = %d, cap = %d\n"</span><span class="token punctuation">,</span> newSlice<span class="token punctuation">,</span> <span class="token operator">&amp;</span>newSlice<span class="token punctuation">,</span> <span class="token function">len</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">cap</span><span class="token punctuation">(</span>newSlice<span class="token punctuation">)</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"After array = %v\n"</span><span class="token punctuation">,</span> array<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>打印输出：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">Before slice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">20</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200c0040</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">2</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
Before newSlice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">20</span> <span class="token number">50</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200c0060</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">3</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
After slice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">30</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200c0040</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">2</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
After newSlice <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">30</span> <span class="token number">50</span><span class="token punctuation">]</span><span class="token punctuation">,</span> Pointer <span class="token operator">=</span> <span class="token number">0xc4200c0060</span><span class="token punctuation">,</span> <span class="token builtin">len</span> <span class="token operator">=</span> <span class="token number">3</span><span class="token punctuation">,</span> <span class="token builtin">cap</span> <span class="token operator">=</span> <span class="token number">4</span>  
After array <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token number">10</span> <span class="token number">30</span> <span class="token number">50</span> <span class="token number">40</span><span class="token punctuation">]</span>

</code></pre>
<p>把上述过程用图表示出来，如下图。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_10.png" alt=""></p>
<p>通过打印的结果，我们可以看到，在这种情况下，扩容以后并没有新建一个新的数组，扩容前后的数组都是同一个，这也就导致了新的切片修改了一个值，也影响到了老的切片了。并且 append() 操作也改变了原来数组里面的值。一个 append() 操作影响了这么多地方，如果原数组上有多个切片，那么这些切片都会被影响！无意间就产生了莫名的 bug！</p>
<p>这种情况，由于原数组还有容量可以扩容，所以执行 append() 操作以后，会在原数组上直接操作，所以这种情况下，扩容以后的数组还是指向原来的数组。</p>
<p>这种情况也极容易出现在字面量创建切片时候，第三个参数 cap 传值的时候，如果用字面量创建切片，cap 并不等于指向数组的总容量，那么这种情况就会发生。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">slice <span class="token operator">:=</span> array<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token number">2</span><span class="token punctuation">:</span><span class="token number">3</span><span class="token punctuation">]</span>

</code></pre>
<p><strong>上面这种情况非常危险，极度容易产生 bug 。</strong></p>
<p>建议用字面量创建切片的时候，cap 的值一定要保持清醒，避免共享原数组导致的 bug。</p>
<p>情况二：</p>
<p>情况二其实就是在扩容策略里面举的例子，在那个例子中之所以生成了新的切片，是因为原来数组的容量已经达到了最大值，再想扩容， Go 默认会先开一片内存区域，把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组。</p>
<p>所以建议尽量避免情况一，尽量使用情况二，避免 bug 产生。</p>
<h2 id="五.-切片拷贝">五. 切片拷贝</h2>
<p>Slice 中拷贝方法有2个。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">slicecopy</span><span class="token punctuation">(</span>to<span class="token punctuation">,</span> fm slice<span class="token punctuation">,</span> width <span class="token builtin">uintptr</span><span class="token punctuation">)</span> <span class="token builtin">int</span> <span class="token punctuation">{</span>  
    <span class="token comment">// 如果源切片或者目标切片有一个长度为0，那么就不需要拷贝，直接 return </span>
    <span class="token keyword">if</span> fm<span class="token punctuation">.</span><span class="token builtin">len</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token operator">||</span> to<span class="token punctuation">.</span><span class="token builtin">len</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
        <span class="token keyword">return</span> <span class="token number">0</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// n 记录下源切片或者目标切片较短的那一个的长度</span>
    n <span class="token operator">:=</span> fm<span class="token punctuation">.</span><span class="token builtin">len</span>
    <span class="token keyword">if</span> to<span class="token punctuation">.</span><span class="token builtin">len</span> <span class="token operator">&lt;</span> n <span class="token punctuation">{</span>
        n <span class="token operator">=</span> to<span class="token punctuation">.</span><span class="token builtin">len</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 如果入参 width = 0，也不需要拷贝了，返回较短的切片的长度</span>
    <span class="token keyword">if</span> width <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
        <span class="token keyword">return</span> n
    <span class="token punctuation">}</span>
    <span class="token comment">// 如果开启了竞争检测</span>
    <span class="token keyword">if</span> raceenabled <span class="token punctuation">{</span>
        callerpc <span class="token operator">:=</span> <span class="token function">getcallerpc</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>to<span class="token punctuation">)</span><span class="token punctuation">)</span>
        pc <span class="token operator">:=</span> <span class="token function">funcPC</span><span class="token punctuation">(</span>slicecopy<span class="token punctuation">)</span>
        <span class="token function">racewriterangepc</span><span class="token punctuation">(</span>to<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>width<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> callerpc<span class="token punctuation">,</span> pc<span class="token punctuation">)</span>
        <span class="token function">racereadrangepc</span><span class="token punctuation">(</span>fm<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>width<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> callerpc<span class="token punctuation">,</span> pc<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 如果开启了 The memory sanitizer (msan)</span>
    <span class="token keyword">if</span> msanenabled <span class="token punctuation">{</span>
        <span class="token function">msanwrite</span><span class="token punctuation">(</span>to<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>width<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
        <span class="token function">msanread</span><span class="token punctuation">(</span>fm<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token operator">*</span><span class="token function">int</span><span class="token punctuation">(</span>width<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    size <span class="token operator">:=</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token punctuation">)</span> <span class="token operator">*</span> width
    <span class="token keyword">if</span> size <span class="token operator">==</span> <span class="token number">1</span> <span class="token punctuation">{</span> 
        <span class="token comment">// TODO: is this still worth it with new memmove impl?</span>
        <span class="token comment">// 如果只有一个元素，那么指针直接转换即可</span>
        <span class="token operator">*</span><span class="token punctuation">(</span><span class="token operator">*</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">(</span>to<span class="token punctuation">.</span>array<span class="token punctuation">)</span> <span class="token operator">=</span> <span class="token operator">*</span><span class="token punctuation">(</span><span class="token operator">*</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">(</span>fm<span class="token punctuation">.</span>array<span class="token punctuation">)</span> <span class="token comment">// known to be a byte pointer</span>
    <span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
        <span class="token comment">// 如果不止一个元素，那么就把 size 个 bytes 从 fm.array 地址开始，拷贝到 to.array 地址之后</span>
        <span class="token function">memmove</span><span class="token punctuation">(</span>to<span class="token punctuation">.</span>array<span class="token punctuation">,</span> fm<span class="token punctuation">.</span>array<span class="token punctuation">,</span> size<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token keyword">return</span> n
<span class="token punctuation">}</span>

</code></pre>
<p>在这个方法中，slicecopy 方法会把源切片值(即 fm Slice )中的元素复制到目标切片(即 to Slice )中，并返回被复制的元素个数，copy 的两个类型必须一致。slicecopy 方法最终的复制结果取决于较短的那个切片，当较短的切片复制完成，整个复制过程就全部完成了。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_11.png" alt=""></p>
<p>举个例子，比如：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    array <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">10</span><span class="token punctuation">,</span> <span class="token number">20</span><span class="token punctuation">,</span> <span class="token number">30</span><span class="token punctuation">,</span> <span class="token number">40</span><span class="token punctuation">}</span>
    slice <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">,</span> <span class="token number">6</span><span class="token punctuation">)</span>
    n <span class="token operator">:=</span> <span class="token function">copy</span><span class="token punctuation">(</span>slice<span class="token punctuation">,</span> array<span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>n<span class="token punctuation">,</span>slice<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>还有一个拷贝的方法，这个方法原理和 slicecopy 方法类似，不在赘述了，注释写在代码里面了。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">slicestringcopy</span><span class="token punctuation">(</span>to <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> fm <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">int</span> <span class="token punctuation">{</span>  
    <span class="token comment">// 如果源切片或者目标切片有一个长度为0，那么就不需要拷贝，直接 return </span>
    <span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>fm<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token operator">||</span> <span class="token function">len</span><span class="token punctuation">(</span>to<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
        <span class="token keyword">return</span> <span class="token number">0</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// n 记录下源切片或者目标切片较短的那一个的长度</span>
    n <span class="token operator">:=</span> <span class="token function">len</span><span class="token punctuation">(</span>fm<span class="token punctuation">)</span>
    <span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>to<span class="token punctuation">)</span> <span class="token operator">&lt;</span> n <span class="token punctuation">{</span>
        n <span class="token operator">=</span> <span class="token function">len</span><span class="token punctuation">(</span>to<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 如果开启了竞争检测</span>
    <span class="token keyword">if</span> raceenabled <span class="token punctuation">{</span>
        callerpc <span class="token operator">:=</span> <span class="token function">getcallerpc</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>to<span class="token punctuation">)</span><span class="token punctuation">)</span>
        pc <span class="token operator">:=</span> <span class="token function">funcPC</span><span class="token punctuation">(</span>slicestringcopy<span class="token punctuation">)</span>
        <span class="token function">racewriterangepc</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>to<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token punctuation">)</span><span class="token punctuation">,</span> callerpc<span class="token punctuation">,</span> pc<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 如果开启了 The memory sanitizer (msan)</span>
    <span class="token keyword">if</span> msanenabled <span class="token punctuation">{</span>
        <span class="token function">msanwrite</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>to<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
    <span class="token comment">// 拷贝字符串至字节数组</span>
    <span class="token function">memmove</span><span class="token punctuation">(</span>unsafe<span class="token punctuation">.</span><span class="token function">Pointer</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>to<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">stringStructOf</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>fm<span class="token punctuation">)</span><span class="token punctuation">.</span>str<span class="token punctuation">,</span> <span class="token function">uintptr</span><span class="token punctuation">(</span>n<span class="token punctuation">)</span><span class="token punctuation">)</span>
    <span class="token keyword">return</span> n
<span class="token punctuation">}</span>

</code></pre>
<p>再举个例子，比如：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    slice <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token number">3</span><span class="token punctuation">)</span>
    n <span class="token operator">:=</span> <span class="token function">copy</span><span class="token punctuation">(</span>slice<span class="token punctuation">,</span> <span class="token string">"abcdef"</span><span class="token punctuation">)</span>
    fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span>n<span class="token punctuation">,</span>slice<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre>
<p>输出：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token number">3</span> <span class="token punctuation">[</span><span class="token number">97</span><span class="token punctuation">,</span><span class="token number">98</span><span class="token punctuation">,</span><span class="token number">99</span><span class="token punctuation">]</span>

</code></pre>
<p>说到拷贝，切片中有一个需要注意的问题。</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go"><span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>  
    slice <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token number">10</span><span class="token punctuation">,</span> <span class="token number">20</span><span class="token punctuation">,</span> <span class="token number">30</span><span class="token punctuation">,</span> <span class="token number">40</span><span class="token punctuation">}</span>
    <span class="token keyword">for</span> index<span class="token punctuation">,</span> value <span class="token operator">:=</span> <span class="token keyword">range</span> slice <span class="token punctuation">{</span>
        fmt<span class="token punctuation">.</span><span class="token function">Printf</span><span class="token punctuation">(</span><span class="token string">"value = %d , value-addr = %x , slice-addr = %x\n"</span><span class="token punctuation">,</span> value<span class="token punctuation">,</span> <span class="token operator">&amp;</span>value<span class="token punctuation">,</span> <span class="token operator">&amp;</span>slice<span class="token punctuation">[</span>index<span class="token punctuation">]</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span>
<span class="token punctuation">}</span>

</code></pre>
<p>输出：</p>
<p>Go</p>
<pre class=" language-go"><code class="prism  language-go">value <span class="token operator">=</span> <span class="token number">10</span> <span class="token punctuation">,</span> value<span class="token operator">-</span>addr <span class="token operator">=</span> c4200aedf8 <span class="token punctuation">,</span> slice<span class="token operator">-</span>addr <span class="token operator">=</span> c4200b0320  
value <span class="token operator">=</span> <span class="token number">20</span> <span class="token punctuation">,</span> value<span class="token operator">-</span>addr <span class="token operator">=</span> c4200aedf8 <span class="token punctuation">,</span> slice<span class="token operator">-</span>addr <span class="token operator">=</span> c4200b0328  
value <span class="token operator">=</span> <span class="token number">30</span> <span class="token punctuation">,</span> value<span class="token operator">-</span>addr <span class="token operator">=</span> c4200aedf8 <span class="token punctuation">,</span> slice<span class="token operator">-</span>addr <span class="token operator">=</span> c4200b0330  
value <span class="token operator">=</span> <span class="token number">40</span> <span class="token punctuation">,</span> value<span class="token operator">-</span>addr <span class="token operator">=</span> c4200aedf8 <span class="token punctuation">,</span> slice<span class="token operator">-</span>addr <span class="token operator">=</span> c4200b0338

</code></pre>
<p>从上面结果我们可以看到，如果用 range 的方式去遍历一个切片，拿到的 Value 其实是切片里面的值拷贝。所以每次打印 Value 的地址都不变。</p>
<p><img src="https://ob6mci30g.qnssl.com/Blog/ArticleImage/57_12.png" alt=""></p>
<p>由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，需要通过  <code>&amp;slice[index]</code>获取真实的地址。</p>
<blockquote>
<p>reference:<br>
<a href="https://halfrost.com/go_slice/">https://halfrost.com/go_slice/</a></p>
</blockquote>

