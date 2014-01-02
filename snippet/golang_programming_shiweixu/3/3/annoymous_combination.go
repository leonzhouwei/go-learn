package main

import "fmt"
import "log"
import "os"

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

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Start() {
	job.Println("starting now...")
	job.Println("started")
}

func main() {
	foo := new(Foo)
	foo.Bar()
	out, err := os.Create("3.3.log")
	if (err != nil) {
		os.Exit(-1)
	}
	prefix := "oops: "
	flag := 0
	log := log.New(out, prefix, flag)
	job := &Job{"cmd", log}
	job.Start()
}