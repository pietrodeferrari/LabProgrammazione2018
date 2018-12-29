package main

import "fmt"

type Persona struct {
	nome string
	cognome string
}

func main() {

	anagrafica := []Persona{
		{"Rick", "Sanchez"},
		{"Morty", "Smith"},
		{"Jerry", "Smith"},
		{"Summer", "Smith"},
		{"Beth", "Smith"}}

	fmt.Printf("Tipo anagrafica %T\n", anagrafica)
	fmt.Println()

	for _, p := range anagrafica {
		fmt.Println(p.nome, p.cognome)
		fmt.Println()

		anagrafica[0].cognome = "Martinez"

		for _, p := range anagrafica {
			fmt.Println(p.nome, p.cognome)
		}
	}
}