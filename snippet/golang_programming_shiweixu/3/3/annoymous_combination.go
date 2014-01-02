package main

import "fmt"

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base.Foo()")
}

func (base *Base) Bar() {
	fmt.Println("Base.Bar()")
}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	fmt.Println("Foo.Bar()")
	foo.Base.Bar()
}

func main() {
	foo := new(Foo)
	foo.Bar()
}