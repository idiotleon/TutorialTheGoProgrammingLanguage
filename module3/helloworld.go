package main

import "fmt"

func main() {

	// declare a variable
	var message string
	// assignment to a variable
	message = "Hello Go World!"

	fmt.Printf(message)

	// declare a variable with initialization
	var messageWithInit string = "Hello Go World with Initialization!"
	fmt.Printf(messageWithInit)

	// declaration and initialization
	message3 := "Hello Go World!"
	fmt.Printf(message3)

	// declare of multiple variables
	// with initialization
	var a, b, c int = 1, 2, 3
	// Println accepts multiple variables
	fmt.Println(message, a, b, c)

}
