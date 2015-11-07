package main

import "fmt"

func Greet(prefix int){

     for i:= 0; i < prefix; i++{
     	 if true {
	    fmt.Println( i )
	 }
     }
}


func main(){

     Greet(5)

     fmt.Println()
}