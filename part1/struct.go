package main

import "fmt"

type A struct {
	name string
	num int
}

func newA() *A {
	a := A{}
	a.name = "hwan"
	a.num = 10

	return &a
}

func (a *A) print() {
	fmt.Println(a)
}

func main() {
	a := newA()

	a.print() 
}