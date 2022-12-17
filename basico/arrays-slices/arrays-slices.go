package main

import "fmt"

func main() {

	/* Array (inmutables) */
	var array [4]int

	array[0] = 1
	array[1] = 2
	array[1] = 2

	fmt.Println("Array", array)
	fmt.Println("(Array): Cantidad de elementos:", len(array))
	fmt.Println("(Array): Capacidad de elementos:", cap(array))

	/* Slices (mutables) */
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println("Slice", slice)
	fmt.Println("(Slice): Cantidad de elementos:", len(slice))
	fmt.Println("(Slice): Capacidad de elementos:", cap(slice))

	/* Métodos en slice */
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  /* hasta el index N */
	fmt.Println(slice[2:4]) /* entre A y B */
	fmt.Println(slice[4:])  /* desde N en adelante */

	/* Agregar elemento a un slice */
	slice = append(slice, 7)

	fmt.Println("Append element slice", slice)

	/* Agregar una lista */
	newSlice := []int{8, 9, 10}

	slice = append(slice, newSlice...)

	fmt.Println("Append list slice", slice)

}
