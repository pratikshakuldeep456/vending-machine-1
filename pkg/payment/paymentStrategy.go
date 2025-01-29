package payment

import (
	"fmt"
)

type PaymentStrategy interface {
	PayAmount(amount int)
}

type ProductItem struct {
	Amount   int
	Code     string
	Quantity int
}

type PaymentContext struct {
	PaymentMethod PaymentStrategy
	Product       ProductItem
}

func (p *PaymentContext) SetMethod(ps PaymentStrategy) {
	p.PaymentMethod = ps

}

func (p *PaymentContext) Checkout() {

	fmt.Println("checkout started ")

	p.PaymentMethod.PayAmount(p.Product.Amount)

}
