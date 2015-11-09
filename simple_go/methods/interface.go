package main

import "fmt"


type Shaper interface {
   Area() int
}

type Rectangle struct {
   length, width int
}


func (r Rectangle) Area() int {
   return r.length * r.width
}

func main() {
   r := Rectangle{length:5, width:3}
   fmt.Println("Rectangle r details are: ", r)  
   fmt.Println("Rectangle r's area is: ", r.Area())  

   s := Shaper(r)
   fmt.Println("Area of the Shape r is: ", s.Area())  
   
}