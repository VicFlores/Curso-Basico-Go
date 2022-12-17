package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := 1
	value2 := 2

	/* If tradicional */

	if value1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("Es", value1)
	}

	/* If with and */
	if value1 == 1 && value2 == 2 {
		fmt.Println("Es verdad AND")
	}

	/* If with or */
	if value1 == 2 || value2 == 2 {
		fmt.Println("Es verdad OR")
	}

	/* Convertir texto a numero */
	value, err := strconv.Atoi("777")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Texto a numero", value)

}
