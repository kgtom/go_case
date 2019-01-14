
默认情况下，go出现错误会呈现以下信息，哪个文件、哪个函数、哪一行。为了防止个人信息泄露及被逆向，进行以下两个操作。




~~~

/Users/tom/goprojects/src/otw/api/handlers/orderHander.go:27 (0x1740524)
        (*OrderClient).CreateOrder: res, err := cli.CreateOrder(context.TODO(), &orderPb.Order{})
/Users/tom/goprojects/src/otw/api/main.go:33 (0x1740cc3)
        (*OrderClient).CreateOrder-fm: router.GET("/add/:name", orderApi.CreateOrder)
/Users/tom/goprojects/src/github.com/gin-gonic/gin/context.go:108 (0x1523152)
 

~~~
### 1.删除调试信息
~~~
go build -ldflags "-s -w” [<your/package]
~~~
（需要Go版本大于1.7）

这里的 -ldflags 参数最终会在 go tool link 的时候传给它， go tool link -h解释如下

  ~~~
 -s    disable symbol table
 -w    disable DWARF generation
 
 ~~~
### 2.删除trace文件信息
trace文件信息来自于 GOPATH GOROOT 环境变量，编译时GO会从 $GOPATH 寻找我们自己的代码src，从$GOROOT 提取标准库，在打包时将GOROOT改写为GOROOT_FINAL并作为trace信息的一部分写入目标文件。于是我们可以修改环境变量，隐藏真实的。代码如下：

~~~
ACTUAL_GOPATH="/Users/tom/goprojects"
export GOPATH='/tmp/go'
export GOROOT_FINAL=$GOPATH
[ ! -d $GOPATH ] && ln -s "$ACTUAL_GOPATH" "$GOPATH"
[[ ! $PATH =~ $GOPATH ]] && export PATH=$PATH:$GOPATH/bin
~~~

>reference
* [wx](https://mp.weixin.qq.com/s/YbaM-_vs_D2BS1lV6Z-u4g)
