// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string{
	return "Konnichiwa!"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string{
	return "Konnichiwa, "+name+"!"
}

// greetWithNameAndAge returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) (string) {
	greeting := "Konnichiwa, watashi no namae wa "+name+" desu, sotishe watashi wa "+strconv.Itoa(age)+" toshi desu."
	return greeting
}

func greetWithNameAndAge2(name string, age int) (greeting string) {
	greeting = "Konnichiwa, watashi no namae wa "+name+" desu, sotishe watashi wa "+strconv.Itoa(age)+" toshi desu."
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(a,b int) (int, error){
	if(b==0){
		return 0, errors.New("Cannot divide by 0")
	}
	return a/b, nil
}

func main() {
	// invoke greet function
	// fmt.Println(greetWithNameAndAge("Hareem", 23))

	// invoke greetWithName function
	// fmt.Println(greetWithName("Toni"))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(4, 0))
}
