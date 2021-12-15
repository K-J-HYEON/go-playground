package main

import "fmt"

type Stringer interface {
	String1() string
}

type Student struct {
	Name string
	Age int
}

func (s Student) String() string {
	return fmt.Sprintf("Hello! I'm %dyears old %s, nice to meet you!", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = studet
	fmt.Printf("%s\n", stringer.String())
}