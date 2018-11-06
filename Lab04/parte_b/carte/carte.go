package main

import "fmt"

func main() {
	const ACE = '\U0001F0B1'
	const TEN = '\U0001F0B1' + 9

	for c := ACE; c <= TEN; c++ {
		fmt.Println("Simbolo:", string(c), "- Codice numerico in base 10:", c)
	}

}
