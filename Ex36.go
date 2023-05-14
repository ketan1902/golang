package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "my age is", p.age)

}

func main() {

	p1 := person{
		first: "Ketan",
		last:  "Gaikwad",
		age:   23,
	}
	p1.speak()

}
