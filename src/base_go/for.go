package main

import "fmt"

func main9() {
	a, b := 6, 20
	var sum int
	for a >= 0 && b > a {
		sum += a
		if a == 4 {
			break
		}
		a, b = a-1, b/2
	}

	fmt.Println(sum, a, b)
}
