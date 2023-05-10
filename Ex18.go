package main

import "fmt"

func main() {

	arr := [5]int{10, 100, 1000, 10000, 100000}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)

	fmt.Println(arr)

}
