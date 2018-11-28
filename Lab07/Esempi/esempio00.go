package main

import "fmt"

func main() {

	var a int = 10
	var b float64 = 12.5
	var c byte = 'A'

	// Le prime tre chiamate non modificano in alcun modo i valori delle 3 variabili
	fInt(a)
	fFloat(b)
	fByte(c)
	fmt.Printf("%d %f %c\n",a, b, c) // 10 12.500000 A

	// Le chiamate di funzioni con ritorno di un valore permettono invece di modificare
	// il valore di una variabile solo se il valore ritornato dalla funzione è assegnato
	// alla variabile
	a = fIntRet(a)
	b = fFloatRet(b)
	c = fByteRet(c)
	fmt.Printf("%d %f %c\n",a, b, c) // 20 22.500000 K

	// Passando ad una funzione un puntatore ad una variabile, il valore di quest'ultima può
	// essere modificato direttamente all'interno della funzione chiamata
	fIntPtr(&a)
	fFloatPtr(&b)
	fBytePtr(&c)
	fmt.Printf("%d %f %c\n",a, b, c) // 30 32.500000 U
}

func fInt(x int) {
	x = x + 10
}

func fFloat(x float64) {
	x = x + 10
}

func fByte(x byte) {
	x = x + 10
}

func fIntRet(x int) int {
	return x + 10
}

func fFloatRet(x float64) float64 {
	return x + 10
}

func fByteRet(x byte) byte {
	return x + 10
}

func fIntPtr(x *int) {
	*x += 10
}

func fFloatPtr(x *float64) {
	*x += 10
}

func fBytePtr(x *byte) {
	*x += 10
}
