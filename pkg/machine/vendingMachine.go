package machine

import (
	"fmt"
	"pratikshakuldeep456/vending-machine/pkg/balance"
	"pratikshakuldeep456/vending-machine/pkg/item"
	"pratikshakuldeep456/vending-machine/pkg/payment"
)

type VendingMachine struct {
	Items    []item.Item
	Balances balance.Balance
}

func (vm *VendingMachine) AddItems(item item.Item) {
	vm.Items = append(vm.Items, item)
}

func (vm *VendingMachine) DisplayItem() []item.Item {
	var items []item.Item

	for _, v := range vm.Items {
		items = append(items, v.GetCode())
	}

	return items
}

func (vm *VendingMachine) DispenseItem(p1 payment.ProductItem) {
	fmt.Println(p1)
	fmt.Println(
		"dispense method called")
	for _, v := range vm.Items {
		if v.Code == p1.Code {
			calculateAmount := p1.Quantity * v.Price

			if calculateAmount > p1.Amount {
				fmt.Println(" Amount provided is less for given quantity")
			}

			remainingAmount := (p1.Amount - calculateAmount)
			err := v.RemoveQuantity(p1.Quantity)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Dispensing item: ", p1.Code, "with Quantity", p1.Quantity)

			if remainingAmount > 0 {
				fmt.Println("returning extra coin to user", remainingAmount)
			}
			break
		} else {
			fmt.Println(" invalid product code")
		}
	}
}
