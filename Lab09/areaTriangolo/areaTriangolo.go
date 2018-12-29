package main

import (
	"./triangolo"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))

	var (
		base, altezza int
		area          float64
	)
	t := triangolo.NewTriangolo()

	fmt.Println("Le 10 coppie di valori casuali generati per base ed altezza sono:")
	for i := 0; i < 10; i++ {
		b := rand.Intn(100) + 1
		h := rand.Intn(100) + 1
		fmt.Printf("%4d - %4d\n",b,h)
		triangolo.SetBase(t, b)
		triangolo.SetAltezza(t, h)
		if triangolo.Area(t) > area {
			area, base, altezza = triangolo.Area(t), b, h
		}
	}

	fmt.Println("\nIl triangolo con area maggiore è il seguente.\n")

	triangolo.SetBase(t, base)
	triangolo.SetAltezza(t, altezza)
	fmt.Println(triangolo.String(t))

}
