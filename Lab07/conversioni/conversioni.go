package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for _, v := range os.Args[1:] {
		convertiBool(v)
		convertiInt(v)
		convertiUInt(v)
		convertiFloat(v)
	}

}

func convertiBool(argomento string) {
	stampaRiga()
	defer stampaRiga()

	fmt.Println("CONVERTO IN VALORE BOOLEANO")

	if valore, err := strconv.ParseBool(argomento); err == nil {
		fmt.Println("valore boolean =", valore)
	} else {
		fmt.Printf("NON è bool (errore: %v)\n", err)
	}
}

func convertiInt(argomento string) {
	stampaRiga()
	defer stampaRiga()

	sizes := []int{strconv.IntSize, 32, 64}

	fmt.Println("CONVERTO IN VALORE INT")

	for i, sizeToConvert := range sizes {
		sizeToDisplay := strconv.Itoa(sizeToConvert)
		if i == 0 {
			sizeToDisplay = ""
		}
		if valore, err := strconv.ParseInt(argomento, 10, sizeToConvert); err == nil {
			fmt.Printf("valore int%s = %v\n", sizeToDisplay, valore)
		} else {
			fmt.Printf("NON è int%s (errore: %v)\n", sizeToDisplay, err)
		}
	}
}

func convertiUInt(argomento string) {
	stampaRiga()
	defer stampaRiga()

	sizes := []int{8, 16, 32, 64}

	fmt.Println("CONVERTO IN VALORE UINT")

	for _, s := range sizes {
		if valore, err := strconv.ParseUint(argomento, 10, s); err == nil {
			fmt.Printf("valore uint%d = %v\n", s, valore)
		} else {
			fmt.Printf("NON è uint%d (errore: %v)\n", s, err)
		}
	}
}

func convertiFloat(argomento string) {
	stampaRiga()
	defer stampaRiga()

	sizes := []int{32, 64}

	fmt.Println("CONVERTO IN VALORE FLOAT")

	for _, s := range sizes {
		if valore, err := strconv.ParseFloat(argomento, s); err == nil {
			fmt.Printf("valore float%d = %v\n", s, valore)
		} else {
			fmt.Printf("NON è float%d (errore: %v)\n", s, err)
		}
	}
}

func stampaRiga() {
	fmt.Println(strings.Repeat("-", 10))
}
