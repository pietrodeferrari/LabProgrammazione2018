package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	occorrenzeMaiuscole, occorrenzeMinuscole := make(map[rune]int), make(map[rune]int)

	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	for scanner.Scan() {
		testo += scanner.Text()
	}

	for _, l := range testo {
		if unicode.IsUpper(l) {
			occorrenzeMaiuscole[l]++
		} else if unicode.IsLower(l) {
			occorrenzeMinuscole[l]++
		}
	}

	fmt.Println("Occorrenze:")
	for k, v := range occorrenzeMaiuscole {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", v))
	}
	for k, v := range occorrenzeMinuscole {
		fmt.Printf("%c: %s\n", k, strings.Repeat("*", v))
	}
}
