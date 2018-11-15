package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Introduci la riga di testo:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str := scanner.Text()

		risultato := ""
		for _, l := range str {
			// Prima versione: maiuscola -> minuscola
			if l >= 'A' && l <= 'Z' {
				l = l - 'A' + 'a'
			}
			// Seconda versione: minuscola -> maiuscola
			//if l >= 'a' && l <= 'z' {
			//	l = l - 'a' + 'A'
			//}
			risultato += string(l)
		}

		fmt.Println(risultato)
	}

}
