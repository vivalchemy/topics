package main

// package datetime

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for i := range 10000000 {
		// fmt.Println(i)
		go func() {
			i = i
		}()
	}
	timeTaken := time.Since(t)
	fmt.Println("The time taken is ", timeTaken)
}
