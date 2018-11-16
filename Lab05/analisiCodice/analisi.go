package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		var a, b int = 5, 7
		if a == b {
			fmt.Println("Sono uguali")
		} else {
			fmt.Println("Sono diversi")
		}
	}

	fmt.Println(strings.Repeat("=", 20))

	{
		var a int = 7
		var c bool
		b := 5
		c = a == b
		fmt.Println(c)
	}

	fmt.Println(strings.Repeat("=", 20))

	{
		var a, b, c int = 7, 5, 12
		fmt.Println(a + b*c)
	}

	fmt.Println(strings.Repeat("=", 20))

	{
		var zero float64 = 0
		res := zero / zero
		fmt.Println(res)
		fmt.Printf("res = %f, tipo di res = %[1]t\n", res)
	}

	fmt.Println(strings.Repeat("=", 20))

	{
		var numero float64 = 0
		fmt.Println(numero)
		numero = 1 / numero
		fmt.Println(numero)
		numero = 1 / numero
		fmt.Println(numero)
		numero = 1 / numero
		numero = numero - numero
		fmt.Println(numero)
	}
}
