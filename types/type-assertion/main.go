// types/variables/begin/main.go
package main

import "fmt"

func whatAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is an int", val)
	case string:
		return fmt.Sprintf("%v is a string", val)
	default:
		return fmt.Sprintf("%v is not a supported type", val)

	}
}

func main() {
	var i any = "hello"
	i = 1
	//fmt.Println(i.(string))

	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	} else {
		fmt.Printf("%v is an int\n", i)
	}

	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI(true))

}
