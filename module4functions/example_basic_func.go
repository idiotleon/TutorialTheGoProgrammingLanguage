package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	fmt.Println(salutation.name, salutation.greeting)
}

func main() {
	var s = Salutation{"Leon", "Hello!"}
	Greet(s)
}
