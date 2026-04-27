package io

import (
	"bytes"
	"fmt"
	"os/exec"
)

func SysCall() {
	//查看系统命令所在的目录，确保命令已安装
	cmd_path, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("could not found command go")
	}
	fmt.Printf("command go in path %s\n", cmd_path)

	cmd := exec.Command("go", "version") //相当于命令go version，注意Command的每一个参数都不能包含空格

	//cmd.Output()运行命令并获得其输出结果

	if output, err := cmd.Output(); err != nil {
		fmt.Println("ot output failed", err)
	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("python3.12", "/data/gopath/myproject/src/base_go/io/hello.py")

	if output, err := cmd.Output(); err != nil {
		fmt.Println("python execute failed", err)
	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("rm", "/data/gopath/myproject/src/base_go/data/biz.log")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ":" + stderr.String())
	} else {
		fmt.Println(out.String())
	}
}
