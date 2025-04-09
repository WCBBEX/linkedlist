package main

import (
	"fmt"
	"github.com/WCBBEX/linkedlist"
)

func main() {
	l := linkedlist.NewWithSlice([]int{1, 2, 3, 4, 5})

	v, err := l.PopBack()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	println(*v)

	a := 1546
	l.PushFront(&a)
	for i := range l.Iter() {
		println(*i)
	}
}
