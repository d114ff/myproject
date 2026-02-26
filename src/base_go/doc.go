package main

import "fmt"

func addd() func() int {
	a := 3
	fmt.Printf("in addd func,address of a is %p\n", &a)
	return func() int {
		fmt.Printf("in addd func,address of a is %p\n", &a)
		a++
		return a
	}
}

func main38() {
	af := addd()
	fmt.Println(af())
	fmt.Println(af())

	bf := addd()
	fmt.Println(bf())
	fmt.Println(bf())
}
