package main

import "fmt"


type Salutation struct {
	name     string
	greeting string
}

func TypeSwitchTest(x interface{}) {
  switch x.(type) {
    case int:
      fmt.Println("int")
    case string:
      fmt.Println("string")
    case Salutation:
      fmt.Println("salutation")
    default:
      fmt.Println("unknown")

  }
  return
}


func main(){

  TypeSwitchTest("hello")
}
