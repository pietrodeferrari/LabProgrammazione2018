package main

import "fmt"

func main() {
	var a int = 1<<2
	var b float64 = 12.5
	var e // var e int
	var c int

	c := a + b // c = a + int(b)
	d := a/b // d := float64(a)/b

	fmt.Println(a, b, c, d, f) // fmt.Println(a, b, c, d, e)
}
