package main

import "fmt"

func main2() {
	var MyName int
	fmt.Print(MyName)
	fmt.Println(MyName) //使用变量

	var a int = 8
	var b = a // 自动推断b的类型
	_ = b
	c := b //第一次出现 ： 声明 相当于var
	a = c  // 第二次出现不能加：和var
	var (
		d uint16
		e int8
		f float32
		g float64
	)

	a = -5
	d = 05    //前缀0表示八进制
	a = 0o75  //前缀0o表示八进制
	a = 0xab3 //前置0x表示十六进制
	a = 5_8_7
	a = 13_000_000 //表示13兆
	f = 1.432665522366653555
	g = 34.
	m := 33.
	var n bool = true
	_, _, _, _ = d, e, f, m

	fmt.Printf("a = %d, g = %.2f, n = %t\n", a, g, n)
	fmt.Printf("f = %g, f = %e\n", f, f)
	fmt.Printf("f = %[1]f, g = %[2]f, g = %[2]g, f = %[1]e\n", f, g)

}
