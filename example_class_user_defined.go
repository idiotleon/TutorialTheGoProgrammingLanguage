package main

import "fmt"

type SalutationSimple string

type SalutationStructed struct {
	name     string
	greeting string
}

func main() {

	var message SalutationSimple = "Hello Go World from user defined type! [simple]"

	fmt.Print(message)

	var messageStructed = SalutationStructed{"leon", "Hello Go World from user defined type! [structed]"}

	fmt.Println(messageStructed.name, messageStructed.greeting)
}
