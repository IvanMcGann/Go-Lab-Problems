
package main

import "fmt"

func main() {
  var n, smallest, biggest int
  x := []int{
    13,98,46,23,
    57,64,
  }

  for _,v:=range x {
    if v>n {
      fmt.Println(v,"is greater than",n)
      n = v
      biggest = n
    } else {
      fmt.Println(v,"is less than",n)
    }
  }

  fmt.Println("The biggest number is: ", biggest)
  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
    } else {
      fmt.Println(v,"<",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is: ", smallest)
}