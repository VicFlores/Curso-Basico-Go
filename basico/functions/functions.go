package main

import "fmt"

func sayHello(message string) {
	fmt.Println(message)
}

func manyArguments(a, b int, c string) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func returnValue(a int) int {
	return a * 4
}

func doubleArguments(a int) (b, c int) {
	result1 := a + 5
	result2 := a * 2

	return result1, result2
}

func main() {

	sayHello("Hola Vic Flores")

	manyArguments(2, 4, "Wenas")

	value := returnValue(2)
	fmt.Println("Return value:", value)

	value1, value2 := doubleArguments(2)
	fmt.Println("Return value1:", value1)
	fmt.Println("Return value2:", value2)

	value1, _ = doubleArguments(2)
	fmt.Println("Return value1:", value1)
}
