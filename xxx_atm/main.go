package main

import (
	"chonlatee/gozero/atm"
	"chonlatee/gozero/bank"
	"fmt"
)

func main() {
	b := bank.New(10000)

	a := atm.New(b, 10000)

	fmt.Printf("\t**** This is simple ATM machine ****")
	finish := false
	for !finish {
		choice := printMenu()
	switchChoices: // this is label
		switch choice {
		case 1:
			var amount int
			fmt.Printf("How much you want to deposit: ")
			fmt.Scan(&amount)
			err := a.Deposit(amount)
			if err != nil {
				fmt.Printf("Deposit error: %v", err.Error())
				break switchChoices // break from switch case not outer loop
			}
		case 2:
			var amount int
			fmt.Printf("How much you want to withdraw: ")
			fmt.Scan(&amount)
			v, err := a.Withdraw(amount)
			if err != nil {
				fmt.Printf("Withdraw error: %v", err.Error())
				break switchChoices
			}

			fmt.Printf("This is your money ( %d ) ... grahhh grahhh (atm machine sound)", v)
		case 3:
			v := a.Balance()
			fmt.Printf("This is your balance: %d", v)

		case 4:
			break switchChoices
		case 5:
			finish = true
			break switchChoices
		default:
			fmt.Printf("Unknow choice. Please try again")
		}
	}
}

func printMenu() int {
	var choice int
	fmt.Printf("\n")
	fmt.Printf("\t\t Please select ATM menu:\n")
	fmt.Printf("\t\t 1. Deposit\n")
	fmt.Printf("\t\t 2. Withdraw\n")
	fmt.Printf("\t\t 3. Check Balance\n")
	fmt.Printf("\t\t 4. Show menu\n")
	fmt.Printf("\t\t 5. Quit\n")
	fmt.Printf("Your choice: ")
	fmt.Scan(&choice)
	return choice
}
