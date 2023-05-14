package main

import "fmt"

func main() {

	ans := foo()

	fmt.Println(ans)

	ans2, ans3 := bar()
	fmt.Println(ans2, ans3)

}

func foo() int {
	a := 5
	b := 10
	sum := a + b

	return sum

}

func bar() (int, string) {
	x := 10
	y := "This is ketan"

	return x, y
}
