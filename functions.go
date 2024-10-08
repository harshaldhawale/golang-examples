package main

import "fmt"

func plus(a int,b int)int{
	return a + b
}

func plusPlus(a int,b int,c int)int{
	return a + b + c
}

// multiple return values

func vals()(int,int){
	return 3,7
}


func main(){
	res := plus(1,2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1 + 2 + 3 =", res)

	// for multiple returns
	a,b := vals()
	fmt.Println("a :", a)
	fmt.Println("b :", b)

	_, c := vals()
	fmt.Println("c :", c)

}