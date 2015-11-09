package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}


func main(){
     

     fmt.Println("Hello")

     slice := []Salutation{
     	   {"Bob", "Hello"},
	   {"Joe", "Hi"},
	   {"Mary", "What is up?"},
     }

     slice = append(slice[:1], slice[2:]...)

     fmt.Println(slice)
}
