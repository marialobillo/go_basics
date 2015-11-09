package main

import "fmt"


func GetPrefix(name string) (prefix string){

     prefixMap := map[string]string {
     	      "Bob": "Mr ",
	      "Joe": "Dr ",
	      "Amy": "Dr ",
	      "Mary": "Mrs ",
     }

     //Update
     prefixMap["Joe"] = "Jr "

     //Delete
     delete(prefixMap, "Amy")
     
     
     return prefixMap[name]

     return
}

func main(){

     fmt.Println(GetPrefix("Mary") + "Mary")
}

