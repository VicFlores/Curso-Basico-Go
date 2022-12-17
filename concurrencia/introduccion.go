package main

import (
	"fmt"
	"sync"
)

func saySomething(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello Vic")

	wg.Add(2)

	go saySomething("Hi Mr Escobar", &wg)

	go func(text string, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println(text)
	}("Hello Mr Flores", &wg)

	wg.Wait()

}
