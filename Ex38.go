package main

import "fmt"

func main() {
	foo()

}

func foo() {
	fmt.Println("This is from foo")

	func() {
		fmt.Println("This is from anonymous function")
	}()
}
