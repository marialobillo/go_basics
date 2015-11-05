package main

import "fmt"

func Greet(salutation string, isFormal bool){

  fmt.Println(salutation)
  if prefix := "Mr "; isFormal {
    fmt.Println(prefix + "Is Formal")
  }
}

func main(){

  Greet("Hello World", true)
}
