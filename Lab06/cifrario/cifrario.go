package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Dichiaro uno Scanner per leggere una riga di testo
	fmt.Print("Frase da cifrare: ")
	scanner := bufio.NewScanner(os.Stdin)

	// Uso un if solamente per essere sicuro di eseguire la cifratura SOLO se ho effettivamente letto qualcosa
	// ma la funzione Scan può essere usata anche da sola, ignorando il valore di ritorno
	if scanner.Scan() {
		riga := scanner.Text()

		// Leggo una chiave numerica
		var chiave byte
		fmt.Print("Chiave: ")
		fmt.Scan(&chiave)

		// Eseguo la funzione che cifra la stringa di testo e stampo il valore restituito
		fmt.Println("Frase cifrata:", cifraRiga(riga, chiave))
	}
}

// La funzione riceve come input una stringa e una chiave numerica e restituisce in output
// la stringa cifrata usando la chiave data in input
func cifraRiga(riga string, chiave byte) (rigaCifrata string) {
	for _, ch := range riga {
		rigaCifrata += string(cifraCarattere(byte(ch), chiave))
	}
	return
}

// La funzione riceve in input un carattere e la chiave numerica e restituisce il carattere cifrato
func cifraCarattere(ch, chiave byte) byte {
	// Se il carattere è una lettera maiuscola e gli sottraiamo 'A', otteniamo un numero che rappresenta
	// il suo indice all'interno dell'alfabeto (partendo da 0, 'A'=0, 'B'=1, 'C'=2,...)
	// Idem per le lettere minuscole se gli sottraiamo 'a'
	// Applicando poi la chiave otteniamo l'indice di un'altra lettera dell'alfabeto,
	// e convertiamo l'indice in lettera sommando 'A' (o 'a' per le lettere minuscole)
	// L'operatore % (che restituisce il resto della divisione) viene usato per essere sicuri che
	// dopo aver sommato indice e chiave si ottenga comunque un numero compreso tra 0 e 25, ovvero un numero
	// che sia indice delle 26 lettere dell'alfabeto inglese
	switch {
	case ch >= 'A' && ch <= 'Z':
		return (ch-'A'+chiave)%26 + 'A'
	case ch >= 'a' && ch <= 'z':
		return (ch-'a'+chiave)%26 + 'a'
	default:
		return ch
	}
}
