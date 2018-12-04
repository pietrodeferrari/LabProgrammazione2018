package main

import (
	"fmt"
	"strings"
)

func main() {

	const DIMENSIONE = 10

	var a []int

	for i := 0; i < DIMENSIONE; i++ {
		a = append(a, i)
		fmt.Println("Iterazione", i, ":", a, len(a), cap(a))
	}


	fmt.Println(strings.Repeat("=", 10))
	fmt.Println(a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println(a, len(a), cap(a))
	a = a[:DIMENSIONE]

	b := append(a[DIMENSIONE/2:], a[:DIMENSIONE/2]...)

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b))
	a = a[:cap(a)]
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b))
	a = a[:DIMENSIONE]


	c := append(a[:DIMENSIONE/4], a[1 + 3*DIMENSIONE/4:]...)

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println(a, len(a), cap(a), b, len(b), cap(b), c, len(c), cap(c))
}

