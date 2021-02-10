package main

import(
	"fmt"
)

func main(){
	// print 0-5
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
	
	// newline
	fmt.Println("\n")

	// create while loop with for loop-like syntax by removing params
	// prints 1-4
	j := 1
	for j < 5 {
		fmt.Println(j)
		j++
	}
}
