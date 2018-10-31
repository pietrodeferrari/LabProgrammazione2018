package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
