package main

import "fmt"

func main() {

	message := "Hello Go World with Pointer!"

	var greeting *string = &message

	fmt.Println(message, greeting)
	fmt.Println(message, *greeting)

	// Pointer Assignment
	*greeting = "Hello Go World with Pointer 2!"
	fmt.Println(message, *greeting)
}
