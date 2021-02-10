package main

import (
	"fmt"
)

func main() {
	// variable instantiation -> type
	var x int

	// assignment
	x = 5
	
	// print
	fmt.Println(x)

	var y int
	y = 9
	fmt.Println(y)

	var sum int
	sum = x+y
	fmt.Println(sum)

	// can be rewritten by using walrus operator -> := 
	// avoid `var` and `int` declaration, inferred
	a := 4
	b := 18
	total := a + b
	fmt.Println(total)
	
}