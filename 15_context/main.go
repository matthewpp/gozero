package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context package in go provides a way to carry deadlines, cancellation signals and other
	request-scoped value across API boundaries and between processes.
*/

func main() {

	/* timeout context */
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation time out:", ctx.Err())
	}

	/* cancel context */

	// ctx, cancel := context.WithCancel(context.Background())

	// go worker(ctx)

	// time.Sleep(3 * time.Second)

	// fmt.Println("Main: Sending cancel signal to worker")
	// cancel()

	// time.Sleep(1 * time.Second)
	// fmt.Println("Main: Done")

	/* context with value */
	// ctx := context.Background()

	// userIDKey := key("userID")
	// value := "123"

	// ctx = context.WithValue(ctx, userIDKey, value)

	// showCTXVal(ctx, userIDKey)

	/* context http cancel */
	httpRequest()
}

/*
	good practice
	- Pass context.Context as the first parameter in functions.
	- Avoid storing contexts in structs.
	- Use context.WithTimeout and context.WithDeadline to control execution time.
	- Minimize usage of context.WithValue; prefer structs for complex data.
	- Always handle context cancellation by checking ctx.Done() or ctx.Err().
	- Shouldn't be used for general-purpose storage.
	- Use for storing request-scoped data
*/

/*
	in Go, the HTTP request context is automatically canceled
	when the request is complete, regardless of success, failure, or cancellation.
*/
