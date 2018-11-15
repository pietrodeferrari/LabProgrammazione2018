package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	// Sono necessari 4 contatori
	maiuscole, minuscole := 0, 0
	consonanti, vocali := 0, 0

	// Creo uno scanner per leggere una riga di testo
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// Se lo scanner ha letto qualcosa, allora recupero il testo letto tramite il metodo Text()
		str := scanner.Text()

		// Uso il for range per prendere in considerazione ogni lettera del testo
		for _, l := range str {

			// Se è una lettera (avrei potuto usare anche i codici ASCII)
			if unicode.IsLetter(l) {
				// Incremento il contatore corretto (maiuscole se l è una maiuscola e minuscole altrimenti)
				if unicode.IsUpper(l) {
					maiuscole++
				} else {
					minuscole++
				}

				// Trasformo la lettera in una lettera minuscola:
				// in questo modo controllo solo le vocali minuscole
				// (non devo controllare che sia 'A' o 'a', 'E' o 'e', 'I' o 'I', ecc.)
				l = unicode.ToLower(l)

				// Incremento il contatore adeguato
				switch l {
				// Questa versione dello switch incrementa il contatore vocali se la variabile l
				// è una di queste 5 lettere
				case 'a', 'e', 'i', 'o', 'u':
					vocali++
				default:
					consonanti++
				}
			}

		}
	}

	fmt.Println("Maiuscole:", maiuscole)
	fmt.Println("Minuscole:", minuscole)
	fmt.Println("Vocali:", vocali)
	fmt.Println("Consonanti:", consonanti)
}
