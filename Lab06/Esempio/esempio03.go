package main

import "fmt"

func test(x int) y int, z int // func test(x int) (y int, z int) {
{ // la parentesi graffa va sulla stessa riga
	y := x + 1 // y = x+1
	z := x + 2 // z = x+2
	return
}

func main() {
	var a, b int // da rimuovere
	a, b := test(10)
	fmt.Println(a, b)
}
