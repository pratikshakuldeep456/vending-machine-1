package machine

import (
	"fmt"
	"pratikshakuldeep456/vending-machine/pkg/balance"
	"pratikshakuldeep456/vending-machine/pkg/item"
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

func (vm *VendingMachine) DispenseItem(code string, quantity int) {
	fmt.Println(
		"dispense method called")
	for _, v := range vm.Items {
		if v.Code == code {
			err := v.RemoveQuantity(quantity)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Dispensing item: ", code)

		}
	}
}
