package main

import "fmt"

func Greet(salutation string, isFormal bool){

  fmt.Println(salutation)
  if isFormal {
    fmt.Println("Is Formal")
  }
}

func main(){

  Greet("Hello World", true)
}
