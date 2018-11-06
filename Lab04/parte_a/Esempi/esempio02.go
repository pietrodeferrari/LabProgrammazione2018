package main

import "fmt"

// Cosa produce in output il programma?
func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i *= 2 {
		fmt.Print(i, " ")
	}

	fmt.Println()

}
