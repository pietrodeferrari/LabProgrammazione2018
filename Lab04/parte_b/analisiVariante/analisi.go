package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	vmaiuscole, vminuscole := 0, 0
	cmaiuscole, cminuscole := 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str := scanner.Text()

		for _, l := range str {
			if unicode.IsLetter(l) {
				if unicode.IsUpper(l) {
					switch l {
					case 'A', 'E', 'I', 'O', 'U':
						vmaiuscole++
					default:
						cmaiuscole++
					}
				} else {
					switch l {
					case 'a', 'e', 'i', 'o', 'u':
						vminuscole++
					default:
						cminuscole++
					}
				}

			}

		}
	}

	fmt.Println("Vocali maiuscole:", vmaiuscole)
	fmt.Println("Consonanti maiuscole:", cmaiuscole)
	fmt.Println("Vocali minuscole:", vminuscole)
	fmt.Println("Consonanti minuscole:", cminuscole)
}
