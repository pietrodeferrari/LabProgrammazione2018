package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6}
	b := a // viene fatta una copia dell'intero array a
	c := a[:]
	d := a[2:]
	// c e d invece sono delle fette dell'array a
	// c va da a[0] fino a a[5]
	// d va da a[2] fino a a[5]

	// Stampa tutti i valori della fetta d,
	// ovvero i valori da a[2] a a[5]
	for i, v := range d {
		fmt.Print(i,": ",v,"\n")
	}
	fmt.Println()

	// Le modifiche sugli elementi di una fetta si ripercuotono
	// sugli elementi dell'array
	c[0] = c[1] + c[2]
	d[1] = d[0] + c[0]
	// questi assegnamenti portano ad avere a[0]=5 e a[3]=8

	// Stampa tutti i valori di a
	for i, v := range a {
		fmt.Print(i,": ",v,"\n")
	}
	fmt.Println()

	// Stampa tutti i valori di c (ovvero tutto i valori di a) ma su una sola riga
	for _, v := range c {
		fmt.Print(v," ")
	}
	fmt.Println()

	// Stampa tutti i valori di b, che non è stato modificato ed è quindi [1,2,3,4,5,6]
	for i := range b {
		fmt.Print(i,": ",b[i],"\n")
	}
	fmt.Println()

	// ad ogni iterazione, nella variabile v c'è una copia del valore di un elemento di b
	// modificando v quindi si modifica una copia, non l'elemento originale in b
	for _, v := range b {
		v *= 2
	}

	// Stampa nuovamente i valori di b, ancora [1,2,3,4,5,6]
	for i := 0; i<len(b); i++ {
		fmt.Print(i,": ",b[i],"\n")
	}

}

