package main

import "fmt"

func main() {

	var message string = "Hello Go World!"
	var greeting *string = &message

	fmt.Println(message, *greeting)
}
