package main

import "fmt"

type Persona struct {
	nome string
	cognome string
}

func main() {

	var p1, p2, p5 Persona

	p1 = Persona{"Rick", "Sanchez"}
	p2.nome = "Mario"
	p2.cognome = "Rossi"

	p3 := Persona{"Morty", "Smith"}

	p4 := Persona{cognome: "Bianchi", nome: "Paolo"}

	p6 := Persona{cognome: "Verdi"}

	fmt.Println("p1:",p1)
	fmt.Println("p5:",p5)

	fmt.Printf("Nome: %-10s - Cognome: %-10s\n", p1.nome, p1.cognome)
	fmt.Printf("%T\t%v\n", p3, p3)
	fmt.Printf("%#v\n", p4)
	fmt.Printf("%#v\n", p6)

}
