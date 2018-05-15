---


---

<h3 id="一、grpc-是什么">一、GRPC 是什么</h3>
<p>gRPC 基于如下思想：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。gRPC 默认使用 <a href="https://developers.google.com/protocol-buffers/">pb文件</a> 作为接口定义语言，来描述服务接口和有效载荷消息结构。</p>
<h4 id="四种服务类型">四种服务类型</h4>
<ul>
<li>单项 RPC，即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。</li>
</ul>
<pre><code>rpc SayHello(HelloRequest) returns (HelloResponse){
}

</code></pre>
<ul>
<li>服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。</li>
</ul>
<pre><code>rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
}

</code></pre>
<ul>
<li>客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。</li>
</ul>
<pre><code>rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
}

</code></pre>
<ul>
<li>双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。</li>
</ul>
<pre><code>rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
}
</code></pre>
<h3 id="二、为什么使用-grpc">二、为什么使用 gRPC?</h3>
<p>一次性的在一个 .proto 文件中定义服务并使用任何支持它的语言去实现客户端和服务器。</p>
<h3 id="三、实践案例">三、实践案例</h3>
<p>通过学习教程中例子，你可以学会如何：</p>
<ul>
<li>在一个 .proto 文件内定义服务。</li>
<li>用 protocol buffer 编译器生成服务器和客户端代码。</li>
<li>使用 gRPC 的 Go API 为你的服务实现一个简单的客户端和服务器。</li>
</ul>
<p>// todo</p>
<blockquote>
<p>reference<a href="http://doc.oschina.net/grpc?t=60133">addr:</a>.<br>
<a href="https://www.cnblogs.com/YaoDD/p/5504881.html">单项rpc例子:</a><br>
<a href="https://www.epubit.com/selfpublish/article/1922">双向流式rpc例子</a></p>
</blockquote>

