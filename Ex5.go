package main

import "fmt"

type ketan int

var x ketan
var y int

func main() {

	fmt.Println(x)

	fmt.Printf("%T", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T", y)

}
