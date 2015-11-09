package main

import "fmt"


type Salutation struc {
     Name string
     Greeting string
}

type Salutations []Salutation


func (salutations Salutations)Greet(salutation Salutations, times int){
     for _, s := range salutations {
     	 
     }
}

func main(){

     salutations := Salutation{
     	   {"Bob", "Hello"},
	   {"Joe", "Hi"},
	   {"Mary", "What is up?"},
     }

     salutations.Greet()
     
}
