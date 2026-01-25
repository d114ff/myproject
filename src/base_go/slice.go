package main

import "fmt"

func slice_init() {
	var s []int //切片声明，len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{} //初始化，len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3) // 初始化，len=cap=3
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) //初始化 len=3,cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5} //初始化，len=cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	fmt.Println("===========================")
	//二维切片初始化
	s2d := [][]int{
		{1},
		{2, 3}, // 二维数字各行的列数是相等的，但二维切片各行的len可以不等.
	}
	fmt.Printf("s2d len %d cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len %d cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1] len %d cap %d\n", len(s2d[1]), cap(s2d[1]))
	fmt.Println("===========================")
}

func slice_append() {
	arr := make([]int, 3, 6)
	brr := append(arr, 8) //arr and brr 共享底层数组，但它们的len不同
	brr[0] = 9

	fmt.Printf("arr[0] = %d, cap of arr %d, len of arr %d\n", arr[0], cap(arr), len(arr))
	fmt.Printf("brr[0] = %d, cap of brr %d, len of brr %d\n", arr[0], cap(brr), len(brr))
	fmt.Println("brr[3] = ", brr[3])
	fmt.Println("===========================")
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	s = append(s, 4, 5) // 可以一次append多个元素
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	// 容量不够用了，需要申请新的内存地址，把老数据拷贝被到新内存是上执行
	s = append(s, 6)
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	fmt.Println("===========================")
}

func sub_slice() {
	//截取一部分创造子切片，现在字母切片共享内存，母切片容量子切片可以使用
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	fmt.Printf("s[1] address %p\n", &s[1])
	sub_slice := s[1:3] //从切片创造子切片，len=cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	// 母切片容量允许子切片append操作
	sub_slice = append(sub_slice, 6, 7)
	sub_slice[0] = 8
	fmt.Printf("s=%v,sub_slice=%v, s[1] address %p , sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])
	// 母切片用完了，子切片重新申请数据，把老数据拷贝到新内存
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("s=%v,sub_slice=%v, s[1] address %p , sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr[1] address %p\n", &arr[1])
	sub_slice = arr[1:3] // 从数组创造子切片，len=2 ，cap=4
	//母体数组允许子切片执行append操作
	sub_slice = append(sub_slice, 6, 7)
	sub_slice[0] = 8
	fmt.Printf("arr=%v,sub_slice=%v, arr[1] address %p , sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])

	// 母体数组的容量用完，子切片执行append会申请新内存
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("arr=%v,sub_slice=%v, arr[1] address %p , sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])

	array := [...]int{1, 2, 3, 4, 5}
	brr := array[:] // 截取数组得到切片
	fmt.Printf("%v,%T\n", brr, brr)
	crr := brr[1:2:4] //len=2-1, cap=4-1
	fmt.Printf("len(crr) = %d, cap(crr) = %d\n", len(crr), cap(crr))
}

// 清空slice
func clear_slice(arr *[]int) {
	*arr = []int{}
}

// 探究capacity扩容规律

func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			fmt.Printf("capacity从%d 变成 %d \n", prevCap, currCap)
			prevCap = currCap
		}
	}
	fmt.Println("===========================")
}

// 由于传的是底层数组的指针，可以直接修改

func update_slice(s []int) {
	s[0] = 888
}

func main16() {
	slices := []int{1, 2, 3, 4, 5}
	//slice_init()
	//slice_append()
	//sub_slice()
	//clear_slice(&slices)
	//expansion()
	update_slice(slices)
	fmt.Printf("%#v\n", slices)
}
