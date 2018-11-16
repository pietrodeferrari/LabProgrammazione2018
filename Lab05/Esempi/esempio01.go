package main

import "fmt"

func main() {

	var x int
	var y int = 'a'
	x = 'A'
	fmt.Print(string(x), " ")
	for x := 1; x < 10; x++ {
		fmt.Print(string(y+x), " ")
	}
	fmt.Println()
}
