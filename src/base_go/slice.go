package main

import "fmt"

func slice_init() {
	var s []int //切片声明，len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
}

func main() {
	slice_init()
}
