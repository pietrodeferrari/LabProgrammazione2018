package main

import (
	"fmt"
	"strings"
)

type Persona struct {
	nome, cognome string
}

func main() {

	p := Persona{"Rick", "Sanchez"}

	fmt.Println(p)

	f(p)

	fmt.Println(p)

	g(&p)

	fmt.Println(p)
}


func f(p Persona) {
	p.nome, p.cognome = strings.ToUpper(p.nome), strings.ToUpper(p.cognome)
}

func g(p *Persona) {
	p.nome, p.cognome = strings.ToUpper(p.nome), strings.ToUpper(p.cognome)
}