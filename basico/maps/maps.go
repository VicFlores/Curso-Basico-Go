package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Jack"] = 19
	m["Jill"] = 24

	fmt.Println(m)

	/* Recorrer map */
	for i, valor := range m {
		fmt.Printf("index: %s, valor: %d\n", i, valor)
	}

	/* Encontrar un valor */
	value, ok := m["Jose"]

	fmt.Println("value:", value)
	fmt.Println("ok:", ok)
}
