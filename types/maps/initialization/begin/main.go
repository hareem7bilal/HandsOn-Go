// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	// var authors map[string]author

	// initialize the map with make
	// authors = make(map[string]author)

	// add authors to the map
	// authors["tm"] = author{
	// 	name: "Tony Morrison",
	// }

	// authors["ma"] = author{
	// 	name: "Marcus Aurelius",
	// }

	authors:=map[string]author{
		"tm": {name: "Tony Morrison"},
		"ma": {name: "Marcus Aurelius"},
		"jr": author{name: "Justine Rochester"},
	}

	// print the map with fmt.Printf
	fmt.Printf("%#v", authors)

	// read a value from the map with a known key
	fmt.Println("JR: ", authors["jr"])
}
