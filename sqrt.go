package main

import (
	"fmt"
)

func Sqrt(x, z float64) float64 {
	for i:=0; i < 10; i+=1 {
		z = z - (z*z-x)/(2*z)	
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 1))
	fmt.Println(Sqrt(2, 2))
	fmt.Println(Sqrt(2, 3))
	fmt.Println(Sqrt(2, 4))
	fmt.Println(Sqrt(2, 5))
	fmt.Println(Sqrt(2, 10))
}
