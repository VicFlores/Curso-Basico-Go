package main

import (
	"fmt"
)

func main() {

	/* For */
	for i := 0; i < 10; i++ {
		fmt.Println("index", i)
	}

	fmt.Printf("\n")

	/* For while */
	counter := 0

	for counter < 10 {
		fmt.Println("counter", counter)
		counter++
	}

	/* For forever */
	counterForever := 0

	for {
		fmt.Println("Counter Forever")
		counterForever++
	}

}
