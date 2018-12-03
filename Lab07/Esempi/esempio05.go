package main

import "fmt"

func main() {
	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	a[1][1] = 10
	a[1][0] = a[1][1] + a[2][1]

	for i, riga := range a {
		fmt.Print(i, ": ")
		for _, cella := range riga {
			fmt.Print(cella, " ")
		}
		fmt.Println()
	}
}
