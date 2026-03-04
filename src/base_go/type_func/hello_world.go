// Package main 定义了一个可独立执行的程序
package main

import "fmt" // 导入fmt包，用于格式化输入输出

// main1 函数演示基本的Go程序输出
// 注意：函数名不是main，所以不会作为程序入口点
func main1() {
	fmt.Println("hello go")               // 打印字符串并换行
	fmt.Print("调出命令行 ctrl+~，删除ctrl+w \n") // 打印快捷键提示信息
}
