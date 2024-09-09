package main

import "fmt"

// Go supports only for loop

func main(){
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j:= 0; j < 3; j++{
		fmt.Println(j)
	}
	
	for {
		fmt.Println("loop")
		break
	}

}