//adapted from https://gist.github.com/pyk/10519339


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
      fmt.Println(v,"is greater than",n)
    } else {
      fmt.Println(v,"is less than",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is: ", smallest)
}