package main

import (
	"fmt"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

func main(){

	var s []string

	fmt.Println("uninit",s, s == nil, len(s) == 0)

	s = make([]string,3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len of slice", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	// Slices can also be copy’d
	c := make([]string, len(s))
	copy(c,s)

	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high]
	l := s[2:5]
	fmt.Println("sl1", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2", l)

	// slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3", l)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g","h","i"}
	fmt.Println("dcl", t)

	// t2 := []string{"g","h","i"}

	// if slices.Equal(t,t2){
	// 	fmt.Println("t == t2")
	// }

	twoD := make([][]int,3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0 ; j < innerLen ; j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D:", twoD)

}