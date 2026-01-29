package main

import (
	"fmt"
)

// map中的key可以是任意能够用==操作符比较的类型，即需要是comparable类型，
// 不能是函数、map、切片，以及包含上述3中类型成员变量的的struct
// map的value可以是任意类型

func key_type() {
	type f func(int) bool
	type m map[int]byte
	type s []int
	type i int
	var ml map[i]f
	fmt.Println(ml)
	/**
	函数、map、切片不能当key
		Map 的键要“老实”，数字、字符串、布尔型都可以。
		Map 的键“不老实”，切片、Map、函数统统不行。
	**/
	// var m2 map[f]bool
	// fmt.Println(m2)
	// var m3 map[m]bool
	// fmt.Println(m3)
	// var m4 map[s]bool
	// fmt.Println(m4)
	type user struct {
		scores float32 //如果scores是slice，则user不能作为map的key
	}
	u := user{3.2}
	m5 := make(map[user]interface{})
	m5[u] = 5
	fmt.Println(m5)

}

func update_map() {
	var m map[string]int                                      //声明
	m = make(map[string]int)                                  // 初始化，容量0
	m = make(map[string]int, 5)                               // 容量为5
	m = map[string]int{"语文": 0, "数学": 39, "物理": 57, "历史": 49} //初始化可以直接添加
	m["英语"] = 59                                              //往map添加K-V
	fmt.Println(m["数学"])                                      //读取key对应的value，如果key不在返回。
	delete(m, "数学")                                           //从map里删除key-value对
	fmt.Println(m["数学"])
	//取key对应的value建议，先判断key存在不
	if value, exists := m["语文"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("map不存在[语文]这个key")
	}
	// 获取map的长度，无法获取map的cap
	fmt.Printf("map里有%d对KV\n", len(m))
	//多次遍历map返回的顺序是不一样的，但相对顺序是一样的，因为每次随机选择一个开始位置，然后顺序遍历
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Println("-----------")
	//一边遍历一边修改
	for key, value := range m {
		m[key] = value + 1
	}
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Println("-----------")
	//for range取得的是值拷贝
	for _, value := range m {
		value += 1
	}
	for key, value := range m {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Println("-----------")
}

func main19() {
	//key_type()
	update_map()
}
