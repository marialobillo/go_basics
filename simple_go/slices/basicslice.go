package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}


func main(){

     var s[]int
     s = make([]int, 3, 10 )
     s[0] = 1
     s[1] = 10
     s[2] = 500

     fmt.Println("Hello")

     slice := []Salutation{
     	   {"Bob", "Hello"},
	   {"Joe", "Hi"},
	   {"Mary", "What is up?"},
     }

     fmt.Println(slice)
}
