
## 简单并发性能demo
### 背景
~~~

  型号名称：	MacBook Pro
  型号标识符：	MacBookPro16,3
  处理器名称：	Quad-Core Intel Core i7
  内存：	16 GB
  GO version: go version go1.14.7 darwin/amd64
        
~~~

### 代码
 
 ~~~
 func GoTask(cnt int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("cpu核数=", runtime.NumCPU())
	fmt.Println("并发数量=", cnt)
	startTs := time.Now()
	for i := 0; i < cnt; i++ {
		go func() {}()
	}
	endTs := time.Now()
	fmt.Printf("cost sec: %.3fs\n", endTs.Sub(startTs).Seconds())
}
func main() {

	fmt.Println("start")
	cnt := 100 * 10000 //百万，比java 快10%
	GoTask(cnt)
	fmt.Println()
	cnt = 1000 * 10000 //千万，比java 快2倍
	GoTask(cnt)
	fmt.Println()
	cnt = 10000 * 10000 //亿，27s内完成，java因为线程数量大，程序崩溃
	GoTask(cnt)

	fmt.Println("end")
  
  }
 
 ~~~
