package main

import "fmt"

func main6() {
	// 定义User结构体，表示用户信息
	// 这是基础结构体，将被其他结构体组合使用
	type User struct {
		Name string // 用户姓名
		Age  int    // 用户年龄
	}

	// 定义Vedio结构体，表示视频信息
	// 使用匿名嵌套User结构体，实现Go语言的"继承"效果
	// 匿名嵌套会自动提升内层结构体的字段和方法到外层
	type Vedio struct {
		Length int    // 视频长度（分钟）
		Name   string // 视频名称
		User          // 匿名嵌套：直接写类型名，不指定字段名
	}

	// 创建User实例
	// 使用字段名初始化方式，代码更清晰
	u := User{Name: "大洋", Age: 18}

	// 创建Vedio实例
	// 匿名嵌套的初始化：使用类型名User作为字段名
	v := Vedio{Length: 120, Name: "go学习教程", User: u}

	// 访问Vedio结构体的直接字段
	fmt.Println(v.Length) // 输出：120
	fmt.Println(v.Name)   // 输出：go学习教程

	// 访问匿名嵌套结构体User的字段
	// 方式1：完整访问 v.User.Name（明确指定嵌套结构体）
	fmt.Println(v.User.Name) // 输出：大洋

	// 方式2：字段提升 v.Age（直接访问，Go自动提升内层字段）
	// 这是匿名嵌套的核心优势：语法更简洁，类似传统继承
	fmt.Println(v.Age) // 输出：18
}
