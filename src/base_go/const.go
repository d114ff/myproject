package main

import (
	"fmt"
)

func main3() {
	const (
		PI = 3.14 // 常量默认名字是大写，常量不能修改默认值
		E  = 2.71
	)
	fmt.Printf("PI = %.2f, E = %.2f\n", PI, E)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Printf("a = %d,b = %d, c =%d ,d = %d\n", a, b, c, d)

	const (
		a1 = iota
		b1 = 50
		c1 = iota
		d1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d, d1 = %d\n", a1, b1, c1, d1)

	const (
		NOT_USE = 1 << (10 * iota) // iota = 0  <<<位移运算
		KB
		MB
		GB
	)
	fmt.Printf("KB = %d, MB = %d ,GB = %d\n", KB, MB, GB)

	const (
		a2, b2 = iota + 1, iota + 2
		c2, d2
		e, f
	)
	fmt.Printf("a2 = %d, b2 = %d, c2 = %d, d2 = %d, e = %d, f = %d\n", a2, b2, c2, d2, e, f)
}
