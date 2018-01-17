package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

// example: function with one return
func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

func Greet(salutation Salutation) {

	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func main() {
	var s = Salutation{"Leon", "Hello"}
	Greet(s)
}
