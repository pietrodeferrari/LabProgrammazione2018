package triangolo

import (
	"fmt"
	)

type triangolo struct {
	base, altezza int
}

// Crea un nuovo triangolo restituendone il puntatore.
// Base e altezza sono inizializzati al valore di default 0.
func NewTriangolo() *triangolo {
	return new(triangolo)
}

// Calcola l'area di un triangolo (base x altezza) e restituisce il valore.
func Area(t *triangolo) float64 {
	return float64(t.base * t.altezza) / 2
}

// Imposta l'altezza del triangolo.
// Il primo argomento è il puntatore al triangolo.
// Il secondo argomento è il valore da assegnare all'altezza.
func SetAltezza(t *triangolo, h int) {
	t.altezza = h
}

// Imposta la base del triangolo.
// Il primo argomento è il puntatore al triangolo.
// Il secondo argomento è il valore da assegnare alla base.
func SetBase(t *triangolo, b int) {
	t.base = b
}

// Trasforma il triangolo passato come argomento in un testo.
// Il testo restituisto dalla funzione riposta i valori di base e altezza e l'area.
func String(t *triangolo) string {
	s := fmt.Sprintf("Triangolo con base e altezza rispettivamente pari a %d e %d.\nL'area del triangolo è pari a %.3f", t.base, t.altezza, Area(t));
	return s  
}

