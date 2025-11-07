package exam

import "fmt"

// implement
type CreditCard struct {
}

func (c CreditCard) Pay(v int) {
	fmt.Printf("pay with credit card: %v\n", v)
}

type MobileBanking struct {
	Name string
}

func (m MobileBanking) Pay(v int) {
	fmt.Printf("pay with mobile banking with %s: %v\n", m.Name, v)
}

func Basic() {
	c := CreditCard{}
	pay(c)

	m := MobileBanking{Name: "John"}
	pay(m)
}

func pay(p Payment) {
	price := 100
	p.Pay(price)
}
