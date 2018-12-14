package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numeri := leggiNumeri()

	fmt.Println("Valori inseriti:")
	fmt.Println(numeri)

	stampaConsecutivi(numeri)

}

func leggiNumeri() (numeri []int) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), " ") {
			if n, err := strconv.Atoi(s); err == nil && n > 0 {
				numeri = append(numeri, n)
			}
		}
	}
	return
}

func stampaConsecutivi(numeri []int) {
	fmt.Println("Terne consecutivamente ordinate:")
	for i:=0; i<len(numeri); i++ {
		for j:=i+1; j<len(numeri); j++ {
			for z:=j+1; z<len(numeri); z++ {
				if numeri[i] < numeri[j] && numeri[j] < numeri[z] {
					fmt.Printf("[%d<%d<%d]\n", numeri[i], numeri[j], numeri[z])
				}
			}
		}
	}
}