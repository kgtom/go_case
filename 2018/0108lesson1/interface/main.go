package main

import (
	"0108LESSON1/interface/pay"
	"fmt"
)

func main() {

	//第一种根据客户端要实现那个实例去执行相应策略的方法
	//cn := cash.CashNormal{}
	//cndemo := cash.NewPayContextByEntity(cn)
	cndemo := pay.NewPayContextByEntity(pay.PayNormal{})
	var m = cndemo.PayMoney(3.00)
	fmt.Println("第一种方式：原价", m)

	//第二种根据参数，利用工厂模式去执行相应的方法
	money := 100.0
	cc := pay.NewPayContext("八折")
	money = cc.PayMoney(money)
	fmt.Println("第二种方式：100八折实际金额为", money)

	money = 200
	cc = pay.NewPayContext("原价")
	money = cc.PayMoney(money)
	fmt.Println("第二种方式：200原价为", money)
}
