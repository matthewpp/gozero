package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

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

type key string

func showCTXVal(ctx context.Context, userIDKey key) {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found in context")
	}
}

func httpRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

	fmt.Println("Response status code:", resp.StatusCode)
}