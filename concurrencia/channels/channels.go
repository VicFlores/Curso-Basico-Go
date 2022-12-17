package main

import "fmt"

func saySomething(text string, c chan<- string) {
	c <- text
}

func sayBye(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	fmt.Println("Hello Vic")

	go saySomething("I love pupusas", c)
	go sayBye("Bye", c)

	fmt.Println(<-c)
	fmt.Println(<-c)

}
