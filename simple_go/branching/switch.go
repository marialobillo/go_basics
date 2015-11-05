package main

import "fmt"

func GetPrefix(name string) (prefix string){
  switch name {
    case "Bob": prefix = "Mr "
    case "Joe", "Amy": prefix = "Dr"
    case "Mary": prefix = "Mrs"
    default: prefix = "Dude "

  }
  return
}

func Greet(name string, isFormal bool){

  if prefix := GetPrefix(name); isFormal {
    fmt.Println(prefix + " " + name)
  }
}

func main(){

  Greet("Mary", true)
}
