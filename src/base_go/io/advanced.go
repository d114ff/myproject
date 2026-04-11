package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
先建立一个心智模型：Reader 就像一根水管，数据从水管里流出来，流进你准备好的"桶"（buffer）里。
*/
func LimitReader() {
	reder := strings.NewReader("daqiaoqiao")
	limitReader := io.LimitReader(reder, 6) //limitReader截取了reader的前6个字节
	content := make([]byte, 100)
	if n, err := limitReader.Read(content); err == nil { //content是目的地，不是读content，是往content里写！
		fmt.Printf("read %s\n", string(content[:n])) //daqiao
	}
	if _, err := limitReader.Read(content); err == io.EOF {
		fmt.Println("no more data available")
	}
}

func MultiReader() {
	r1 := strings.NewReader("黄梅时节家家雨\n")
	r2 := strings.NewReader("青草池塘处处蛙\n")
	r3 := strings.NewReader("有约不来过夜半\n")
	r4 := strings.NewReader("闲敲棋子落灯花\n")
	r := io.MultiReader(r1, r2, r3, r4) //注意，这里是有序的
	io.Copy(os.Stdout, r)               //把r流拷贝到标准输出流
	//Copy(目的地, 源头)os.Stdout（目的地，终端屏幕）
}

func MultiWriter() { //一个入口 → 多个目的地
	var (
		write1 bytes.Buffer //是内存里的缓冲区，可以往里写数据，也可以读出来。
		write2 bytes.Buffer
	)
	multiWriter := io.MultiWriter(&write1, &write2) //把两个桶接在一起，形成广播结构
	multiWriter.Write([]byte("黄梅时节家家雨\n"))          //数据从入口进来，同时广播到两个桶：
	fmt.Print(write1.String())
	fmt.Print(write2.String())
	//借助于MultiWriter可以把一条日志输出到多个文件里面去
}

func TeeReader() { //TeeReader = 窃听器，数据流过去的时候，偷偷复制一份，主流不受影响。
	var writer bytes.Buffer
	reader := strings.NewReader("黄梅时节家家雨\n")
	teeReader := io.TeeReader(reader, &writer) //从reader里读取的内容既会进入teeReader，也会进入writer
	//io.Copy(os.Stdout, reader)                 //如果打开此行，则reader里的内容已经被读完了，下面2行不会输出任何内容
	io.Copy(os.Stdout, teeReader) //如果把此行注释掉，则没有任何数据经过teeReader，writer里也不会有任何内容
	fmt.Print(writer.String())    //数据流过 teeReader 的时候，自动复制一份到 writer，原来的流继续走。
}

func PipeIO() {
	reader, writer := io.Pipe() //writer的内容会直接进入reader，中间没有buffer
	go func() {
		writer.Write([]byte("hello")) //因为中间没有buffer，所以Write操作会阻塞，直到另一个协程准备从reader里读取内容
		writer.Close()
	}()
	content := make([]byte, 100)
	if n, err := reader.Read(content); err == nil {
		fmt.Printf("read %s\n", string(content[:n]))
	}
	reader.Close()
}
