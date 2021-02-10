package main

import (
	"fmt"
)

func main(){
	// map -> Go : dict -> Python
	// use `make` function to create a map
	races := make(map[string]int)

	// assing key-value pairs (but these are string-value pairs)
	races["marathon"] = 26
	races["half"] = 13
	races["10K"] = 6
	races["5K"] = 3
	
	// print the whole map
	fmt.Println(races)
	// print value for a sepcific key
	fmt.Println(races["half"])

	// delete a key-value pair
	delete(races, "10K")
	fmt.Println(races)	
}