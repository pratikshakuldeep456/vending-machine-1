package payment

import (
	"fmt"
	"pratikshakuldeep456/vending-machine/pkg/item"
)

type PaymentStrategy interface {
	PayAmount(amount int)
}

type PaymentContext struct {
	PaymentMethod PaymentStrategy
	Item          item.Item
}

func (p *PaymentContext) SetMethod(ps PaymentStrategy) {
	p.PaymentMethod = ps

}

func (p *PaymentContext) Checkout() {

	fmt.Println("checkout started ")

	p.PaymentMethod.PayAmount(p.Item.Price * p.Item.Count)

}
