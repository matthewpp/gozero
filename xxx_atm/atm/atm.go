package atm

import (
	"fmt"
	"slices"
	"sync"
)

type Banker interface {
	Deposit(int) error
	Withdraw(int) (int, error)
	Balance() int
}

type myATM struct {
	money int
	bank  Banker
	mu    *sync.Mutex
}

func (a *myATM) Deposit(v int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if v < 0 {
		return fmt.Errorf("deposit amount must greater than zero")
	}

	if !slices.Contains([]int{1000, 2000, 3000, 4000, 5000}, v) {
		return fmt.Errorf("accept only [1000, 2000, 3000, 4000, 5000]")
	}

	a.money += v
	err := a.bank.Deposit(v)
	if err != nil {
		return err
	}

	return nil
}

func (a *myATM) Withdraw(v int) (int, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if v < 0 {
		return 0, fmt.Errorf("withdraw amount must greater than zero")
	}

	if a.money == 0 {
		return 0, fmt.Errorf("no money on atm machine")
	}

	if a.money < v {
		return 0, fmt.Errorf("atom money not enough")
	}

	if !slices.Contains([]int{1000, 2000, 3000, 4000, 5000}, v) {
		return 0, fmt.Errorf("accept only [1000, 2000, 3000, 4000, 5000]")
	}

	a.money -= v

	val, err := a.bank.Withdraw(v)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (m *myATM) Balance() int {
	return m.bank.Balance()
}

func New(b Banker, m int) *myATM {
	return &myATM{
		money: m,
		bank:  b,
		mu:    &sync.Mutex{},
	}
}
