package main

import "fmt"

func main() {

	ans := foo()

	fmt.Println(ans)

}

func foo() int {
	a := 5
	b := 10
	sum := a + b

	return sum

}
