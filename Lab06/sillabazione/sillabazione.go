package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// il programma principale fa due cose:
	// 1) legge un testo
	// 2) chiava una funzione che lo converta e ne stampa il risultato
	// Suddividere il proprio codice in funzioni lo rende più leggibile e comprensibile
	testo := leggiTesto()
	fmt.Println(sillaba(testo))
}

// La funzione restituisce una variabile string contenente il testo su più righe letto da tastiera
// La lettura si interrompe alla digitazione del carattere terminatore (CTRL+D)
// Racchiudere del codice che usate frequentemente in funzioni, vi permette di riusarlo in modo molto rapido
func leggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// La variabile testo è inizializzata al valore di default "" (stringa vuota) direttamente alla chiamata della funzione
		testo += scanner.Text() + "\n"
	}
	return
}

// La funzione effettua la sillabazione basata sull'ordine delle lettere
// Riceve in input la stringa da sillabare e restituisce la stringa sillabata
func sillaba(testo string) (risultato string) {
	// inizializzo una variabile chiamata 'precedente':
	// questa variabile verrà usata per memorizzare il carattere che precede il carattere corrente
	// all'interno del ciclo for
	precedente := ' '

	// Per ogni carattere del testo dato in input
	for _, l := range testo {
		// Se è vero che:
		//  - il carattere precedente era una lettera
		//  - il carattere corrente è una lettera
		//  - il carattere precedente ha un codice superiore al codice del carattere corrente
		//    (ovvero, il carattere precedente compare dopo nell'alfabeto rispetto al carattere corrente)
		if unicode.IsLetter(precedente) && unicode.IsLetter(l) && precedente > l {
			// allora aggiungo un trattino per la sillabazione
			risultato += "-"
		}
		risultato += string(l) // aggiungo il carattere corrente alla stringa risultato
		precedente = l // aggiorno il valore della variabile 'precedente' che prende il valore del carattere
		// attuale per l'iterazione successiva
	}

	return
}
