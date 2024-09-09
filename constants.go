package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean, and numeric values.
// const declares a constant value
const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 500000000
	
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

// output:

// constant
// 6e+11
// 600000000000
// -0.28470407323754404