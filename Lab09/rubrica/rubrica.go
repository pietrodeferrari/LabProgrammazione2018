package main

import "fmt"

func main() {

	var rubrica Rubrica

	rubrica = aggiungiContatto(rubrica,"Mario", "Rossi", "02503111", "Via Festa del Perdono", 11, "20122", "Milano")
	rubrica = aggiungiContatto(rubrica,"Anna", "Rossi", "02503111", "Via Festa del Perdono", 11, "20122", "Milano")
	rubrica = aggiungiContatto(rubrica,"Carlo", "Rossi", "02503111", "Via Festa del Perdono", 11, "20122", "Milano")

	stampaRubrica(rubrica)

	rubrica = eliminaContatto(rubrica, "Mario", "Rossi")

	stampaRubrica(rubrica)

	rubrica = aggiornaContatto(rubrica, "Anna", "Rossi", "Via S. Sofia", 25, "20122", "Milano")

	stampaRubrica(rubrica)
}

type Indirizzo struct {
	via, cap, città string
	numero uint
}

type Contatto struct {
	nome, cognome string
	telefono string
	indirizzo Indirizzo
}

type Rubrica []Contatto

func stampaIndirizzo(indirizzo Indirizzo) {
	fmt.Printf("%s %d, %s, %s", indirizzo.via, indirizzo.numero, indirizzo.cap, indirizzo.città)
}

func stampaContatto(persona Contatto) {
	fmt.Printf("%s %s, Tel. %s, ", persona.nome, persona.cognome, persona.telefono)
	stampaIndirizzo(persona.indirizzo)
}

func stampaRubrica(rubrica Rubrica) {
	fmt.Println("Lista dei", dimensioneRubrica(rubrica), "contatti")
	for i, p := range rubrica {
		fmt.Printf("[%d] - ", i+1)
		stampaContatto(p)
		fmt.Println()
	}
}

func dimensioneRubrica(rubrica Rubrica) int {
	return len(rubrica)
}

func aggiungiContatto(rubrica Rubrica, nome, cognome, telefono, via string, numero uint, cap, città string) Rubrica {
	return append(rubrica, Contatto{nome, cognome, telefono,
		Indirizzo{via, cap, città, numero}})
}

func eliminaContatto(rubrica Rubrica, nome, cognome string) Rubrica {
	for i, contatto := range rubrica {
		if contatto.nome == nome && contatto.cognome == cognome {
			rubrica = append(rubrica[:i], rubrica[i+1:]...)
		}
	}
	return rubrica
}

func aggiornaContatto(rubrica Rubrica, nome, cognome, via string, numero uint, cap, città string) Rubrica {
	for i, contatto := range rubrica {
		if contatto.nome == nome && contatto.cognome == cognome {
			rubrica[i].indirizzo = Indirizzo{via, cap, città, numero}
		}
	}
	return rubrica
}