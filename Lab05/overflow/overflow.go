package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var n32 int32 = int32(n)
	var n64 int64 = int64(n)

	for n < n+1 && n32 < n32+1 && n64 < n64+1 {
		fmt.Println(n, n32, n64)
		n++
		n32++
		n64++
	}

	fmt.Println(n, n32, n64)
	fmt.Println("Overflow: ", n+1, n32+1, n64+1)
}
