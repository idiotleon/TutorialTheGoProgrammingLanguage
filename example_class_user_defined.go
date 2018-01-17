package main

import "fmt"

type SalutationSimple string

/*
*	With a Captical S, it means that this is publicly visible
*	outside of this package
*	So if we were to put this in its own package and import that package,
*	anything that has a capital is going to be visible and anything,
*	which is NOT the case in terms of lower case initials.
*	So there is no need for private and protected in public
 */
type SalutationStructed struct {
	name     string
	greeting string
}

func main() {

	var message SalutationSimple = "Hello Go World from user defined type! [simple]"

	fmt.Println(message)

	var messageStructed = SalutationStructed{"leon", "Hello Go World from user defined type! [structed]"}

	fmt.Println(messageStructed.name, messageStructed.greeting)

	// another way to assign values
	var messageStructed2 = SalutationStructed{}
	messageStructed2.name = "Leon"
	messageStructed2.greeting = "Hello Go World from user structed type 2! [structed 2]"

	fmt.Println(messageStructed2.name, messageStructed2.greeting)
}
