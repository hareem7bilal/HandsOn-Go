// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")
	fmt.Println("Work in main...")

	// defer recovery
	defer func(){
		if err:=recover(); err!=nil{
			fmt.Println("Recovered from panic: ", err)
		}
	}()

	// panic
	panic("Something bad happened")
}
