package main

import "fmt"

var b int = 3
var a float64

func main() {
	var c int
	var d int = 4

	fmt.Println("Variabile globale non inizializzata =", a)
	fmt.Println("Variabile globale inizializzata =", b)
	fmt.Println("Variabile locale non inizializzata =", c)
	fmt.Println("Variabile locale inizializzata =", d)
	//fmt.Println("Variabile non dichiarata =", e)
}
