package main

import "fmt"

type humanoid interface {
	speak()
	walk()
}

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Printf("%s speaking!\n", p.name)
}
func (p person) walk() {
	fmt.Printf("%s walking!\n", p.name)
}
func (p person) String() string{
	return fmt.Sprintf("Hello, my name is %s.\n", p.name)
}

func (d dog) walk() {
	fmt.Printf("Dog walking!\n")
}

func (d dog) speak() {
	fmt.Printf("Dog speaking!\n")
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}

type dog struct{}

func main() {
	p := person{name: "Naomi"}
	doHumanThings(p)

	doHumanThings(dog{})
	fmt.Println(p)

}
