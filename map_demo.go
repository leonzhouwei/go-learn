package main

import "fmt"

type Person struct {
	Id string
	Name String
	Address String
}

func main() {
	var personDb map[string]Person
	personDb = make(map[string]Person)
	//
	personDb["12345"] = Person{"12345", "Tom", "Room 203,..."}
	personDb["1"] = Person{"1", "Jack", "Room 101,..."}
	//
	person, ok := personDb["1234"]
}
