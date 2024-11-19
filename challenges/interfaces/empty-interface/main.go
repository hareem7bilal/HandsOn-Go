package main

import "fmt"

func whatAmI(val interface{}) string {
	return fmt.Sprintf("%v %T", val, val)
}

func main(){
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
}