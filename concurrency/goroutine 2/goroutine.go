package main

import (
	"fmt"
	"time"
) 


func main(){
	// this will run forever so `count(fish)` will never run
	count("sheep")
	count("fish")
}

func count(thing string){
	// infinite loop
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}