package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age int
}

type Students []Student

func (s Students) Len() int { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age}
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i]}

func main() {
	s := []Student{
		{"진", 33},
		{"폴", 30},
		{"화랑", 27},
		{"백두산", 60},
		{"헤이야치", 70},
	}
	sort.Sort(Students(s))
	fmt.Println(s)
}