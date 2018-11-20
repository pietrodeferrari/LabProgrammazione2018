package main

import (
	"fmt"
)

func main() {
	var p1, p2 string

	// Leggo due parole
	fmt.Print("Prima stringa: ")
	fmt.Scan(&p1)
	fmt.Print("Seconda stringa: ")
	fmt.Scan(&p2)
	// La differenza tra la lettura con Scan e uno Scanner, è che la prima legge fino al primo spazio (o fine riga)
	// mentre lo scannere può leggere un intero testo fino al fine riga, compresi gli spazi

	fmt.Print("Massimo prefisso: '", prefisso(p1, p2), "'\n")
	fmt.Print("Massimo suffisso: '", suffisso(p1, p2), "'\n")
}

// La funzione riceve in input due stringhe e restituisce il massimo prefisso
// La funzione controlla carattere per carattere le due parole e si ferma nel momento in cui trova un carattere diverso
func prefisso(parola1, parola2 string) (risultato string) {
	// Prima calcolo la lunghezza della parola più corta: non ci può essere prefisso comune più lungo
	lunghezzaMinima := len(parola1)
	if lunghezzaMinima > len(parola2) {
		lunghezzaMinima = len(parola2)
	}

	// Uso un ciclo for per confrontare carattere per carattere le due parole
	for index := 0; index < lunghezzaMinima; index++ {
		// Per ogni indice da 0 a lunghezzaMinima (non ha senso andare oltre)
		if parola1[index] == parola2[index] { // confronto i caratteri corrispondenti all'indice index
			// se sono uguali, aggiorno il risultato (il carattere fa parte del prefisso)
			risultato += string(parola1[index])
		} else {
			// altrimenti il prefisso è terminato (ho trovato caratteri diversi)
			return
		}
	}
	return
}

// La funzione riceve in input due stringhe e restituisce il massimo suffisso
// La funzione controlla carattere per carattere le due parole e si ferma nel momento in cui trova un carattere diverso
// La funzione è identica alla precedente, con l'unica differenza che invece di andare in 'avanti' (dal carattere di indice 0,
// al carattere di indice lunghezzaMinima-1), va all'indietro
func suffisso(parola1, parola2 string) (risultato string) {
	lunghezzaMinima := len(parola1)
	if lunghezzaMinima > len(parola2) {
		lunghezzaMinima = len(parola2)
	}

	for index := 0; index < lunghezzaMinima; index++ {
		if parola1[len(parola1)-1-index] == parola2[len(parola2)-1-index] {
			risultato = string(parola1[len(parola1)-1-index]) + risultato
		} else {
			return
		}
	}
	return
}
