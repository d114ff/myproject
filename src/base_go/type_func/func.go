package main

import (
	"fmt"
	"time"
)

// a,b是形参，邢参是函数内部的局部变量，实参的值会拷贝给邢参
func arg1(a, b int) { // 参数类型相同时可以只写一次
	a = a + b // 在函数内部修改形参的参数值，不影响实参
	return    // 函数返回值，return后面的语句不会再执行
	fmt.Println("我不会被输出")
}

func arg2(a, b int) {
	a = a + b
	// 不写return，默认执行完最后一行代码函数返回
}

func arg3(a, b *int) { // 如果想通函数修改实参，需要指针类型
	*a = *a + *b
	*b = *a * *b
}

func no_arg() { // 函数可以没有参数，也没有返回值
	fmt.Println("欢迎来到goland之旅")
}

func return1(a, b int) int { // 函数需要返回一个int类型数据
	a = a + b
	c := a
	return c
}

func return2(a, b int) (c int) { // 返回值变量c已经声明好了
	a = a + b
	c = a    // 可以直接使用C
	return c // 由于函数要求有返回值，即使C赋值了，也需要显式写return
}

func return3() (int, int) { // 可以没有形参，可以返回多个参数
	now := time.Now()
	return now.Hour(), now.Hour()
}

// 不定长参数
func variable_ength_arg(a int, other ...int) int { //调用该函数时，other可以对应0个参数也可以对应多个参数
	sum := a
	//不定长参数实际上是slice类型
	for _, ele := range other {
		sum += ele
	}
	if len(other) > 0 {
		fmt.Printf("first ele %d len %d cap %d\n", other[0], len(other), cap(other))
	} else {
		fmt.Printf("len %d cap %d\n", len(other), cap(other))
	}
	return sum
}

func sum(arr ...int) int {
	s := 0
	if len(arr) == 0 {
		return s
	}
	s = s + arr[0]
	s = s + sum(arr[1:]...)
	return s

}

func main13() { // 13
	//c, d := 3, 5
	//fmt.Println(c, d)
	//arg1(c, d)
	//arg3(&c, &d)
	//fmt.Println(c, d)
	//fmt.Println(return1(c, d))
	//sum := variable_ength_arg(1, 2, 3, 4)
	sum := sum(1, 2, 3)
	fmt.Println(sum)
}
