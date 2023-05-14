package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6}

	n := foo(xi...)

	fmt.Println(n)

	x2 := []int{11, 22, 33, 44}

	m := bar(x2)
	fmt.Println(m)

}

func foo(xi ...int) int {

	total := 0

	for _, v := range xi {
		total = total + v
	}

	return total

}

func bar(x2 []int) int {
	sum := 0

	for _, v := range x2 {
		sum = sum + v
	}

	return sum

}
