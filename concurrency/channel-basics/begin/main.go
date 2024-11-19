// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	// invoke the sum function as a goroutine
	go sum(nums, ch)

	result := <-ch

	fmt.Println("Result:", result)

	// force main thread to sleep
	//time.Sleep(100 * time.Millisecond)
	ch2 := make(chan string, 2) // buffered channel
	ch2 <- "Hitori"
	ch2 <- "Yamaguchi"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
