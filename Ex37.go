package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length, width float64
}

type CIRCLE struct {
	radius float64
}

type SHAPE interface {
	area() float64
}

func main() {

	s1 := SQUARE{
		length: 4,
		width:  5,
	}

	c1 := CIRCLE{
		radius: 6,
	}

	info(s1)
	info(c1)

}

func (s SQUARE) area() float64 {
	return s.length * s.width
}

func (c CIRCLE) area() float64 {

	return c.radius * c.radius * math.Pi

}

func info(s SHAPE) {

	ans := s.area()

	fmt.Println(ans)

}
