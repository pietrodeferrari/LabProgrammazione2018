package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
	"math/rand"
	"time"
)

const (
	NNUM     int = 10
	INTERVAL int = 20000
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var sequenze [][]float64

	/* S0 */
	var sl0 []float64
	fmt.Println("Inserisci una sequenza di numeri (uno per riga); premi due volte invio per terminare.")
	for scanner.Scan() {
		if number := scanner.Text(); number != "" {
			if value, err := strconv.ParseFloat(number,64); err == nil {
				sl0 = append(sl0, value)
			} else {
				fmt.Println("NaN... il valore non verrà considerato!")
			}
		} else {
			break
		}
	}
	sequenze = append(sequenze, sl0)
	n := len(sl0)

	/* S1 */
	sl1 := make([]float64, 2*(n-1))
	copy(sl1, sl0[1:])
	copy(sl1[n-1:], sl0[1:])
	for i:=0; i<n-1; i++ {
		sl1[i] *= math.Pow(2,float64(i+1))
	}
	for i:=0; i<n-1; i++ {
		sl1[n-1+i] *= math.Pow(2,float64(-(i+1)))
	}
	sequenze = append(sequenze, sl1)

	/* S2 */
	oneThird := n/3
	twoThird := int(2*(float64(n)/3))
	sl2 := make([]float64, oneThird + (n - twoThird) )
	copy(sl2[:oneThird], sl0[:oneThird])
	copy(sl2[oneThird:], sl0[twoThird:])
	sequenze = append(sequenze, sl2)

	/* S3 */
	sl3 := make([]float64, oneThird + NNUM + (n - twoThird))
	copy(sl3[:oneThird], sl0[:oneThird])
	copy(sl3[oneThird + NNUM:], sl0[twoThird:])

	rand.Seed(int64(time.Now().Nanosecond()))
	a := sl3[oneThird:oneThird+NNUM]
	for i := range a {
		a[i] = float64(rand.Intn(INTERVAL) - INTERVAL/2)
		fmt.Println(a[i])
	}

	sequenze = append(sequenze, sl3)

	for i, s := range sequenze {
		esaminaSequenza("S"+strconv.Itoa(i), s)
	}

}


func esaminaSequenza(nomeSequenza string, sl []float64) {

	fmt.Print("Sequenza ", nomeSequenza, ":\n")
	fmt.Println(sl)

	min, max, sum := sl[0], sl[0], sl[0]
	for i:=1; i<len(sl); i++ {
		if min > sl[i] {
			min = sl[i]
		}
		if max < sl[i] {
			max = sl[i]
		}
		sum += sl[i]
	}

	fmt.Print("Minimo in ", nomeSequenza, ": ", min, "\n")
	fmt.Print("Massimo in ", nomeSequenza, ": ", max, "\n")
	fmt.Print("Valore medio in ", nomeSequenza, ": ", sum/float64(len(sl)), "\n")
	fmt.Println()
}