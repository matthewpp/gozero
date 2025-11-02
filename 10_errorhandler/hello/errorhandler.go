package hello

import (
	"errors"
	"fmt"
)

var thirdErr = errors.New("this is third error")
var OverPayment = errors.New("over price to pay")
var LowerPayment = errors.New("lower price to pay")

func FirstError() (string, error) {
	return "", fmt.Errorf("this is first error")
}

func SecondError() (string, error) {
	return "", errors.New("this is second error")
}

func ThirdError() (string, error) {
	return "", thirdErr
}

type myError struct {
	msg string
}

func (m myError) Error() string {
	return m.msg
}

func FourthError() (string, error) {
	return "", myError{
		msg: "this is fourth error",
	}
}

func Payment(v int) error {

	if v < 100 {
		return LowerPayment
	}

	if v > 10000 {
		return OverPayment
	}

	return nil
}

type PayErr struct {
	msg     string
	Details map[string]string
}

func (p PayErr) Error() string {
	return p.msg
}

func (p PayErr) Info() map[string]string {
	return p.Details
}

func NewPayment(v int) error {
	if v < 100 {
		return PayErr{
			msg: "you pay lower price",
			Details: map[string]string{
				"code": "001",
			},
		}
	}

	if v > 10000 {
		return PayErr{
			msg: "you pay over price",
			Details: map[string]string{
				"code": "002",
			},
		}
	}

	return nil
}
