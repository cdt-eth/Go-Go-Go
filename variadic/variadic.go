package main

import "fmt"

// take undefined num of ints as args with `...`
func sum(nums ...int){
	fmt.Print(nums, " ")
	total := 0
	
	// use blank identifier `_` bc we don't need indecies
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1,7,2,11,9)
	sum(1,3,6,8,4,3,5,6,8,7,6,54,4,444,2,4,2,3)
	sum(1,1)
}