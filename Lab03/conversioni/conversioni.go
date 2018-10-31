/**
Convertitore

Permette di scegliere tra varie conversioni di grandezze numeriche
*/
package main

import "fmt"

func main() {
	var valore float64
	var scelta int = 0

	// questo ciclo permette di ripetere l'inserimento di una scelta fino a che non si inserisce un valore valido (da 1 a 8)
	for scelta < 1 || scelta > 8 {
		fmt.Print(
			"Scegli la conversione:\n" +
				"1) secondi -> ore\n" +
				"2) secondi -> minuti\n" +
				"3) minuti -> ore\n" +
				"4) minuti -> secondi\n" +
				"5) ore -> secondi\n" +
				"6) ore -> minuti\n" +
				"7) minuti -> giorni e ore\n" +
				"8) minuti -> anni e giorni\n: ")
		fmt.Scan(&scelta)
	}

	fmt.Print("Inserisci il valore da convertire: ")
	fmt.Scan(&valore)

	// in base alla scelta effettuata, si calcola la conversione
	if scelta == 1 {
		fmt.Print(valore, " secondi corrispondono a ", valore/3600, " ore\n")
	} else if scelta == 2 {
		fmt.Print(valore, " secondi corrispondono a ", valore/60, " minuti\n")
	} else if scelta == 3 {
		fmt.Print(valore, " minuti corrispondono a ", valore/60, " ore\n")
	} else if scelta == 4 {
		fmt.Print(valore, " minuti corrispondono a ", valore*60, " secondi\n")
	} else if scelta == 5 {
		fmt.Print(valore, " ore corrispondono a ", valore*3600, " secondi\n")
	} else if scelta == 6 {
		fmt.Print(valore, " ore corrispondono a ", valore*60, " minuti\n")
	} else if scelta == 7 {
		giorni := int(valore) / (60 * 24)
		fmt.Print(valore, " minuti corrispondono a ", giorni, " giorni e ", valore/60-float64(24*giorni), " ore\n")
	} else if scelta == 8 {
		anni := int(valore) / (60 * 24 * 365)
		fmt.Print(valore, " minuti corrispondono a ", anni, " anni e ", valore/(60*24)-float64(365*anni), " giorni\n")
	}
}
