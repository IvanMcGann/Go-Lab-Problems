//adapted from https://gist.github.com/sergiotapia/7884903 

//implements main method
package main

//import the format package
import "fmt"

//main method function
//initially i is set to 1 to begin the for loop, if i is a modulus of 3 AND 5 eg 15 Fizzbuzz will be displayed to the screen
//However if i is a modulus of 3 e.g. 9 only Fizz will be displayed to the console
//Finally if i is only a modulus of 5 Buzz will be displayed to the user

func main() {
    i := 1
    for i <= 100 {
        if (i % 3 == 0 && i % 5 == 0) {
            fmt.Println("Fizzbuzz")
        } else if (i % 3 == 0) {
            fmt.Println("Fizz")
        } else if (i % 5 == 0) {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
        i = i + 1
    }
}