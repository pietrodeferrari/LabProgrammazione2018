package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, d := leggiInput()
	if d >= len(n) {
		fmt.Println(math.NaN())
	} else {
		fmt.Println("numero migliore:", numeroMigliore(n, d))
	}
}

func leggiInput() (n string, d int) {
	n = os.Args[1]
	d, _ = strconv.Atoi(os.Args[2])
	return
}

func numeroMigliore(n string, d int) (minimo int) {
	minimo, _ = strconv.Atoi(n)

	for i:=0; i < int(math.Pow(2, float64(len(n)))); i++ {
		formato := fmt.Sprintf("%%0%db", len(n))
		binario := fmt.Sprintf(formato, i)
		if strings.Count(binario,"0") == d {
			minimo = testNumero(minimo, filtraStringa(n, binario))
		}
	}

	return
}

func filtraStringa(numero, binario string) (risultato string){
	for i, l := range numero {
		if binario[i] != '0' {
			risultato += string(l)
		}
	}
	return
}

func testNumero(minimo int, daTestare string) int {
	if valore, err := strconv.Atoi(daTestare); err == nil && minimo > valore {
		minimo = valore
	}
	return minimo
}
