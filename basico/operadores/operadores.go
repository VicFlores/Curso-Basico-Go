package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	y := 50

	/* Suma */
	result := x + y

	fmt.Println("Suma:", result)

	/* Resta */
	result = y - x

	fmt.Println("Resta:", result)

	/* Multiplicación */
	result = x * y

	fmt.Println("Multiplicación:", result)

	/* Division */
	result = y / x

	fmt.Println("Division:", result)

	/* Modulo o Residuo */
	result = y % x

	fmt.Println("Modulo:", result)

	/* Incremental */
	x++

	fmt.Println("Incremental:", x)

	/* Decremental */
	x--

	fmt.Println("Decremental:", x)

	/* Area de rectángulo */
	base := 10
	altura := 2

	result = base * altura

	fmt.Println("Area de rectángulo:", result)

	/* Area de trapecio */
	a := 5
	b := 10
	altura = 4

	result = (a + b) / 2 * altura

	fmt.Println("Area de trapecio:", result)

	/* Area de circulo */
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)

	fmt.Println("Area de circulo:", areaCirculo)

}
