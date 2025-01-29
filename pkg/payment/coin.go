package payment

import "fmt"

type Coin struct {
}

func (c *Coin) PayAmount(amount int) {
	fmt.Println(" coin amount paid is", amount)
}
