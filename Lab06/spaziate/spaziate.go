package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	// Come per l'esercizio sillabazione, questo codice legge un testo e lo converte
	// (aggiungendo uno spazio tra i caratteri)

	fmt.Println("Inserisci un testo:")
	testo := leggiTesto()

	fmt.Println("Testo a doppia spaziatura:")
	fmt.Println(aggiungiSpaziatura(testo))
}

// La funzione è identica a quella per l'esercizio sillabazione e legge un testo da tastiera
// (Vedi esercizio sillabazione/sillabazione.go per ulteriori commenti)
func leggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

// La funzione aggiunge uno spazio tra ogni coppia di caratteri quando almeno uno di questi è diverso da ' ' (spazio)
// La funzione riceve come argomento un testo e restituisce la stringa trasformata
func aggiungiSpaziatura(testo string) (risultato string) {
	// Come per l'esercizio sulla sillabazione, uso una variabile di supporto per
	// sapere se il carattere precedente fosse un spazio
	èSpazio := true
	for _, l := range testo {
		// Se il carattere precedente non era uno spazio e non lo è nemmeno quello corrente
		if !èSpazio && !unicode.IsSpace(l) {
			// aggiungi uno spazio
			risultato += " "
		}
		// Aggiorna il risultato e cambia il valore della variabile èSpazio in base al carattere corrente
		risultato += string(l)
		èSpazio = unicode.IsSpace(l)
	}
	return
}
