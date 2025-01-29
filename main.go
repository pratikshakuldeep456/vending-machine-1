package main

import (
	"pratikshakuldeep456/vending-machine/pkg/item"
	"pratikshakuldeep456/vending-machine/pkg/machine"
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

	newVendingMachine.DispenseItem("Ankit", 1)
}
