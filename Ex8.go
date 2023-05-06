package main

import "fmt"

func main() {
	a := (25 == 25)
	b := (25 != 95)
	c := (25 <= 89)
	d := (25 >= 26)
	e := (25 < 90)
	f := (25 > 12)

	fmt.Println(a, b, c, d, e, f)

}
