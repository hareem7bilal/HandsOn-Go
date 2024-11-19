// generics/standard-library/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sumInts(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

// create a numeric interface with a type set
type numeric interface {
	constraints.Integer | constraints.Float
	//grow()
}

// update sum function to use a numeric interface with a type set
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func (si specialInt) grow() {}

// equal returns true if a and b are equal.
func equal[T comparable](a, b T) bool {
	return a == b

}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(2.5, 2.2))

	// invoke equal with comparable types
	fmt.Println("equal(1, 1):", equal(1, 1))
	fmt.Println("equal(\"one\", \"two\"):", equal("one", "two"))

	// invoke equal with a custom type
	type c struct{ f string }
	fmt.Println("equal(c{f: \"a\"}, c{f: \"a\"}):", equal(c{f: "a"}, c{f: "a"}))
	fmt.Println("equal(c{f: \"a\"}, c{f: \"b\"}):", equal(c{f: "a"}, c{f: "b"}))
}
