package main

import "fmt"

type person struct {
	first   string
	last    string
	fav_ice []string
}

func main() {

	p1 := person{
		first: "Ram",
		last:  "Gaikwad",
		fav_ice: []string{
			"Butterscotch",
			"Blueberry",
			"Mango",
		},
	}

	p2 := person{
		first: "Sham",
		last:  "More",
		fav_ice: []string{
			"Vanila",
			"Strawberry",
			"Pista",
		},
	}

	for i, v := range p1.fav_ice {
		fmt.Println(i, v)
	}

	for i, v := range p2.fav_ice {
		fmt.Println(i, v)
	}
}
