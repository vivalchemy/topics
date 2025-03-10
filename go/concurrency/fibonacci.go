package main

import (
	"fmt"
	"runtime"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// go worker(jobs, results) // single cpu

	// multiple cpus
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	// utilize all cores
	// get the count of cores
	cores := runtime.GOMAXPROCS(runtime.NumCPU())
	for range cores {
		go worker(jobs, results)
	}

	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	for i := range 100 {
		jobs <- i
	}
	close(jobs) // no more jobs

	for range 100 {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	switch {
	case n < 2:
		return n
	default:
		return fib(n-1) + fib(n-2)
	}
}
