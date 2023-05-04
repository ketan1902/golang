package main

import "fmt"

type ketan int

var x ketan = 10

func main() {

	fmt.Println(x)

	fmt.Printf("%T", x)

	x = 42

	fmt.Println(x)

}
