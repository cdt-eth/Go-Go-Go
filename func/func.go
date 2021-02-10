package main

import (
	"fmt"
	"errors" // built-in error functions
	"math" // built-in square root function
)


func main(){
	// print func run
	fmt.Println(sum(18,67))
	
	fmt.Print("\n")
	
	// assign result
	result := sum(19,999)
	fmt.Println(result)
	
	fmt.Print("\n")
	
	res,err := square(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	
}

// declare func
// important: type goes after variable
func sum(x int, y int)int{
	return x + y
}

// the (float64, error) denotes what the `square` function CAN return
// the error is a built-in type to
func square(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("ERROR:negative numbers")
	}
	return math.Sqrt(x), nil
}