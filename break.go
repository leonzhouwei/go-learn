package main

import "fmt"

func main() {
	i, j, k := 0, 0, 0
	for ; k < 3; k++ {
		for ; j < 5; j++ {
			for ; i < 10; i++ {
				if i > 5 {
					break JLoop
				}
				fmt.Println(i)
			}
		}
	JLoop:
	}

	fmt.Println(i, j, k)
}
