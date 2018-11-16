package main

import "fmt"

func main() {
	var n float32

	fmt.Scan(&n)

	var n64 = float64(n)

	for n < n+1 && n64 < n64+1 {
		fmt.Println(n, n64)
		n++
		n64++
	}

	fmt.Println(n, n64)
	fmt.Println("Overflow: ", n+1, n64+1)
}
