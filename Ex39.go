package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5}

	s := sum(ii...)

	fmt.Println(s)

}

func sum(xi ...int) int {
	total := 0
	for i := 0; i < len(xi); i++ {
		total = total + xi[i]
	}
	return total
}
