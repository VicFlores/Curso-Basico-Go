package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println("My brand pc", myPc.brand)
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println("valor de A:", a)
	fmt.Println("Espacio en memoria de B:", b)
	fmt.Println("Valor en espacio de memoria:", *b)

	*b = 100

	fmt.Println("Cambio de valor", a)

	/* Aplicación */
	myPc := pc{ram: 16, disk: 200, brand: "HP"}

	fmt.Println("My Pc:", myPc)

	/* Function ping */
	myPc.ping()

	/* Function duplicate */
	fmt.Println("My pc:", myPc)

	myPc.duplicateRAM()

	fmt.Println("My pc:", myPc)
}
