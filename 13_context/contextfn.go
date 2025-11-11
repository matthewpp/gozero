package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func contextTimeout() {
	/* timeout context */
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	//case <-time.After(2 * time.Second):
	case <-time.After(4 * time.Second):
		fmt.Println("operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation time out:", ctx.Err())
	}
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

type key string

func showCTXVal(ctx context.Context, userIDKey key) {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found in context")
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

func processTask(ctx context.Context, id int, shouldFail bool) error {
	for i := 0; i < 10; i++ {
		// เช็คว่า context ถูก cancel หรือไม่
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d: Cancelled at step %d\n", id, i)
			return ctx.Err()
		default:
		}

		// จำลองการทำงาน
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Task %d: Step %d\n", id, i)

		// จำลอง error
		if shouldFail && i == 3 {
			return fmt.Errorf("task %d failed at step %d", id, i)
		}
	}
	return nil
}

func handleErrorGroupWithContext() {
	g, ctx := errgroup.WithContext(context.Background())

	// Task 1: จะทำงานปกติ
	g.Go(func() error {
		return processTask(ctx, 1, false)
	})

	// Task 2: จะ error ที่ step 3
	g.Go(func() error {
		return processTask(ctx, 2, true)
	})

	// Task 3: จะถูก cancel เมื่อ Task 2 error
	g.Go(func() error {
		return processTask(ctx, 3, false)
	})

	// รอและแสดงผล
	if err := g.Wait(); err != nil {
		fmt.Printf("\n❌ Error occurred: %v\n", err)
	} else {
		fmt.Printf("\n✅ All tasks completed\n")
	}
}

func httpRequest() {
	/* context http cancel */
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
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
