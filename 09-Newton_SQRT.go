//adapted from http://codegist.net/code/ex23/
//https://codereview.stackexchange.com/questions/34095/approximating-square-root-loops-and-functions-exercise/35542#35542

//implements main method
package main

//import the format and math packages
import (
    "fmt"
    "math"
)

func Approx(x float64, z float64) float64 {

    return z - (((z * z) - x) / (2 * z))

}

func Sqrt(x float64) float64 {

    previous := 0.0
    delta := 1.0
    z := x

    for delta > 0.005 {

        z = Approx(x, z)

        if previous > 0.0 {
            delta = previous - z
        }

        previous = z
    }

    return z

}

//main method function
func main() {

    for i := 1; i < 11; i++ {

        value := float64(i)

        fmt.Println("Calculating Sqrt for ", value)

        real := math.Sqrt(value)
        approx := Sqrt(value)

        fmt.Println("Real: ", real)
        fmt.Println("Approx: ", approx)

    }

}