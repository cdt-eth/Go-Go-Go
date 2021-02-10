package main

import (
	"fmt"
)

func main(){
	i := 4
	
	fmt.Println(i)

	for i < 15 {
		// send pointer with `&`
		inc(&i)
		fmt.Println(i)
	}
}

// pass variable by reference using * symbol (pointer)
// it's not getting a copy, it's getting the value at that memory reference 
func inc(x *int){
	// deallocate pointer then increment
	*x++
}