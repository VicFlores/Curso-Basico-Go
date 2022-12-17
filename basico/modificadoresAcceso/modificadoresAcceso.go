package main

import (
	pkCar "curso-basico-go/basico/modificadoresAcceso/car"
	"fmt"
)

func main() {

	/* Public struct*/
	var myCarPublic pkCar.CarPublic

	myCarPublic.Brand = "Ferrari"
	myCarPublic.Year = 2009

	fmt.Println("my public car", myCarPublic)

	/* Public function */
	pkCar.PrintMessage1("Hola Vic")

}
