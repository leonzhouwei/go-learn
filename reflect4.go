package main

import (
	"fmt"
	"reflect"
)

type E struct {
	A int
	B string
}

type T struct {
	A int
	B string
	E E
}

func main() {
	t := T{203, "mh203", E{0, "mh0"}}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	// modify the A filed of T via reflect
	// fieldA := s.Field(0)
	// fieldA.SetInt(204)
	// fieldB := s.Field(1)
	// fieldB.SetString("mh204")
	fieldE := s.Field(2)
	replacement := E{1, "mh1"}
	fieldE.Set(reflect.ValueOf(replacement))
	fmt.Println(t)
}