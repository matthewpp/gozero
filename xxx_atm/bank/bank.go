package bank

import "fmt"

type mtlBank struct {
	money int
}

func (b *mtlBank) Deposit(v int) error {
	b.money += v
	return nil
}

func (b *mtlBank) Withdraw(v int) (int, error) {

	if b.money < v {
		return 0, fmt.Errorf("money in bank not enough")
	}

	b.money -= v

	return v, nil
}

func (b *mtlBank) Balance() int {
	return b.money
}

func New(m int) *mtlBank {
	return &mtlBank{
		money: m,
	}
}
