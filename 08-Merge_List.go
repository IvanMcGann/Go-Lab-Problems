//adapted from https://play.golang.org/p/Ma2GXvj3XP
//https://blog.golang.org/slices
//http://golanging.blogspot.ie/2013/04/go-slices-and-merge-sort.html


//implements main method
package main

//import the format package
import (
	"fmt"
)

//function for merging the list
func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...) //append retrieves element of a slice
		}
		if len(r) == 0 {
			return append(ret, l...) //append retrieves element of a slice
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0]) //append retrieves element of a slice
			l = l[1:]
		} else {
			ret = append(ret, r[0]) //append retrieves element of a slice
			r = r[1:]
		}
	}
	return ret
}

//function for sorting the list
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

//main method function
func main() {
	s := []int{9, 4, 3, 6, 1, 2, 10,82, 95, 44,63, 5, 7, 8} //slice called s containing a unsorted list to merge and sort
	fmt.Printf("%v\n%v\n", s, MergeSort(s)) //prints the unsorted and sorted list to the console
}