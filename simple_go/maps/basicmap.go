package main

import "fmt"


func GetPrefix(name string) (prefix string){

     var prefixMap map[string]string
     prefixMap = make(map[string]string)

     prefixMap["Bob"] = "Mr "
     prefixMap["Joe"] = "Dr "
     prefixMap["Amy"] = "Dr "
     prefixMap["Mary"] = "Mrs "

     return prefixMap[name]
}

func main(){

     fmt.Println(GetPrefix("Mary") + "Mary")
}
