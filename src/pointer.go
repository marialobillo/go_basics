package main

import "fmt"

func main() {

	message := "Hello Go World!"

	var greeting *string = &message

	fmt.Println(message, greeting)
}
