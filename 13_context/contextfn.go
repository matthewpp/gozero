package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func contextTimeout() {
	/* timeout context */
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation time out:", ctx.Err())
	}
}

func cancelContext() {
	/* cancel context */
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	fmt.Println("Main: Sending cancel signal to worker")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Done")
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Stopping work due to cancel signal")
			return
		default:
			fmt.Println("Worker Working....")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func valueContext() {
	/* context with value */
	ctx := context.Background()

	userIDKey := key("userID")
	value := "123"

	ctx = context.WithValue(ctx, userIDKey, value)

	showCTXVal(ctx, userIDKey)
}

type key string

func showCTXVal(ctx context.Context, userIDKey key) {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found in context")
	}
}

func httpRequest() {
	/* context http cancel */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tt := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/delay/3", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Request completed in:", time.Since(tt))

	fmt.Println("Response status code:", resp.StatusCode)
}
