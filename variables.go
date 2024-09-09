package main

import "fmt"

func main(){

	// In Go, variables are explicitly declared and used by the compiler to 
	// e.g. check type-correctness of function call

	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

// output:

// initial
// 1 2
// true
// 0
// apple