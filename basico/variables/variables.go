package main

import "fmt"

func main() {

	/* Declaración de constantes */
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	/* Variables enteras */
	var area int
	var altura int = 14
	base := 12

	fmt.Println("area:", area)
	fmt.Println("altura:", altura)
	fmt.Println("base:", base)

	/* Zero value */
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("zeroInt:", zeroInt)
	fmt.Println("zeroFloat:", zeroFloat)
	fmt.Println("zeroString:", zeroString)
	fmt.Println("zeroBool:", zeroBool)

	/* Area cuadrado */
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("areaCuadrado:", areaCuadrado)

}
