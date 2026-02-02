package main

import "fmt"

/*
何时会发生panic:
1. index out of range和divide by zero是常见的runtime error,会发生panic
2. 通常系统初始化发生重大问题会主动调用panic(any)

panic时会依次执行：
1. 执行已经注册的defer(后注册的先执行)，未注册的defer不执行
2. 打印错误信息和堆栈调用信息
3. 调用os.Exit(2)结束go进程
*/
func defer_panic() {
	defer fmt.Println("1")
	var arr []int
	n := 0 //会执行
	//defer fmt.Println(1 / n) //在注册defer时就要计算1/n，发生panic
	defer func() { //在defer后跟func时，defer注册时不计算1/n，而是在panic发生时调用func时才计算
		_ = arr[n]
		_ = 1 / n
		fmt.Println("2") //不会执行
	}()

	defer fmt.Println("3")
}

// recover()一定要放在函数的最开始位置
func soo() {
	// defer recover()//recover()必须放在defer func() {}里,不能直接放在defer后面,否则不生效
	defer func() {
		if panicinfo := recover(); panicinfo != nil {
			fmt.Println("soo函数中发生了panic:", panicinfo)
		}
	}()
	fmt.Println("enter soo")
	fmt.Println("regist recover")

	defer fmt.Println("hello")
	defer func() {
		n := 0
		_ = 3 / n                  //发生panic
		defer fmt.Println("world") //不会执行
	}()
}

func main20() {
	//defer_panic()
	soo()
}
