package main

import "fmt"

const (
	a     = 10
	b int = 45
)

func main() {
	a := 25

	fmt.Println(a, b)
	fmt.Printf("%T %T", a, b)
}
