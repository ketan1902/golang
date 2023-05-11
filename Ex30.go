package main

import "fmt"

type vehical struct {
	doors string
	color string
}

type truck struct {
	vehical

	fourWheel bool
}

type sedan struct {
	vehical
	luxury bool
}

func main() {

	t := truck{
		vehical: vehical{
			doors: "2",
			color: "Black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehical: vehical{
			doors: "4",
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(s)
	fmt.Println(t)

	fmt.Println(s.vehical)
	fmt.Println(t.vehical)

}
