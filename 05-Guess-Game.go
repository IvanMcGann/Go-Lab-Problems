package main
import (
    "fmt"
    "math/rand"
    "time"
)

func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 10)
    tries := 0
    var guess int

    fmt.Println("Welcome to Guess My Number Game!")
    for guess != myrand {
        fmt.Println("Take a guess...")
        fmt.Scanf("%v", &guess)
        tries++
        if guess > myrand {
            fmt.Println("Too high")
        } else if guess < myrand {
            fmt.Println("Too low")
        } else {
            fmt.Printf("Good job! You guessed it in %v tries", tries)
            break
        }
    }
}