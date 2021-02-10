package main

import (
	"fmt"
)

// create struct type
type anime struct {
	title string
	epCount int
}

// func newAnime(name string) *anime {
// 	a := anime{title: title}
// 	a.epCount = 20
// 	return &a
// }


func main() {
	a := anime{title: "One Piece", epCount:931}
	fmt.Println(a)
	fmt.Println(a.title)
	fmt.Println(a.epCount)

	// fmt.Println(newAnime{title:"Attack On Titan", epCount:75})
	// fmt.Println(newAnime{"Attack On Titan"})
}

