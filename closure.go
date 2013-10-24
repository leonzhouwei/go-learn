package main 

import "fmt"

func main() {
	f := func() {
		fmt.Println("closure")
	}
	f()
	func() {
		fmt.Println("lalala")
	}()
}