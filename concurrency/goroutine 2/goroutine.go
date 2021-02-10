package main

import (
	"fmt"
	"time"
) 


func main(){
	// // this will run forever so `count(fish)` will never run
	// count("sheep")

	// created a goroutine so that now they can run side-by-side
	go count("sheep")
	count("fish")

	// if we ran both functions as goroutines we'd enter a deadlock because both functions are queued and waiting for another line of code that won't run
}

func count(thing string){
	// infinite loop
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}