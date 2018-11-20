package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Per leggere una stringa di testo è necessario usare lo Scanner:
	// la creazione di un nuovo Scanner si ottiene con la procedura NewScanner del package bufio
	// mentre Stdin è un file speciale (creato nel package os) che indirizza allo standard input
	scanner := bufio.NewScanner(os.Stdin)

	// La funzione Scan restituisce true quando riesce a leggere qualcosa e false altrimenti
	// Mettendola come condizione in un for significa:
	// "continua a leggere una riga di testo alla volta finché riesci"
	// Messa in un if invece, significa:
	// "se riesci a leggere qualcosa, allora fai ..."
	for scanner.Scan() {
		// Il metodo Text restituisce il testo letto dallo Scanner con tipo string
		riga := scanner.Text()

		// il programma deve terminare in due casi:
		// 1) si usa il carattere terminatore (su Linux CTRL+D)
		// 2) si inserisce una riga vuota
		// Il caso 1 è automatico: se si inserisce il carattere terminatore
		// la funzione Scan (messa nella condizione del for) restituisce false e il ciclo for termina.
		// Il caso 2 è da gestire, in questo caso con un if
		if riga != "" {
			// Se la riga non è vuota -> applica l'algoritmo che trasforma le vocali in 'u'
			fmt.Println(modificaRiga(riga))
		} else {
			// Se riga vuota -> termina
			fmt.Println("Riga vuota, termino.")
			break
		}
	}
}

// Questa funzione riceve in input un carattere e lo trasforma:
// se è una vocale la trasforma in 'u' (maiuscola o minuscola),
// altrimenti restituisce il carattere stesso
func modificaCarattere(ch rune) rune {
	switch ch{
	case 'a', 'e', 'i', 'o':
		// se io carattere è una vocale minuscola, restituisco una 'u'
		return 'u'
	case 'A', 'E', 'I', 'O':
		// se io carattere è una vocale maiuscola, restituisco una 'U'
		return 'U'
	default:
		// altrimenti restituisco il carattere stesso
		return ch
	}
}

// La funzione riceve in input una stringa e restituisce in output una nuova stringa dove,
// per ogni carattere, si applica la funzione modificaCarattere()
func modificaRiga(riga string) (conversione string) {
	for _, ch := range riga {
		conversione += string(modificaCarattere(ch))
	}
	return
}
