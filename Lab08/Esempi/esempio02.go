package main

import "fmt"

func main() {

	const DIMENSIONE = 10

	var a []int

	for i:=0; i<DIMENSIONE; i++ {
		a = append([]int{i}, a...)
	}

	fmt.Println(a)

	b := make([]int, DIMENSIONE)

	copy(b, a[DIMENSIONE/2:])

	fmt.Println(b)

	c := make([]int, DIMENSIONE)

	for i:=0; i<DIMENSIONE; i+=2 {
		copy(c[DIMENSIONE-2-i:], a[i:i+2])
	}

	fmt.Println(c)
}
