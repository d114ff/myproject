package main

import "fmt"

func array1d() {
	var arr1 [5]int = [5]int{} //数组必须指定长度和类型，且长度类型指定后不能改变
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}               // 给前面两个元素赋值
	var arr4 = [5]int{2: 15, 4: 30}       // 给指定index赋值
	var arr5 = [...]int{3, 4, 5, 2, 1, 8} // 根据{}里的元素个数推断出数组长度
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}} //数组的元素类型有匿名结构体给定
	fmt.Printf("arr1=%#v\n", arr1)
	fmt.Printf("arr2=%#v\n", arr2)
	fmt.Printf("arr3=%#v\n", arr3)
	fmt.Printf("arr4=%#v\n", arr4)
	fmt.Printf("arr5=%#v\n", arr5)
	fmt.Printf("arr6=%v\n", arr6)
	// 通过index访问数组里的元素
	fmt.Printf("arr5[0] = %d\n", arr5[0])
	fmt.Printf("arr5[%d]=%d\n", len(arr5)-1, arr5[len(arr5)-1])
	fmt.Printf("数组的地址：%p\n", &arr5)
	//遍历数组
	for i, ele := range arr5 {
		fmt.Printf("index = %d element = %d\n", i, ele)
	}
	fmt.Println("---------------------")
	// 第二种遍历
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("index = %d element = %d\n", i, arr5[i])
	}
	// 数组长度容量不变
	fmt.Printf("len(arr5) = %d\n", len(arr5))
	fmt.Printf("cap(arr5) = %d\n", cap(arr5))
}

func array2d() {
	var arr1 = [5][3]int{{1}, {2, 3}} //5列3行，只给前两行赋值
	var arr2 = [...][3]int{{1}, {2, 3}, {3, 4, 5}, {5, 6, 7}, {8, 9, 10}}
	//根据index访问数组里的元素
	fmt.Printf("arr[1][1]=%d\n", arr1[1][1])
	fmt.Printf("arr[4][2]=%d\n", arr2[4][2])
	// 遍历二维数组
	for row, array := range arr2 { //先取出某一行
		for col, ele := range array { //再遍历这一行
			fmt.Printf("arr2[%d],[%d] = %d\n", row, col, ele)
		}
	}
	//对于多维数组，其cap和len指第一维的长度
	fmt.Printf("len(arr1) = %d\n", len(arr1))
	fmt.Printf("cap(arr1) = %d\n", cap(arr1))
}

// go语言中函数传参都是按值传递，即传递数组实际上传的是数组的拷贝
func update_array1(arr [5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 8888
}

// 传数组的指针
func update_array2(arr *[5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 8888 // 传数组指针可以在原来空间内存直接改
}

func update_array3(arr [5]*int) {
	*arr[0] = 88 //因为传的是int指针，所以直接在原来的内存空间上进行修改

}

// for range取得的是值拷贝

func for_range_arry() {
	arr := [...]int{1, 2, 3}
	for i, ele := range arr { // ele 是 arr中元素的拷贝
		arr[i] += 8 // 修改 arr里的元素，不影响ele
		fmt.Printf("i = %d，arr[%d] = %d ele = %d\n", i, i, arr[i], ele)
		ele += 1 // 修改 ele不影响arr
		fmt.Printf("i = %d，arr[%d] = %d ele = %d\n", i, i, arr[i], ele)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i = %d，arr[%d] = %d\n", i, i, arr[i])
	}
}

func main15() {
	//array1d()
	//array2d()
	//arr := [5]int{5, 4, 3, 2, 1}
	//update_array1(arr)
	//fmt.Println(arr[0])
	//update_array2(&arr)
	//fmt.Println(arr[0])
	//a, b, c := 1, 2, 3
	//brr := [5]*int{&a, &b, &c}
	//update_array3(brr)
	//fmt.Println(*brr[0])
	//for_range_arry()
}
