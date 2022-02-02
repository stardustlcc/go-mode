package main

import "fmt"

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type Cash struct{}
type Bank struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}

/*
   定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。
*/
func main() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	payment = NewPayment("Bob", "00002", 888, &Bank{})
	payment.Pay()
}
