package main

import "fmt"

func main() {

	var x,y int

	fmt.Println("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)

	max := x
	if max < y {
		max = y
	}

	min := x
	if min > y {
		min = y
	}

	fmt.Println("Maggiore:", max)
	fmt.Println("Minore:", min)

	fmt.Println("Somma:", x+y)
	fmt.Println("Differenza:", max - min)
	fmt.Println("Prodotto:", x*y)
	fmt.Println("Divisione:", x/y)

	potenza := 1
	for i:=0; i<y; i++ {
		potenza *= x
	}
	fmt.Println("Potenza:", potenza)

	fmt.Println("Valore medio:", (x+y)/2)
}

