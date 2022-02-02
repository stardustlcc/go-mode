package main

import "fmt"

/*
	定义了策略家族，分别封装起来，让它们之间可以互相替换，此模式让策略的变化，不会影响到使用策略的客户（多态）。实际上客户调用的是抽象的策略接口，不用关心策略的底层实现。
	以超市打折为例，正常收费，打折收费和返利收费是具体策略，用一个策略接口去管理具体策略；同时还可以与简单工厂模式结合，传入策略名称，由工厂去创建策略管理类。
*/

//首先定义了3个常量，用于规范具体的策略名称
const (
	original_price    = iota //0
	ten_percent_off          //1
	full_100_minus_20        //2
)

//然后定义一个打折接口和实现打折接口的具体策略，打折接口Pricer和打折方法Price是要被外界调用的，所以需要首字母大写：
type Pricer interface {
	Price(cash float64) float64
}

//原价
type normal struct{}

//打折
type discount struct {
	percent float64
}

//满减
type reduction struct {
	fullMoney     float64
	discountMoney float64
}

//均实现了Price接口
func (n *normal) Price(cash float64) float64 {
	return cash
}
func (d *discount) Price(cash float64) float64 {
	return cash * d.percent
}
func (r *reduction) Price(cash float64) float64 {
	count := cash / r.fullMoney
	return cash - count*r.discountMoney
}

//最后定义生成具体打折策略的工厂CashFactory，和工厂实现获取具体策略的方法GetCashPrice，由于均需要被外界调用，所以首字母需要大写：
type CashFactory struct{}

func (c CashFactory) GetCashPrice(discountType int) Pricer {
	var cashPricer Pricer
	switch discountType {
	case original_price:
		cashPricer = &normal{}
	case ten_percent_off:
		cashPricer = &discount{0.5}
	case full_100_minus_20:
		cashPricer = &reduction{100, 20}
	}
	return cashPricer
}

func main() {

	var cf CashFactory
	var price Pricer
	price = cf.GetCashPrice(original_price)
	fmt.Println(price.Price(200))

	price = cf.GetCashPrice(ten_percent_off)
	fmt.Println(price.Price(200))

	price = cf.GetCashPrice(full_100_minus_20)
	fmt.Println(price.Price(200))
}
