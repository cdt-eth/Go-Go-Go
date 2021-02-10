package main

import(
	"fmt"
)

func main(){
	// print 0-5
	for i := 0; i < 6; i++ {
		fmt.Print(i)
	}
	
	// newline
	fmt.Println("\n")

	// create while loop with for loop-like syntax by removing params
	// prints 1-4
	j := 1
	for j < 5 {
		fmt.Print(j)
		j++
	}
	
	fmt.Println("\n")
	
	// iterate over slice using `range`
	s := []string{"apple", "banana", "cherry"}
	for index, fruit := range s {
		fmt.Println("index", index, "fruit", fruit)
	}

	fmt.Println("\n")

	// same thing with a map
	m := make(map[string]string)
	m["p"] = "python"
	m["g"] = "go"
	m["s"] = "swift"
	m["c"] = "css"

	for abbrv,lang := range m {
		fmt.Println("abbrv", abbrv, "language", lang)
	}


}
