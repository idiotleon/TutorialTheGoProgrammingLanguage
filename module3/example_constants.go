package main

import "fmt"

const (
	PI       = 3.14
	Language = "Golang"
)

const (
	// iota represents successive untyped integer constants
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E
	F
)

func main() {

	fmt.Println(PI)
	fmt.Println(Language)

	fmt.Println(A, B, C)

	fmt.Println(D, E, F)
}
