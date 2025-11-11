package main

import (
	"errors"
	"fmt"

	"hello/hello"
)

func main() {
	//howToInitError()
	//
	//howToCompareError()
	//
	//howToCompareTypeError()

	howToUseCompareTypeError()

	//badErrorPractice()
}

func howToInitError() {
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
}

func howToCompareError() {
	/* we need to check error and handle it */
	err := hello.Payment(5)
	if err != nil {
		switch err {
		case hello.LowerPayment:
			fmt.Println("you pay lower price")
		case hello.OverPayment:
			fmt.Println("you pay over price")
		default:
			fmt.Println("unknow error")
		}

		//if err == hello.LowerPayment {
		//	fmt.Println("you pay lower price pay again")
		//}

		if errors.Is(err, hello.LowerPayment) {
			fmt.Println("you pay lower price pay again")
		} else if errors.Is(err, hello.OverPayment) {
			fmt.Println("you pay over price pay again")
		} else {
			fmt.Println("unknow error")
		}
	}
}

func howToCompareTypeError() {
	var payErr hello.PayErr
	pErr := hello.PayOverErr
	if errors.As(pErr, &payErr) {
		fmt.Println("this error type is hello.PayErr")
	}

	pErr = hello.PayLowErr
	if errors.As(pErr, &payErr) {
		fmt.Println("this error type is hello.PayErr")
	}

	cErr := hello.CreditLowErr
	if errors.As(cErr, &payErr) {
		fmt.Println("this error type is hello.PayErr")
	} else {
		fmt.Println("this error type is not hello.PayErr")
	}
}

func howToUseCompareTypeError() {
	err := hello.NewPayment(5)
	if err != nil {
		mes := err.Error()
		fmt.Println("payment error message err.Error():", mes)
		//info := err.Info()
		//fmt.Println("payment error message err.Info():", info)
		payErr, ok := err.(hello.PayErr) // type assertion but not recommend
		if ok {
			fmt.Printf("%+v\n", payErr.Info())
		}
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

func badErrorPractice() {
	creditErr := hello.Credit(5)
	if creditErr.Error() != "" {
		switch creditErr.Error() {
		case "you pay lower price":
			fmt.Println("you pay lower price")
		case "you pay over price":
			fmt.Println("you pay over price")
		default:
			fmt.Println("unknow error")
		}

		if errors.Is(creditErr, hello.LowerPayment) {
			fmt.Println("you pay lower price pay again")
		} else if errors.Is(creditErr, hello.OverPayment) {
			fmt.Println("you pay over price pay again")
		} else {
			fmt.Println("unknow error")
		}
	}
}
