package main

import "fmt"

type SalutationSimple string

func main() {

	var message SalutationSimple = "Hello Go World from user defined type [simple]"

	fmt.Print(message)
}
