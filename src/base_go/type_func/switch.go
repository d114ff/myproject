package main

import (
	"fmt"
	"strings"
)

func switch_basic() {
	color := "yellow"
	switch color {
	case "green":
		fmt.Println("go")
	default:
		fmt.Println("stop")
	}

	color = "green"
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("stop")
	default:
		fmt.Println("以上都不是")

	}
	color = "black"
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("stop")
	default:
		fmt.Printf("invalid traffic singal: %s\n", strings.ToUpper(color))
	}
}

func add(a int) int {
	return a + 10
}

func switch_expression() {
	var a int = 5
	switch add(a) {
	case 15:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}

	const B = 15
	switch B {
	case add(a):
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}

func switch_condition() {
	color := "yellow"
	switch color {
	case "green":
		fmt.Println("go")
	case "red", "yellow":
		fmt.Println("stop")
	}

	switch {
	case add(5) > 10:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}

func fall_throth(age int) {
	fmt.Printf("您的年龄是%d,您可以：\n", age)
	switch {
	case age > 50:
		fmt.Println("出任国家首脑")
		fallthrough
	case age > 25:
		fmt.Println("生儿育女")
		fallthrough
	case age > 22:
		fmt.Println("结婚")
		fallthrough
	case age > 18:
		fmt.Println("开车")
		fallthrough
	case age > 16:
		fmt.Println("参加工作")
		fallthrough
	case age > 15:
		fmt.Println("上高中")
		fallthrough
	case age > 3:
		fmt.Println("上幼儿园")
	}
}

func switch_type() {

}

func main11() {
	//switch_basic()
	//switch_expression()
	//switch_condition()
	//fall_throth(30)
}
