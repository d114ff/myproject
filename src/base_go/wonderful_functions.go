package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{3, 7, 4, 6, 8, 1}
	slices.Sort(arr)
	fmt.Println(arr)                          // 按照顺序排序
	slices.SortFunc(arr, func(a, b int) int { // 自定义排序
		//return b - a  // b - a 从大到小
		return a - b // a - b 从小到大
	})
	fmt.Println(arr)

	// 结构体排序
	type User struct {
		Age    int
		Height float32
	}

	brr := []*User{&User{18, 1.8}, &User{25, 1.7}}
	slices.SortFunc(brr, func(a, b *User) int {
		// return int(b.Height - a.Height) 因为返回值是int类不能减掉
		if b.Height > a.Height {
			return 1
		} else if b.Height < a.Height {
			return -1
		} else {
			return 0
		}
	})
	for i := 0; i < len(brr); i++ {
		fmt.Printf("索引: %d, 年龄: %d, 身高: %.2f\n", i, brr[i].Age, brr[i].Height)
	}

	fmt.Println("最大值", slices.Max(arr))
	fmt.Println("最小者", slices.Max(arr))
	fmt.Println("包含", slices.Contains(arr, 8))

	crr := make([]int, len(arr))
	copy(crr, arr) //最多只能拷贝 len(crr) 个元素，性能比自己写for循环要高很多
	fmt.Println(crr)
	fmt.Println("EQ", slices.Equal(arr, crr))
	arr[0]++
	fmt.Println(arr)
	fmt.Println("EQ", slices.Equal(arr, crr))

	fmt.Println("----------------------------------------")
	drr := arr
	fmt.Println("相等", slices.Equal(drr, arr)) //true
	arr[0]++
	fmt.Println(arr)
	fmt.Println("相等", slices.Equal(drr, arr)) //true

}
