package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	myCar := car{brand: "Ford", year: 2020}

	fmt.Println("My car:", myCar)

	/* Otra forma */
	var otherCar car

	otherCar.brand = "Hyundai"
	otherCar.year = 2013

	fmt.Println("Other car", otherCar)
}
