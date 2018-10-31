package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Inserisci i due numeri:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a < b {
		fmt.Println("Il maggiore è", b)
	} else if a > b {
		fmt.Println("Il maggiore è", a)
	} else {
		fmt.Println("I numeri sono uguali")
	}
}
