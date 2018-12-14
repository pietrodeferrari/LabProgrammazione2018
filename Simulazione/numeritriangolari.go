package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if t, err := strconv.Atoi(os.Args[1]); err == nil {

		triangoli, quadrati := make([]int, t), make([]int, t)
		for i:=0; i<t; i++ {
			if i == 0 {
				triangoli[i], quadrati[i] = 0, 0
			} else {
				triangoli[i], quadrati[i] = triangoli[i-1]+i, i*i
			}
		}

		fmt.Println(triangoli)
		fmt.Println(quadrati)

		for i:=0; i<t; i++ {
			for j:=0; j<t; j++ {
				if quadrati[i] == triangoli[j] {
					fmt.Printf("%d triangolare e quadrato\n", quadrati[i])
				}
			}
		}

	}
}
