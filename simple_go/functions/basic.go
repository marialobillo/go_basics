package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation){
  fmt.Println(salutation.name);
  fmt.Println(salutation.greeting)
}

func main() {
	var s = Salutation{"Bob", "Hello"}

  Greet(s)
}
