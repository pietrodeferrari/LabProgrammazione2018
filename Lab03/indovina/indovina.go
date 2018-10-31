// L'utente deve inserire un numero da tastiera
// per cercare di indovinare il numero fissato nel codice
package main

import "fmt"

func main() {
	var daindovinare, x int = 10, 0

	fmt.Println("Inserisci un tentativo:")
	fmt.Scanln(&x)

	if x == daindovinare {
		fmt.Println("Indovinato")
	} else if x < daindovinare {
		fmt.Println("Troppo piccolo")
	} else {
		fmt.Println("Troppo grande")
	}
}
