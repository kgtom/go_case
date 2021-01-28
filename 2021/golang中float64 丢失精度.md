
### 在golang中，针对四舍五入的问题丢失精度问题，并且兼顾速度，采用下面代码
~~~

// Round 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 oriVal 根据指定精度 precision（小数点后数字的数目）进行四舍五入。precision 取值范围[正数 0 负数]。
func Round(oriVal float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(oriVal*p+0.5) / p
}
func main() {

	fmt.Println("start")
	fmt.Println(Round(10.13837, 0))
	fmt.Println(Round(10.13837, 1))
	fmt.Println(Round(10.13837, 2))
	fmt.Println(Round(10.13837, 3))
	fmt.Println(Round(10.13837, -1))
	fmt.Println(Round(712.13837, -2))

	fmt.Println("end")
~~~

> reference
* [so](https://stackoverflow.com/questions/39544571/golang-round-to-nearest-0-05)
