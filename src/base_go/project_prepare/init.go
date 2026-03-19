package projectprepare

import (
	"fmt"
	"regexp"
)

var Reg *regexp.Regexp

//在init（）函数内不依赖外部的其他任何变量

func init() {
	var err error
	Reg, err = regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("init Reg success")
	}
}

func init() {
	fmt.Println("init是否匹配正则表达式", Reg.Match([]byte("hello"))) // 上面init执行完才执行这个
	//对外部有依赖，不能使用init函数
}
