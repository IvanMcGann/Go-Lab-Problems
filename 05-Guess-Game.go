//adapted from http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html
//https://golang.org/pkg/math/rand/

//implements main method
package main

//import the format and math/rome and time packages
import (
    "fmt"
    "math/rand"
    "time"
)

//function for generating random number
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

//main method function
func main() {
    myrand := xrand(1, 10) //range for the random number to be between
    tries := 0 //counter for tries
    var guess int //declares a integer called guess

    //displays headings to console
    fmt.Println("Welcome to the Guess the Number Game!")
    fmt.Println("Make a choice between 1 and 10")

    //
    for guess != myrand {
        fmt.Println("Take a guess...") //asks user to enter guess
        fmt.Scanf("%v \n", &guess) //takes in the guess
        tries++ //counts attempts
        if guess > myrand {
            fmt.Println("Too high")//displays user if the guess is higher than the random correct number
        } else if guess < myrand {
            fmt.Println("Too low")//displays user if the guess is less than the random correct number
        } else {
            fmt.Printf("Good job! You guessed it in %v tries", tries)//tells user when they guessed correctly 
            break //terminates the for loop and ends the program 
        }
    }
}