package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if input := scanner.Text(); input == "" {
			break
		} else {
			str += input + "\n"
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
}
