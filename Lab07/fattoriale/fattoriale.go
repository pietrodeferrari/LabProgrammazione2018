package main

import "fmt"

func main() {

	var numero uint

	fmt.Println("Inserisci un numero:")
	fmt.Scan(&numero)

	fattoriale(numero)
}

func fattoriale(n uint) (fattoriale uint) {
	fattoriale = 1
	for i:=uint(1); i<=n; i++ {
		fattoriale *= i
		fmt.Printf("Il fattoriale di %d è %d\n", i, fattoriale)
	}
	return
}
