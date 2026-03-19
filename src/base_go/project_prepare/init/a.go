package main

import (
	"fmt"
	projectprepare "myproject/src/base_go/project_prepare" //先执行外面包的init函数
	_ "net/http/pprof"                                     //在线pprof，引入包没有使用加下划线，引入目的是自动执行init()函数

	_ "github.com/go-sql-driver/mysql" //注册mysql驱动
)

func init() {
	fmt.Println("init logger")
	fmt.Println("main是否匹配正则表达式", projectprepare.Reg.Match([]byte("hello123")))
}

func main() {
	fmt.Println("server start") //init函数在所有main函数之前运行
}

func init() {
	fmt.Println("init database")
}
