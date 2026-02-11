package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type ETS struct{} //ETS跟标准库的struct{}等价（可以互换）

// 所有的空结构指向同一个地址
func allEmptyStructlsSame() {
	var a ETS
	var b ETS

	fmt.Printf("address of a %p b %p \n", &a, &b)
	fmt.Printf("size a %d b %d \n", unsafe.Sizeof(a), unsafe.Sizeof(b))
	fmt.Printf("size a %d b %d \n", reflect.TypeOf(a).Size(), reflect.TypeOf(b).Size())
}

// 空结构体的应用场景

func scenariosOfEmptyStruct() {
	set := map[int]struct{}{
		1: {},
		4: {},
		7: {},
	}
	if _, exists := set[5]; exists {
		fmt.Println("5是存在的")
	} else {
		fmt.Println("5是不存在的")
	}

	bloker := make(chan struct{}, 0)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("done")
		bloker <- struct{}{}
	}()
	<-bloker
}

func example() {
	bloker := make(chan struct{})

	// 启动子goroutine
	go func() {
		fmt.Println("子goroutine先执行")
		bloker <- struct{}{} // 发送信号
		fmt.Println("子goroutine最后执行")
	}()

	fmt.Println("主goroutine执行")
	<-bloker // 接收信号（阻塞）
	fmt.Println("主goroutine继续执行")
}

func main() {
	//allEmptyStructlsSame()
	//scenariosOfEmptyStruct()
	example()
}
