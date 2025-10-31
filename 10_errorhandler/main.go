package main

import (
	"errors"
	"fmt"
)

var thirdErr = errors.New("this is third error")

var overPayment = errors.New("over price to pay")
var lowerPayment = errors.New("lower price to pay")

func main() {

	_, err := firstError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = secondError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = thirdError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = fourthError()
	if err != nil {
		fmt.Println(err.Error())
	}

	/* we need to check error and handle it */
	err = payment(5)
	if err != nil {
		switch err {
		case lowerPayment:
			fmt.Println("you pay lower price")
		case overPayment:
			fmt.Println("you pay over price")
		default:
			fmt.Println("unknow error")
		}
		// if errors.Is(err, lowerPayment) {
		// 	fmt.Println("you pay lower price pay again")
		// } else if errors.Is(err, overPayment) {
		// 	fmt.Println("you pay over price pay again")
		// } else {
		// 	fmt.Println("unknow error")
		// }
	}

	err = newPayment(5)
	if err != nil {
		var pe payErr
		if errors.As(err, &pe) {
			fmt.Println(pe.Error())
			fmt.Printf("%+v\n", pe.Info())
			fmt.Printf("error code: %s\n", pe.info["code"])
		} else {
			fmt.Println("new payment unknow error")
		}
	}

}

func firstError() (string, error) {
	return "", fmt.Errorf("this is first error")
}

func secondError() (string, error) {
	return "", errors.New("this is second error")
}

func thirdError() (string, error) {
	return "", thirdErr
}

type myError struct {
	msg string
}

func (m myError) Error() string {
	return m.msg
}

func fourthError() (string, error) {
	return "", myError{
		msg: "this is fourth error",
	}
}

func payment(v int) error {

	if v < 100 {
		return lowerPayment
	}

	if v > 10000 {
		return overPayment
	}

	return nil
}

type payErr struct {
	msg  string
	info map[string]string
}

func (p payErr) Error() string {
	return p.msg
}

func (p payErr) Info() map[string]string {
	return p.info
}

func newPayment(v int) error {
	if v < 100 {
		return payErr{
			msg: "you pay lower price",
			info: map[string]string{
				"code": "001",
			},
		}
	}

	if v > 10000 {
		return payErr{
			msg: "you pay over price",
			info: map[string]string{
				"code": "002",
			},
		}
	}

	return nil
}
