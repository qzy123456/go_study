package main
import "fmt"

func main(){
var a  int = 20
var b *int
 b = &a
 fmt.Printf("a",&a)
 fmt.Printf("b",b)
  fmt.Printf("b",*b)
 

}
