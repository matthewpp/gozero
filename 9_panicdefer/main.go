package main

import (
	"fmt"

	"hello/hello"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		if i == 10 {
			hello.MyPanic() // Triggers a panic to demonstrate panic behavior
			//hello.MyRecovery()            // Recovers from panic using recover() in a deferred function
			//hello.MyRecoveryWithTrace()
		}
	}

	// hello.TestDefer()

}
