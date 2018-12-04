package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var a []int

	soglia, _ := strconv.Atoi(os.Args[1])

	rand.Seed(int64(time.Now().Nanosecond()))

	for n := int(soglia)+1; n > int(soglia); {
		n = rand.Intn(100)
		a = append(a, n)
	}


	fmt.Println("Valori generati", a)
	fmt.Println("Valori sopra soglia:", a[:len(a)-1])

}
