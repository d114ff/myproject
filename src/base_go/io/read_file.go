package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file faied:%v\n", err)
	} else {
		defer fin.Close()
		bs := make([]byte, 100)
		fin.Read(bs)
		fmt.Println(string(bs))

		fin.Seek(10, 0) //后面0是首位，10是跳过当前10各字节，在首位跳过当前10各字节。
		fin.Read(bs)
		fmt.Println(string(bs))

		fin.Seek(0, 0) //不加这个seek上面把文件读空了，让它回到文件初始位置
		const BATCH = 10
		buffer := make([]byte, BATCH)
		for {
			n, err := fin.Read(buffer) // 每次读10个字节
			if n > 0 {
				fmt.Print(string(buffer[0:n]))
			}
			if err == io.EOF {
				break
			}
		}
	}
}

func ReadFileWithBuffer() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\n") //删除右边换行符
				fmt.Println(line)
			}
			if err == io.EOF {
				break
			}
		}
	}
}
