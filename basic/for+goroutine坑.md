### 问题：
~~~
for i := 0; i < 3; i++ {

        fmt.Println("i", i)

        go fmt.Println("i2", i)

    }
~~~

**有时候 for执行完了，才执行goroutine，而此时gourine里面参数已变成最后一次迭代值**
解决这个问题参数传递过去。如下：

~~~
func main() {

    for i := 0; i < 5; i++ {

        fmt.Println("i", i)

        go func(val int) {
            fmt.Println("val", val)
        }(i)

        if i < 1 {

            break
        }
    }
    time.Sleep(1 * time.Second)
    fmt.Println("end")
    }

~~~
* 去掉time.Sleep,因为for 迭代完成后，main 主函数也就结束了，gorounte没有机会走。生产环境因为程序一直活着，不会存在执行不到goroutine的时候。
