package hello

import (
	"fmt"
	"runtime"
)

func MyPanic() {
	// Trigger a panic with a message
	panic("just from myPanic")
}

func MyRecovery() {
	defer func() {
		fmt.Println("Defer Function")

		// Attempt to recover from panic
		if r := recover(); r != nil {
			fmt.Println("recover from panic: ", r)
		}

		fmt.Println("End Defer")
	}()

	fmt.Println("Start MyRecovery")

	// Trigger a panic with a message
	panic("just from myRecovery ")
}

func MyRecoveryWithTrace() {
	// Defer a function to handle panic and print stack trace
	defer func() {
		// Attempt to recover from panic
		if r := recover(); r != nil {
			// Create a buffer to store stack trace
			var buf [4096]byte
			// Get the stack trace into the buffer
			v := runtime.Stack(buf[:], false)
			// Print the recovered panic message
			fmt.Println("recover from panic: ", r)
			// Print the stack trace
			fmt.Printf("stack trace %v\n", string(buf[:v]))
		}
	}()

	// Trigger a panic with a message
	panic("just from myRecovery")
}

func TestDefer() int {
	fmt.Println("Function called")

	// defer statements are executed after the return statement is evaluated
	// but before the result is returned to the caller
	defer fmt.Println("Deferred function executed")

	fmt.Println("Before return")

	return 10
}
