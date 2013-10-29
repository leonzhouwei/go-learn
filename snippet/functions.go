package main

import (
	"fmt"
	"errors"
)

func add(x int, y int) int {
	return x + y
}

func divide(numerator int, denominator int) (ret int, err error) {
	if denominator == 0 {
		ret = -1
		err = errors.New("denominator shall not be zero")
		return
	}
	return numerator / denominator, nil
}

func unfixedsized(args ...int) {
	for _, e := range args {
		fmt.Println(e)
	}
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(divide(2, 1))
	fmt.Println(divide(1, 0))
	unfixedsized(1, 2, 3)
}

