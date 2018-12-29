package main

import "fmt"

func main() {

	var a []int = make([]int,5)

	fmt.Println(a)

	f(a)

	fmt.Println(a)

	a = g(a)

	fmt.Println(a)
}

func f(s []int) {
	s[0] = 10
	s = append(s, 10)
	fmt.Println("s:",s)
}

func g(s []int) []int {
	s[0] = 20
	s = append(s, 20)
	fmt.Println("s:",s)
	return s
}