package payment

import (
	"errors"
	"fmt"
	"strings"
)

var ErrPayLower = errors.New("you pay lower price")
var ErrPayOver = errors.New("you pay over price")

type Payment struct {
}

func (p Payment) Pay(v int) error {
	return pay(v)
}

func (p Payment) PayWithCreditCard(v int, number string) error {

	return paymentWithCreditCard(v, number)
}

func pay(v int) error {

	if v < 100 {
		return ErrPayLower
	}

	if v > 1000 {
		return ErrPayOver
	}

	return nil
}

func paymentWithCreditCard(v int, number string) error {
	if strings.HasPrefix(number, "001") {
		return fmt.Errorf("we not accept card start with 001")
	}

	if strings.HasPrefix(number, "003") {
		return fmt.Errorf("we not accept card start with 003")
	}

	if strings.HasPrefix(number, "007") {
		return fmt.Errorf("we not accept card from james bond")
	}

	return nil
}
