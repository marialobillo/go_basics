package main

import "fmt"

func Greet(salutation string, isFormal bool){

  fmt.Println(salutation)
  if prefix := "Mr "; isFormal {
    fmt.Println(prefix + "Is Formal")
  }else{
    fmt.Println(prefix + "Is not Formal")
  }
}

func main(){

  Greet("Hello World", true)
}
