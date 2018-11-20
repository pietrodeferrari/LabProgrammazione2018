package main

import (
	"fmt"
)

var a int = 10

func test1() int {
	// test1 somma 5 al valore della variabile globale a e restituisce il suo valore
	a += 5
	return a
}

func test2(a int) int {
	// test2 somma 6 al valore della variabile locale a e restituisce il suo valore
	a += 6
	return a
}

func test3() int {
	// test3 restituisce la somma tra il valore della variabile globale a e 7 (a non viene modificata)
	return a + 7
}

func main() {
	var a, b, c int // Inizializzati ai valori di default a=0, b=0 e c=0
	a, b, c = test1(), test2(a), test3() // Richiamo le 3 funzioni. Questa riga equivale al seguente codice:
	// temp_a = test1()
	// temp_b = test2(a)
	// temp_c = test3()
	// a, b, c = temp_a, temp_b, temp_c
	// Prima vengono eseguite tutte le funzioni sul lato destro dell'assegnamento.
	// Successivamente, i valori di ritorno vengono assegnati alle variabili sul lato sinistro

	fmt.Println(a, b, c)
}
