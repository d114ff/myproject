package main

import (
	"fmt"
	"strings"
)

// channel.go - 展示Go语言channel的基本操作和特性

// rangeChannel - 展示channel的声明、初始化、读写操作和range遍历
func rangeChannel() {
	var ch chan int // 声明一个int类型的channel，初始为nil
	if ch == nil {
		fmt.Printf("ch is nil,ch len %d cap %d\n", len(ch), cap(ch))
	}
	//ch <- 2 // 注意：不能向nil channel里发送数据，会导致deadlock
	if len(ch) == 0 { // 引用类型未初始化时都是nil，可以对它们执行len()函数，返回0
		fmt.Println("ch length is 0")
	}

	// 创建一个容量为8的带缓冲channel
	ch = make(chan int, 8)

	// 向channel中写入数据
	ch <- 1 // 往管道里写入(send)数据
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	// 查看channel的长度和容量
	fmt.Printf("ch len %d cap %d\n", len(ch), cap(ch))

	// 从channel中读取数据
	v := <-ch // 从管道里取走数据
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
	fmt.Println()

	// 关闭channel
	close(ch)

	// 遍历并取走channel里的剩余元素
	// 当管道里已无剩余元素且没有close管道时，receive操作会一直阻塞，最终报deadlock
	// 当管道为空且被close后，for循环会自动退出
	for ele := range ch {
		fmt.Println(ele)
	}

	// 创建一个新的channel，演示只写和只读channel
	c := make(chan int, 10)
	send(c)
	recv(c)
}

// send - 演示只写channel（chan<- int），只能向channel中写入数据
func send(c chan<- int) {
	c <- 1
}

// recv - 演示只读channel（<-chan int），只能从channel中读取数据
func recv(c <-chan int) {
	v := <-c
	fmt.Printf("从只读channel中读取值: %d\n", v)
}

// slice、map、channel是go语言里的三大引用类型
// 如果只是想改变它们引用的底层数据，不需要传指针，因为传引用类型本质上传的就是底层数据的指针

// changeArray - 演示数组是值类型，函数内修改不影响原数组
func changeArray(arr [3]int) {
	arr[0]++
}

// changeSlice - 演示切片是引用类型，函数内修改会影响原切片
func changeSlice(slc []int) {
	if len(slc) > 0 {
		slc[0]++
	}
}

// changeMap - 演示map是引用类型，函数内修改会影响原map
func changeMap(mp map[int]bool) {
	mp[0] = true
}

// changeChan - 演示channel是引用类型，函数内修改会影响原channel
func changeChan(ch chan bool) {
	if cap(ch) > len(ch) { // 检查channel是否还有容量
		ch <- true
	}
}

// main - 程序入口，演示数组、切片、map和channel的传递特性
func main21() {
	//rangeChannel() // 注释掉，先演示引用类型传递
	fmt.Println(strings.Repeat("*", 20)) // 打印分隔线

	// 演示值类型传递：数组
	arr := [3]int{1, 2, 3}
	changeArray(arr)
	fmt.Println("数组传递后的值:", arr[0]) // 输出1，原数组未被修改

	// 演示引用类型传递：切片
	slc := []int{1, 2, 3}
	changeSlice(slc)
	fmt.Println("切片传递后的值:", slc[0]) // 输出2，原切片被修改

	// 演示引用类型传递：map
	mp := map[int]bool{0: false}
	changeMap(mp)
	fmt.Println("map传递后的值:", mp[0]) // 输出true，原map被修改

	// 演示引用类型传递：channel
	ch := make(chan bool, 3)
	changeChan(ch)
	fmt.Println("channel传递后的值:", <-ch) // 输出true，原channel被修改
}
