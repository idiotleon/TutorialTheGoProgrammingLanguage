package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")

	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	// Variadic Function Example
	message = greeting[1] + " " + name
	alternate = "HEY!" + " " + name
	return
}

func main() {
	var salutation = Salutation{"Leon", "Hello"}
	Greet(salutation)
}
