package main

import "fmt"

func main() {

	var message string = "Hello Go World!"
	var greeting *string = &message

	//variable and pointer is = "hi"
	*greeting = "hi"

	fmt.Println(message, *greeting)
}
