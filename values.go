package main

import "fmt"

func main(){
	// go has various value types including strings, integers, floats
	// booleans etc.

	fmt.Println("go" + "lang") //string

	fmt.Println("1+1 =", 1+1) //integer

	fmt.Println("0.3/0.3 =",0.3/0.3) //float

	fmt.Println("boolean true value =", true) //boolean true
	fmt.Println("boolean false value =", false) //boolean false

	//logical values

	fmt.Println(true && false) //and operator
	fmt.Println(true || false) //or operator
	fmt.Println(!true) //not operartor
}