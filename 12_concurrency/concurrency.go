package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func withOutGoRoutine() {
	/* with out go routine */
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("val is: %s\n", v)
	}
}

func goRoutineWithoutWaitGroup() {
	/* go routine without wait group */
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		go func() {
			sleepRandom := rand.Intn(100)
			time.Sleep(time.Millisecond * (500 + time.Duration(sleepRandom)))
			fmt.Printf("val is: %s\n", v)
		}()
	}

	time.Sleep(550 * time.Millisecond)
}

func goRoutineWithSyncWaitGroup() {
	/* goroutine with sync wait group */
	// with no handle error from goroutine
	var wg sync.WaitGroup

	//wg.Add(11) // for show deadlock or wait forever
	for _, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		go func() {
			wg.Add(1)
			sleepRandom := rand.Intn(100)
			time.Sleep(time.Millisecond * (500 + time.Duration(sleepRandom)))
			fmt.Printf("val is: %s\n", v)

			wg.Done()
		}()
	}

	wg.Wait()
}

func goRoutineWithErrorGroup() {
	/* goroutine with error group */
	// https://pkg.go.dev/golang.org/x/sync/errgroup
	// with handle error from goroutine
	var eg errgroup.Group

	for i, v := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"} {
		eg.Go(func() error {
			sleepRandom := rand.Intn(500)
			time.Sleep(time.Millisecond * (500 + time.Duration(sleepRandom)))
			fmt.Printf("val is: %s\n", v)
			if i%2 == 0 {
				return nil
			} else {
				return fmt.Errorf("some error happened for val %s and time %v \n", v, sleepRandom)
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("error:", err)
	}
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
		fmt.Println(<-results)
	}
	close(results)

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

	fmt.Println("end")
}

func breakOuterLoop() {
outerloop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf("i=%d j=%d k=%d\n", i, j, k)
				if i == 5 && j == 5 && k == 5 {
					fmt.Printf("break all loop\n")
					break outerloop
				}
			}
		}
	}
}
