// goto
package main

import (
	"fmt"
)

func main() {
	myFunc()
}

func myFunc() {
	for i := 0; ; i++ {
		fmt.Println(i)
		if i > 5 {
			goto HERE
		}
	}
	HERE:
}
