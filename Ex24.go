package main

import "fmt"

func main() {

	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, name := range x {

		fmt.Println(i, name)

		for j, val := range name {
			fmt.Printf("%v, %v\n", j, val)
		}
	}

}
