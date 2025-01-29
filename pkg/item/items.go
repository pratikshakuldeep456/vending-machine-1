package item

import "errors"

type Item struct {
	Count int
	Code  string
	Price int
}

func (i *Item) IsAvailable() bool {

	return (i.Count != 0)
}

func (i *Item) AddQuantity(quantity int) {

	i.Count += quantity

}

func (i *Item) RemoveQuantity(quantity int) error {
	if i.Count < quantity {
		return errors.New("insuff quantity bro")
	}
	i.Count -= quantity
	return nil

}

func (i *Item) GetCode() Item {
	return Item{
		Count: i.Count,
		Code:  i.Code,
		Price: i.Price,
	}
}
