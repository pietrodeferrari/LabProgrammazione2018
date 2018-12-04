package main

import "fmt"

func main() {

	m := make(map[string]int)

	for _, s := range []string{"questo", "è", "un", "test"} {
		m[s] = len([]rune(s))
	}

	for k, v := range m {
		fmt.Println(k, "->", v)
	}

}
