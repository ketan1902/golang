package main

import "fmt"

func main() {

	x := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},

		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`}}

	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k)

		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
