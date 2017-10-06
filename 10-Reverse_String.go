//adapted from https://coderwall.com/p/fw6miw/reverse-text-in-golang
//https://golang.org/doc/effective_go.html#for


//implements main method
package main

//import the format package
import "fmt"

//main method function
func main() {
    s := "Hello, playground" //creates a slice(an array in go)  
    fmt.Println(s) //displays a slice of string to the user 
    fmt.Println(Reverse(s)) //reverses the slice using the Reverse method
}

//Reverse method function
func Reverse(s string) string {
    var reverse string //declares a string called reverse
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])//reverses the characters in the string by using the for method 
    }
    return reverse //returns the string
}

