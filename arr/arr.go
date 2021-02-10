package main

import (
	"fmt"
)

func main(){

	// like C++, we can allot space ahead of time to an array
	var arr[5]int
	
	// assinging and printing is the same
	arr[4] = 5
	arr[2] = 9
	fmt.Println(arr)

	// or use walrus to assign everything together
	b := [5]int{1,5,3,21,999}

	fmt.Println(b)

	// can use a slice not instantiate without a fixed size
	c := []int{1,2,3,4,5,6,7}
	fmt.Println(c)
	
	// now append 8th element
	c = append(c,8)
	fmt.Println(c)
}