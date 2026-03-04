package main

import "fmt"

func main8() {
	if 5 > 9 {
		fmt.Printf("A")
	}
	var a int = 10

	if a > 5 {
		fmt.Println("B")
	}

	if b := 8; b > a {
		fmt.Println("b > a")
	} else if 3 < b {
		fmt.Println("3 < b")
	}
	if c, d, e := 1, 4, 7; c < d && (c > e || c < 3) {
		fmt.Println("D")
	}
}
