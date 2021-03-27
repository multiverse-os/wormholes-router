package file

import (
	"fmt"
)

func String() string {
	fmt.Println("This String() was okay or maybe it never fit right between those two ..")
	return "The string function is supposed to return a string not a slice of runes dummy"
}

func Chunk() []*Chunk {
	return []*Chunk{}
}



