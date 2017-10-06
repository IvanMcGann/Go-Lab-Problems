// adapted from http://www.golangpro.com/2016/01/check-string-palindrome-go.html

//implements main method
package main

//import the format and strings packages
import (
 "fmt"
 "strings"
)

//main method function
func main() {

 var ip string //creates string called ip
 fmt.Println("Enter string:") //requests user to enter a string to test
 fmt.Scanf("%s\n", &ip) //reads in the user string through a scanf function
 ip = strings.ToLower(ip) //sets the input string to lowercase making it easier to read
 fmt.Println(isP(ip)) //reads string using method below
}

//Function to test if the string entered in console is a Palindrome
func isP(s string) string {
 mid := len(s) / 2 //calculates the midpoint of the slice
 last := len(s) - 1 //calculates the end of the slice like an array (arrays begin at 0 not 1)
 for i := 0; i < mid; i++ {//if a letter greater <blank> 
  if s[i] != s[last-i] { //compares the slices
   return "NO. It's not a Palimdrome." //displays if false
  }
 }
 return "YES! You've entered a Palindrome" //displays if true

}