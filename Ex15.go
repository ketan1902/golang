package main

import "fmt"

func main() {
	a := 25

	if a == 25 {
		fmt.Println("Correct")
	} else if a > 20 {
		fmt.Println("THis is also correct")
	} else {
		fmt.Println("Wrong")
	}
}
