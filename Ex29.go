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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		for k, val := range v.fav_ice {
			fmt.Println(k, val)
		}
	}
}
