package main

import (
	"pratikshakuldeep456/vending-machine/pkg/item"
	"pratikshakuldeep456/vending-machine/pkg/machine"
	"pratikshakuldeep456/vending-machine/pkg/payment"
)

func main() {

	newVendingMachine := &machine.VendingMachine{}

	newItem := &item.Item{
		Count: 10,
		Code:  "Ankit",
		Price: 100,
	}
	newItem2 := &item.Item{
		Count: 10,
		Code:  "Cookie",
		Price: 100,
	}
	newVendingMachine.AddItems(*newItem)
	newVendingMachine.AddItems(*newItem2)

	newVendingMachine.DisplayItem()

	//payment

	payments := &payment.PaymentContext{
		Item: item.Item{
			Count: 4,
			Code:  "Ankit",
			Price: 10,
		},
	}
	coin := &payment.Card{}

	payments.SetMethod(coin)
	payments.Checkout()

	newVendingMachine.DispenseItem(payments.Item.Code, payments.Item.Count)
}
