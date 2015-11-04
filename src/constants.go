package main

import "fmt"

const (
	PI       = 3.14
	Language = "Go"
	A        = iota
	B        = iota
	C        = iota
)

func main() {

	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A, B, C)
}
