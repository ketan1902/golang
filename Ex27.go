package main

import "fmt"

type person struct {
	first   string
	last    string
	fav_ice string
}

func main() {

	p1 := person{
		first:   "Ram",
		last:    "Gaikwad",
		fav_ice: "Butterscotch",
	}

	p2 := person{
		first:   "Sham",
		last:    "More",
		fav_ice: "Vanila",
	}

	s := []string{p1.fav_ice, p2.fav_ice}

	for i := range s {
		fmt.Println(s[i])
	}
}
