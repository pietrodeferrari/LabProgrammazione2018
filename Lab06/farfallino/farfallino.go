package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inserisci il testo:")

	// Rispetto all'esercizio 'garibaldi', in questo esercizio c'è una piccola diversità
	// nel modo in cui viene letto il testo:
	// in 'garibaldi' viene letta una riga di testo e la funzione è subito applicata a tale riga.
	// In questo esercizio invece, prima viene fatta tutta la lettura del testo, e solo dopo viene applicata la funzione

	// Queste 5 righe permettono di leggere un testo da tastiera e salvarlo in una variabile 'testo'
	var testo string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // In questo esempio, il codice for termina SOLO quando si inserisce il carattere terminatore (CTRL+D)
		testo += scanner.Text() + "\n" // aggiungo un carattere '\n' (a capo) perché questo viene filtrato
		// quando faccio la lettura con lo scanner
	}

	// Una volta letto il testo, per ogni carattere applicao la funzione toFarfallino() e stampo il risultato
	for _, r := range testo {
		fmt.Print(toFarfallino(r))
	}

}

// La funzione prende in input un carattere e stampa la sua conversione (come tipo string)
func toFarfallino(ch rune) string {
	sch := string(ch) // converto il carattere in string
	switch ch {
	case 'A', 'E', 'I', 'O', 'U':
		return sch + "F" + sch // in questo caso differenzio tra minuscole e maiuscola, aggiungendo una 'f' minuscola
	case 'a', 'e', 'i', 'o', 'u':
		return sch + "f" + sch
	default:
		return sch
	}
}
