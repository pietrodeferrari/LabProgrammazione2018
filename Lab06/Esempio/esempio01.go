package main

import "fmt"

func test1(x int) (y int) {
	// Quando la funzione viene chiamata, x = 10 e y = 0
	y = x * 10 // x = 10 e y = 100
	x, y = y, x // x = 100 e y = 10
	return // restituisce il valore di y, quindi 10
}

func main() {

	var a, b int = 10, 5
	b = test1(a) // Richiamo test1 passando a = 10 come parametro

	fmt.Println(a, b) // Stampa 10 10
}
