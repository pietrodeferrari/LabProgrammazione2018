package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Introduci la riga di testo:")

	// Leggo una riga di testo tramite lo scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str := scanner.Text()

		risultato := ""
		for i, l := range str {

			if i%2 == 1 && l >= 'A' && l <= 'Z' {
				// Se la lettera è in posizione dispari ed è maiuscola, la trasformo in minuscola
				l = l - 'A' + 'a'
			} else if i%2 == 0 && l >= 'a' && l <= 'z' {
				// Se la lettera è in posizione pari ed è minuscola, la trasformo in maiuscola
				l = l - 'a' + 'A'
			}
			risultato += string(l)
		}

		fmt.Println(risultato)
	}
}
