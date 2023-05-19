package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5}

	s := sum(ii...)

	fmt.Println(s)

	o := odd(sum, ii...)

	fmt.Println("Sum of odd elements", o)

}

func sum(xi ...int) int {
	total := 0
	for i := 0; i < len(xi); i++ {
		total = total + xi[i]
	}
	return total
}

func odd(f func(xi ...int) int, vi ...int) int {

	var yi []int

	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)

}
