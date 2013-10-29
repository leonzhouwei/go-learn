package main

import "fmt"

func main() {
	a := add(0)
	fmt.Println(a)
	b := add(2)
	fmt.Println(b)
}

func add(x int) int {
	if x == 0 {
		return 5
	} 
	return x
}