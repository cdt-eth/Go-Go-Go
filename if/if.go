package main

import (
	"fmt"
)

func main(){
	i := 22
	// i := 16
	// i := 8
	if i > 20 {
		fmt.Println("over 20!")
	} else if i > 15  {
		fmt.Println("between 15 & 20")
	} else {
		fmt.Println(" :( ")
	}
}