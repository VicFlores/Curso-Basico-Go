package main

import "fmt"

func main() {
	modulo := 4 % 2

	/* Con condición */
	switch modulo {

	case 0:
		fmt.Println("Es par")

	default:
		fmt.Println("No es par")
	}

	/* Sin condición */
	value := 250

	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")

	case value < 0:
		fmt.Println("Es menor a 0")

	default:
		fmt.Println("No condición")
	}

}
