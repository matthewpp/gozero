package calculator

import (
	"errors"
)

var (
	ErrDivideByZero = errors.New("divide by zero")
)

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return float64(a) / float64(b), nil
}
