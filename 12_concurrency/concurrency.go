package main

import (
	"fmt"
	"sync"
	"time"
)

func withOutGoRoutine() {
	/* with out go routine */
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		fmt.Printf("val is: %s\n", v)
		time.Sleep(time.Millisecond * 500)
	}
}

func goRoutineWithoutWaitGroup() {
	/* go routine without wait group */
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		go func() {
			fmt.Printf("val is: %s\n", v)
			time.Sleep(time.Millisecond * 500)
		}()
	}
}

func goroutineWithSyncWaitGroup() {
	/* goroutine with sync wait group */
	var wg sync.WaitGroup

	wg.Add(10)
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		go func() {
			fmt.Printf("val is: %s\n", v)
			time.Sleep(time.Millisecond * 500)
			wg.Done()
		}()
	}
	wg.Wait()
}

func worker(id int, jobs <-chan job, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j.name)
		time.Sleep(j.processTime)
		fmt.Println("worker", id, "finished job", j.name)
		results <- fmt.Sprintf("job %s run success", j.name)
	}
}

type job struct {
	name        string
	processTime time.Duration
}

func workerPool() {
	/*--- worker pool ---*/
	const numJobs = 10
	const workerLimit = 3
	jobs := make(chan job, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= workerLimit; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- job{
			name:        fmt.Sprintf("job %d", j),
			processTime: time.Duration(time.Second * time.Duration(j)),
		}
		//	if j == 5 {
		//		close(jobs) // panic send close channel
		//	}
		//}
	}
	close(jobs)

	//for a := 1; a <= 11; a++ { // dead lock
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("all jobs processed")
}

func selectConcurrency() {
	/* -- select ---*/
	c1 := make(chan int)
	c2 := make(chan int)
	done1 := make(chan bool)
	done2 := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			c1 <- i
		}

		done1 <- true

	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			c2 <- i
		}

		done2 <- true

	}()

	var v1 bool
	var v2 bool

outerloop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received c1", msg1)
		case msg2 := <-c2:
			fmt.Println("received c2", msg2)
		case v1 = <-done1:
		case v2 = <-done2: // not good practice
			if v1 && v2 {
				fmt.Printf("all finish\n")
				break outerloop
			}
		}
	}
}
