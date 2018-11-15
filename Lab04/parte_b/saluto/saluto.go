package main

import (
	"fmt"
	"time"
)

func main() {

	var nome string

	// La lettura si una singola stringa può essere fatta con una semplice Scan
	// Per la lettura di testi più lunghi si consiglia l'uso di uno Scanner
	fmt.Println("Come ti chiami?")
	fmt.Scan(&nome)

	fmt.Println("Ciao", nome, "sai che oggi è il", time.Now())
}
