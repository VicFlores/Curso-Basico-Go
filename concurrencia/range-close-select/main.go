package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println("Length:", len(c))
	fmt.Println("Capacity:", cap(c))

	/* Cierra el canal */
	close(c)

	/* Iterar en los valores del channel */
	for message := range c {
		fmt.Println(message)
	}

	/* Select */
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Email 1", email1)
	go message("Email 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)

		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
