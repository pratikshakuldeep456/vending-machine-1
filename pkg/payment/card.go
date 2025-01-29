package payment

import "fmt"

type Card struct {
}

func (c *Card) PayAmount(amount int) {
	fmt.Println("  card amount paid is", amount)
}
