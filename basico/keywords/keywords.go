package main

import "fmt"

func main() {

	/* Defer */
	defer fmt.Println("Hello")
	fmt.Println("Vic")

	/* Continue - break */
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Es 8, break!")
			break
		}
	}
}
