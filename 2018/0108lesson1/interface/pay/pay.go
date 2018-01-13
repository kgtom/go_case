package pay

type PaySuper interface {
	PayMoney(money float64) float64
}

//普通情况，没有折扣
type PayNormal struct {
}

func newPayNormal() PayNormal {
	instance := new(PayNormal)
	return *instance
}

func (c PayNormal) PayMoney(money float64) float64 {
	return money
}

//打折，传入打折的折扣，如0.8
type PayRebate struct {
	Rebate float64 //折扣
}

func newPayRebate(rebate float64) PayRebate {
	instance := new(PayRebate)
	instance.Rebate = rebate
	return *instance
}

func (c PayRebate) PayMoney(money float64) float64 {
	return money * c.Rebate
}

type PayContext struct {
	Strategy PaySuper
}

func NewPayContext(cashType string) PayContext {
	c := new(PayContext)
	//使用简易工厂模式，用来生产策略,这样PayRebate，PayNormal首字母可以小写
	switch cashType {
	case "八折":
		c.Strategy = newPayRebate(0.8)

	default:
		c.Strategy = newPayNormal()
	}
	return *c
}
func NewPayContextByEntity(cashType PaySuper) PayContext {
	c := new(PayContext)
	c.Strategy = cashType
	return *c
}

//用策略实例后调用策略的函数。
func (c PayContext) PayMoney(money float64) float64 {
	return c.Strategy.PayMoney(money)
}
