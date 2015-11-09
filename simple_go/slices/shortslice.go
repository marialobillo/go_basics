package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}


func main(){

    // s := []int {1, 10 , 500 , 25}
     

     fmt.Println("Hello")

     slice := []Salutation{
     	   {"Bob", "Hello"},
	   {"Joe", "Hi"},
	   {"Mary", "What is up?"},
     }

     slice = slice[0:2]
     slice = slice[1:]
     slice = slice[0:len(slice)]

     slice = append(slice, Salutation{"Harry", "Hi!"})

     fmt.Println(slice)
}
