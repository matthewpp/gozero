package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*
		https://go.dev/talks/2012/waza.slide#1

		Concurrency vs 	Parallelism
		Concurrency is about dealing with lots of things at once.

		Parallelism is about doing lots of things at once.

		Not the same, but related.

		Concurrency is about structure, parallelism is about execution.

		Concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable.
			concurrent or parallel which one is better sometime it's depend on type of work.

			eg. if some work proces have to wait some other work process maybe use parallel not help.
	*/

	// mutex  use when you want to prevent mutual section - ( lock behavior, use with caution for performance issue)
	// channel use when you want to communicate between goroutine - ( use with caution, IMHO: complex to read and  debug for beginner).

	/* -- check number of cpu --*/
	fmt.Printf("%d\n", runtime.NumCPU())

	/* sync wait group */

	//withOutGoRoutine()

	//goRoutineWithoutWaitGroup()

	//goroutineWithSyncWaitGroup()

	workerPool()

	//selectConcurrency()
}
