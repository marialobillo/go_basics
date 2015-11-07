package main

import "fmt"


func Greet(prefix int){

     i := 0
     for i < prefix {

     	 fmt.Println("Hello")
	 i++
     }

}



func main(){

     Greet(6)
}