package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	s1 := s[2:3]
	s2 := s[1:6]
	s3 := s[:8]
	s4 := s[2:8]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
