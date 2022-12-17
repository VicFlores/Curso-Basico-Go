package main

import "fmt"

func main() {

	/* Variables */
	helloMessage := "Hello"
	nameMessage := "Vic"

	/* Println */
	fmt.Println(helloMessage, nameMessage)

	/* Printf */
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) /* %s = string, %d = int */
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) /* %v = any */

	/* Sprintf */
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)

	fmt.Println(message)

	/* Tipo de dato de una variable */
	var1 := "Hello"
	var2 := 69
	var3 := 3.1416

	fmt.Printf("El tipo de la variable(var1): %T\n", var1)
	fmt.Printf("El tipo de la variable(var2): %T\n", var2)
	fmt.Printf("El tipo de la variable(var3): %T\n", var3)

}
