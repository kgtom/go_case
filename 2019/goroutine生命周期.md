

### 一、main goroutine 以及普通 goroutine 从执行到退出goexit0()的整个过程。



goexit0() 主要完成最后的清理工作:

* 1.把 g 的状态从 _Grunning 更新为 _Gdead；

* 2.清空 g 的一些字段；

* 3.调用 dropg 函数解除 g 和 m 之间的关系，其实就是设置 g->m = nil, m->currg = nil；

* 4.把 g 放入 p 的 freeg 队列缓存起来供下次创建 g 时快速获取而不用从内存分配。freeg 就是 g 的一个对象池；

* 5.调用 schedule 函数再次进行调度。

到这里，gp 就完成了它的历史使命，功成身退，进入了 goroutine 缓存池，待下次有任务再重新启用。

而工作线程，又继续调用 schedule 函数进行新一轮的调度，整个过程形成了一个循环。

### 二、我们继续探索非 main goroutine 的退出流程。

gp 执行完后，RET 指令弹出 goexit 函数地址（实际上是 funcPC(goexit)+1），CPU 跳转到 goexit 的第二条指令继续执行，然后执行 goexit0函数


### 总结一下，main goroutine 和普通 goroutine 的退出过程：

* 对于 main goroutine，在执行完用户定义的 main 函数的所有代码后，直接调用 exit(0) 退出整个进程，非常霸道。

* 对于普通 goroutine 则没这么“舒服”，需要经历一系列的过程。
 * 先是跳转到提前设置好的 goexit 函数的第二条指令，
 * 然后调用 runtime.goexit1，接着调用 mcall(goexit0)，而 mcall 函数会切换到 g0 栈，运行 goexit0 函数，
 * 清理 goroutine 的一些字段，并将其添加到 goroutine 缓存池里，然后进入 schedule 调度循环。到这里，普通 goroutine 才算完成使命。
 


> reference
* [wx](https://mp.weixin.qq.com/s/kwKqrT4BoeheM9MvSh-rLw)
