package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		p := recover()
		if p == nil {
			return
		}
		err, ok := p.(error)
		if ok {
			fmt.Println("%#v\n", err)
			return
		}
		panic(p)
	}()
	panic(errors.New("ERROR"))
}