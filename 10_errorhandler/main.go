package main

import (
	"errors"
	"fmt"

	"hello/hello"
)

func main() {

	_, err := hello.FirstError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = hello.SecondError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = hello.ThirdError()
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = hello.FourthError()
	if err != nil {
		fmt.Println(err.Error())
	}

	/* we need to check error and handle it */
	err = hello.Payment(5)
	if err != nil {
		switch err {
		case hello.LowerPayment:
			fmt.Println("you pay lower price")
		case hello.OverPayment:
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

	err = hello.NewPayment(5)
	if err != nil {
		var pe hello.PayErr
		if errors.As(err, &pe) {
			fmt.Println(pe.Error())
			fmt.Printf("%+v\n", pe.Info())
			fmt.Printf("error code: %s\n", pe.Details["code"])
		} else {
			fmt.Println("new payment unknow error")
		}
	}

}
