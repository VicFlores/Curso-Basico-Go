package main

import "fmt"

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindrome")
	} else {
		fmt.Println("No es palindrome")
	}
}

func main() {
	slice := []string{"a", "b", "c"}

	for i, valor := range slice {
		fmt.Printf("index: %d, valor: %s\n", i, valor)

	}

	for i := range slice {
		fmt.Printf("index: %d\n", i)

	}

	for _, valor := range slice {
		fmt.Printf("valor: %s\n", valor)

	}

	isPalindrome("roma")
}
