package main

import (
	"fmt"
)

/*
Maps are Go’s built-in associative data type 
(sometimes called hashes or dicts in other languages)
*/

func main(){
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	// If the key doesn’t exist, the zero value of the value type is returned
	v2 := m["k3"]
	fmt.Println("v2", v2)

	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m,"k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin
	// clear(m)
	// fmt.Println("map:",m)

	/*
		The optional second return value when getting a value from a map indicates if the key was present in the map. 
		This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
		Here we didn’t need the value itself, so we ignored it with the blank identifier _
	*/

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map n:", n)

	// n2 := map[string]int{"foo": 1, "bar": 2}
    // if maps.Equal(n, n2) {
    //     fmt.Println("n == n2")
    // }
}