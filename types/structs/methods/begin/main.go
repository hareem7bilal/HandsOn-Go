// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first , last  string
}

// fullName returns the full name of the author
func (a author) fullName() string{
	return a.first+" "+a.last
}

func main() {
	// initialize author
	a := author{
		first: "Marcus",
		last:  "Aurelius",
	}

	// print the author's full name
	fmt.Printf("%#v\n", a)
	fmt.Println(a.fullName())
}
