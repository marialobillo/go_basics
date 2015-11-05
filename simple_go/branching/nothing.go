package main

import "fmt"

func GetPrefix(name string) (prefix string){
  switch {
    case name == "Bob":
      prefix = "Mr "
    case name == "Joe", name == "Amy", len(name) == 10:
      prefix = "Dr"
    case name == "Mary":
      prefix = "Mrs"
    default:
      prefix = "Dude "

  }
  return
}

func Greet(name string, isFormal bool){

  if prefix := GetPrefix(name); isFormal {
    fmt.Println(prefix + " " + name)
  }
}

func main(){

  Greet("123456790", true)
}
