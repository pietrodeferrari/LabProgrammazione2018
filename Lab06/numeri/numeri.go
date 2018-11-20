package main

import "fmt"


// Il codice è diviso in diverse funzioni, ognuna con un compito ben specifico
// la funzione main stampa il menu e richiama la funzione desiderata
// (NB: si possono fare implementazioni molto più efficienti rispetto a quelle proposte in questo codice,
// ma in questo esercizio focalizzatevi sull'uso e riuso delle funzioni invece che sulla loro efficienza)
func main() {
	var scelta int
	var soglia uint // soglia unsigned int perché l'esercizio richiedeva una soglia non negativa

	fmt.Println("Inserisci una soglia:")
	fmt.Scan(&soglia)

	// Stampo un menu per la scelta del calcolo
	fmt.Print("Scegli una funzione:\n" +
		"1) numeri perfetti\n" +
		"2) numeri amichevoli\n" +
		"3) primi gemelli\n" +
		": ")
	fmt.Scan(&scelta)

	switch scelta {
	case 1:
		perfetti(soglia)
	case 2:
		amichevoli(soglia)
	case 3:
		gemelli(soglia)
	default:
		fmt.Println("Scelta errata!")
	}
}

// La funzione perfetti calcola e stampa i numeri perfetti inferiori ad una soglia passata come argomento
func perfetti(soglia uint) {
	fmt.Println("Numeri perfetti inferiori a", soglia, ":")

	// Per ogni numero i compreso tra 1 e la soglia si controlla se è perfetto, e in caso affermativo lo si stampa
	for i := uint(1); i < soglia; i++ { // il cast uint(1) serve per poter fare un confronto tra uint ed evitare errori di compilazione
		if èPerfetto(i) {
			fmt.Println(i)
		}
	}
}

// La funzione amichevoli calcola e stampa le coppie di numeri amichevoli sotto una soglia
func amichevoli(soglia uint) {
	fmt.Println("Numeri amichevoli inferiori a", soglia, ":")
	// Per confrontare coppie di numeri posso usare due cicli for
	// e confrontare se i e j sono amichevoli
	for i := uint(1); i < soglia; i++ {
		// j parte a i+1 per evitare il confronto tra numeri uguali
		// e doppi confronti (ad esempio fare uno solo tra i confronti i=1 e j=2, e i=2 e j=1)
		for j := i + 1; j < soglia; j++ {
			if sonoAmichevoli(i, j) {
				fmt.Print("(", i, ",", j, ")\n")
			}
		}
	}
}

// La funzione gemelli calcola e stampa le coppie di numeri gemelli sotto una soglia
func gemelli(soglia uint) {
	fmt.Println("Numeri gemelli inferiori a", soglia, ":")
	// Come per la funzione amichevoli, uso due for per scorrere tutte le coppie di numeri,
	// anche se, in questo caso, il secondo for avrebbe potuto essere evitato, dato che affinché due numeri siano gemelli
	// deve essere vero j==i+2
	for i := uint(1); i < soglia; i++ {
		for j := i + 2; j < soglia; j++ {
			if sonoGemelli(i, j) {
				fmt.Print("(", i, ",", j, ")\n")
			}
		}
	}
}

// La funzione restituisce true se il numero passato come argomento è perfetto.
// Un numero è perfetto se la somma dei divisori propri è uguale al numero stesso
// Viene utilizzata la funzione sommaDivisori() per calcolare la somma dei divisori,
// in questo modo il codice risulta più leggibile e semplice da capire
// NB: in go, a differenza di altri linguaggi, si possono usare caratteri non ASCII per i nomi di funzioni e variabili
func èPerfetto(numero uint) bool {
	return sommaDivisori(numero) == numero
}

// La funzione restituisce true se due numeri sono amichevoli.
// Due numeri sono amichevoli se la somma dei divisori del primo è uguale al secondo e viceversa.
// Usando nuovamente la funzione sommaDivisiori(), il corpo della funzione risulta più comprensibile
// e autoesplicativo
func sonoAmichevoli(n1, n2 uint) bool {
	return sommaDivisori(n1) == n2 && sommaDivisori(n2) == n1
}

// La funzione restituisce true se due numei sono gemelli.
// Due numeri sono gemelli se sono entrambi primi e la differenza tra i due è 2.
// Dato che non sappiamo quale sia il maggiore tra i due numeri, possiamo usare
// l'operatore logico or (||), oppure il metodo Abs del package math
// per il calcolo del valore assoluto della differenza tra n1 e n2
func sonoGemelli(n1, n2 uint) bool {
	return èPrimo(n1) && èPrimo(n2) && (n1 == n2+2 || n2 == n1+2)
}

// La funzione restituisce true se il numero è primo.
// In questo caso ho riutilizzato la funzione sommaDivisori() (ogni numero primo ha solo 1 solo divisore proprio
// che è 1, e quindi un numero è primo se la somma dei divisori propri è 1).
func èPrimo(numero uint) bool {
	return sommaDivisori(numero) == 1
}

// La funzione restituisce la somma dei divisori propri di un numero.
func sommaDivisori(numero uint) (somma uint) {
	for i := uint(1); i <= numero/2; i++ {
		if numero%i == 0 {
			somma += i
		}
	}
	return
}
