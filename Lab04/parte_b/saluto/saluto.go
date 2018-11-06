package main

import (
	"fmt"
	"time"
)

func main() {

	var nome string

	fmt.Println("Come ti chiami?")
	fmt.Scan(&nome)

	fmt.Println("Ciao", nome, "sai che oggi è il", time.Now())
}
