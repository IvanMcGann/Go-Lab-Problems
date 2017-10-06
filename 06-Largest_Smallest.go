//adapted from https://gist.github.com/pyk/10519339

//implements main method
package main

//import the format package
import "fmt"

func main() {
  var n, smallest, biggest int //created vriables for n which is the number, the smallest and biggest numbers
  x := []int{
    13,98,46,23,
    57,64,
  } //creates int slice 

  for _,v:=range x {
    if v>n {
      fmt.Println(v,"is greater than",n)// compares the current n starting with the first value v
      n = v                             // n is given the value of v if it is greater than the previous v 
      biggest = n                       // it is concurrently compares to the next value taken in until the slice is read
    } else {
      fmt.Println(v,"is less than",n) // if it is not greater than a println will display that
    }
  }

  fmt.Println("The biggest number is: ", biggest) // biggest value is displayed to console 

  //for loop to find smallest, runs the same code as above but displays the smallest value
  for _,v:=range x {
    if v>n {
      fmt.Println(v,"is greater than",n)
    } else {
      fmt.Println(v,"is less than",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is: ", smallest) // smallest value is displayed to the console
}