// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare a signal channel to read from
	// readChan := make(chan struct{})
	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)
	// use a default case in select to prevent blocking
	// select {
	// case <-readChan:
	// 	fmt.Println("Recieved from readChan")
	// default:
	// 	fmt.Println("No signal recieved")
	// }
	for {
		select {
		case <-timeChan1:
			fmt.Println("timeChan1")
			return
		case <-timeChan2:
			fmt.Println("timeChan2")
			return
		default:
			fmt.Println("default")
			time.Sleep(250 * time.Millisecond)
		}

	}
}
