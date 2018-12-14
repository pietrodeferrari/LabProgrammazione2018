package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Inserisci stringa base:")
	base := leggiParole()
	fmt.Println("Inserisci stringa inverti:")
	inverti := leggiParole()

	fmt.Println("stringa invertita:", invertiParole(base, inverti))

}

func leggiParole() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func invertiParole(base, daInvertire []string) (risultato string) {
	paroleInvertite := make(map[string]string)

	for _, s := range base {
		paroleInvertite[s] = s
	}

	for _, s := range daInvertire {
		if _, ok := paroleInvertite[s]; ok {
			paroleInvertite[s] = invertiParola(s)
		}
	}

	for _, s := range base {
		risultato += paroleInvertite[s] + " "
	}

	return
}

func invertiParola(parola string) (stringaInvertita string) {
	for _, l := range parola {
		stringaInvertita = string(l) + stringaInvertita
	}
	return
}
