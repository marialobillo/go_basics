package main

import "fmt"


type Salutation struct {
     Name string
     Greeting string
}

func (salutation Salutation) Rename(newName string) {
     salutation.Name = newName
}



func (salutations Salutation)Greet(salutation Salutation, times int){
     for _, s := range salutations {
     	 fmt.Println(salutations.Name)
     }
}

func main(){

     salutations := []Salutation{
     	   {"Bob", "Hello"},
	   {"Joe", "Hi"},
	   {"Mary", "What is up?"},
     }

     salutations.Greet(salutations, 2)

     salutations[0].Rename("Sherlock")
     
     
}
