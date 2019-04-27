 ~~~
 “为什么 Go map 遍历输出是不固定顺序？”。而通过这一番分析，原因也很简单明了。就是 for range map 在开始处理循环逻辑的时候，就做了随机播种...
~~~ 

### 看一下汇编
~~~
    ...
    0x009b 00155 (main.go:11)    LEAQ    type.map[int32]string(SB), AX
    0x00a2 00162 (main.go:11)    PCDATA    $2, $0
    0x00a2 00162 (main.go:11)    MOVQ    AX, (SP)
    0x00a6 00166 (main.go:11)    PCDATA    $2, $2
    0x00a6 00166 (main.go:11)    LEAQ    ""..autotmp_3+24(SP), AX
    0x00ab 00171 (main.go:11)    PCDATA    $2, $0
    0x00ab 00171 (main.go:11)    MOVQ    AX, 8(SP)
    0x00b0 00176 (main.go:11)    PCDATA    $2, $2
    0x00b0 00176 (main.go:11)    LEAQ    ""..autotmp_2+72(SP), AX
    0x00b5 00181 (main.go:11)    PCDATA    $2, $0
    0x00b5 00181 (main.go:11)    MOVQ    AX, 16(SP)
    0x00ba 00186 (main.go:11)    CALL    runtime.mapiterinit(SB)
    0x00bf 00191 (main.go:11)    JMP    207
    0x00c1 00193 (main.go:11)    PCDATA    $2, $2
    0x00c1 00193 (main.go:11)    LEAQ    ""..autotmp_2+72(SP), AX
    0x00c6 00198 (main.go:11)    PCDATA    $2, $0
    0x00c6 00198 (main.go:11)    MOVQ    AX, (SP)
    0x00ca 00202 (main.go:11)    CALL    runtime.mapiternext(SB)
    0x00cf 00207 (main.go:11)    CMPQ    ""..autotmp_2+72(SP), $0
    0x00d5 00213 (main.go:11)    JNE    193
    ...
    
    
 ~~~
    
我们大致看一下整体过程，重点处理 Go map 循环迭代的是两个 runtime 方法，如下：


runtime.mapiterinit
runtime.mapiternext



但你可能会想，明明用的是 for range 进行循环迭代，怎么出现了这两个函数，怎么回事？



~~~go 

runtime.mapiterinit
func mapiterinit(t *maptype, h *hmap, it *hiter) {
    ...
    it.t = t
    it.h = h
    it.B = h.B
    it.buckets = h.buckets
    if t.bucket.kind&kindNoPointers != 0 {
        h.createOverflow()
        it.overflow = h.extra.overflow
        it.oldoverflow = h.extra.oldoverflow
    }

    r := uintptr(fastrand())
    if h.B > 31-bucketCntBits {
        r += uintptr(fastrand()) << 31
    }
    it.startBucket = r & bucketMask(h.B)
    it.offset = uint8(r >> h.B & (bucketCnt - 1))
    it.bucket = it.startBucket
    ...

    mapiternext(it)
}

~~~

通过对 mapiterinit 方法阅读，可得知其主要用途是在 map 进行遍历迭代时进行初始化动作。共有三个形参，用于读取当前哈希表的类型信息、当前哈希表的存储信息和当前遍历迭代的数据

为什么
咱们关注到源码中 fastrand 的部分，这个方法名，是不是迷之眼熟。没错，它是一个生成随机数的方法。再看看上下文：

~~~
...
// decide where to start
r := uintptr(fastrand())
if h.B > 31-bucketCntBits {
    r += uintptr(fastrand()) << 31
}
it.startBucket = r & bucketMask(h.B)
it.offset = uint8(r >> h.B & (bucketCnt - 1))

// iterator state
it.bucket = it.startBucket

~~~

在这段代码中，它生成了随机数。用于决定从哪里开始循环迭代。更具体的话就是根据随机数，选择一个桶位置作为起始点进行遍历迭代

因此每次重新 for range map，你见到的结果都是不一样的。那是因为它的起始位置根本就不固定！


   咱们已经选定了起始桶的位置。接下来就是通过 mapiternext 进行具体的循环遍历动作。该方法主要涉及如下：

从已选定的桶中开始进行遍历，寻找桶中的下一个元素进行处理
如果桶已经遍历完，则对溢出桶 overflow buckets 进行遍历处理
通过对本方法的阅读，可得知其对 buckets 的遍历规则以及对于扩容的一些处理（这不是本文重点。因此没有具体展开）

### 总结


在本文开始，咱们先提出核心讨论点：“为什么 Go map 遍历输出是不固定顺序？”。而通过这一番分析，原因也很简单明了。就是 for range map 在开始处理循环逻辑的时候，就做了随机播种... 你想问为什么要这么做？当然是官方有意为之，因为 Go 在早期（1.0）的时候，虽是稳定迭代的，但从结果来讲，其实是无法保证每个 Go 版本迭代遍历规则都是一样的。而这将会导致可移植性问题。因此，改之。也请不要依赖...




> reference:
[segmentfault](https://segmentfault.com/a/1190000018782278)
