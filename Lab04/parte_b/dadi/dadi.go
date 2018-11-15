/*
Simula una partita a dadi

Il numero di giocatori è fissato a 5. I dadi usati sono dadi a 6 facce. Il numero di turni è chiesto come input.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Uso delle costanti per mantenere il numero di facce del dado e la lista di giocatori
	const FACCEDADO = 6
	const GIOCATORI = "ABCDE" // Gestisco la lista dei giocatori come string, così da poter iterare su di essa

	var turni int
	vincitori := "" // Uso una variabile string per memorizzare la lista dei vincitori
	// (ci sono altre strutture dati più indicate, ma per il momento usiamo una string)

	// Inizializzo il seme del generatore di numeri casuali,
	// indispensabile per avere un risultato diverso per esecuzioni differenti
	rand.Seed(int64(time.Now().Nanosecond()))

	fmt.Print("Inserisci il numero di turni di gioco: ")
	fmt.Scan(&turni)

	for i := 1; i <= turni; i++ {
		// Queste variabili memorizzano il vincitore del turno e il punteggio migliore
		vincitoreTurno, lancioMigliore := ' ', 0

		// Per ogni giocatore...
		for _, giocatore := range GIOCATORI {
			// ... lancia due dati e memorizza il risultato
			lancio1, lancio2 := rand.Intn(FACCEDADO)+1, rand.Intn(FACCEDADO)+1

			fmt.Print("Turno ", i, ") Giocatore ", string(giocatore), ", lanci di valore: ", lancio1, " e ", lancio2, "\n")

			// Se il totale dei lanci è maggiore del valore di lancioMigliore (ho fatto un punteggio migliore)...
			if totale := lancio1 + lancio2; totale > lancioMigliore {
				// ... aggiorno i valori delle variabili
				vincitoreTurno, lancioMigliore = giocatore, totale
			}
		}
		fmt.Print("FINE TURNO ", i, " - GIOCATORE VINCITORE: ", string(vincitoreTurno), "\n")
		vincitori += string(vincitoreTurno)
	}

	// Al termine dei turni, per ogni giocatore conto quante volte ha vinto (quante volte compare in vincitori)
	fmt.Println("Nome e numero vittorie di ogni giocatore:")
	for _, giocatore := range GIOCATORI {
		count := 0
		for _, vincitore := range vincitori {
			if vincitore == giocatore {
				count++
			}
		}
		fmt.Println(string(giocatore), "->", count)
	}
}
