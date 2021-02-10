package main

import (
	"fmt"
)

// create struct type
type anime struct {
	show string
	episodes int
}

func main() {
	a := anime{show: "One Piece", episodes:931}
	fmt.Println(a)
	fmt.Println(a.show)
	fmt.Println(a.episodes)
}