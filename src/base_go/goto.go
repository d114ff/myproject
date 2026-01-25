package main

import "fmt"

// basic_goto 演示基本的goto循环用法
func basic_goto() {
	var i int = 4
MY_LABEI: // 定义标签，注意这里有拼写错误（应该是LABEL）
	i += 3
	i *= 2
	if i > 200 {
		return // 当i大于200时退出函数
	}
	fmt.Println(i)
	goto MY_LABEI // 跳转回标签处，形成循环
}

// if_goto 演示在if-else语句中使用goto进行条件跳转
func if_goto() {
	var i int = 4
	if i%2 == 0 {
		goto L1 // 如果i是偶数，跳转到L1标签
	} else {
		goto L2 // 如果i是奇数，跳转到L2标签
	}
L1:
	i += 3
	fmt.Println("L1 = ", i)
L2:
	i *= 3
	fmt.Println("L2 = ", i)
}

// for_goto 演示在嵌套for循环中使用goto进行复杂跳转
func for_goto() {
	const SIZE = 5
L1: // 外层循环标签
	for i := 0; i < SIZE; i++ {
	L2: // 内层循环标签
		fmt.Printf("开始检查第%d行\n", i)
	L3: // 条件判断标签
		if i%2 == 1 { // 如果i是奇数
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					goto L1 // 跳转到外层循环开始处
				} else if j%3 == 1 {
					goto L2 // 跳转到内层循环开始处
				} else {
					goto L3 // 跳转到条件判断处
				}
			}
		}
	}
}

// continue_goto 演示使用continue配合标签进行跳转
func continue_goto() {
	const SIZE = 5
L1: // 外层循环标签
	for i := 0; i < SIZE; i++ {
	L2: // 内层循环标签
		fmt.Printf("开始检查第%d行\n", i)
	L3: // 条件判断标签
		if i%2 == 1 { // 如果i是奇数
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					continue L1 // 跳转到外层循环的下一次迭代
				} else if j%3 == 1 {
					goto L2 // 跳转到内层循环开始处
				} else {
					goto L3 // 跳转到条件判断处
				}
			}
		}
	}
}

func break_goto() {
	const SIZE = 5
L1: // 外层循环标签
	for i := 0; i < SIZE; i++ {
	L2: // 内层循环标签
		fmt.Printf("开始检查第%d行\n", i)
	L3: // 条件判断标签
		if i%2 == 1 { // 如果i是奇数
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					break L1 // 跳转到外层循环的下一次迭代
				} else if j%3 == 1 {
					goto L2 // 跳转到内层循环开始处
				} else {
					goto L3 // 跳转到条件判断处
				}
			}
		}
	}
}

func main10() {
	//basic_goto()  // 基本goto循环示例
	//if_goto()     // if-else中的goto示例
	//for_goto()    // 嵌套for循环中的goto示例
	//continue_goto() // continue配合标签的示例
	break_goto()
}
