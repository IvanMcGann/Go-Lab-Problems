
//Adapted from : https://play.golang.org/p/yW7sAyJpPe
//				 http://golang-jp.org/pkg/math/big/	
// https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion			

//implements main method
package main

//import the format and math/big package
//math/big is used for big integers
import (
	"fmt"
	"math/big" 
)



func factorial(n int64) *big.Int { //get factorial of a number. can handle big numbers using int64
	if n < 0 {
		return big.NewInt(1) 
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))

}

//main method function
func main() {
	fmt.Println(factorial(100)) //print out the factorial of 100
	sum := big.NewInt(0) //set up sum as a big int 
	n := new (big.Int)//to store the factorial
	fact := (factorial(100)) 
	for fact.BitLen() > 0 { //loop through the bits until the end 
		_, n := fact.DivMod(fact, new(big.Int).SetUint64(uint64(10)), n) //DivMod implements Euclidean division and modulus as per reference at the top. 
		sum = sum.Add(sum,n) //add the numers for answer
	} 
	fmt.Println(sum)// print answer
}