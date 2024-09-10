package main

import (
	"fmt"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

func main(){

	var s []string

	fmt.Println("uninit",s, s == nil, len(s) == 0)

	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

}