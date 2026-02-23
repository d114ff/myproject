package main

import (
	"bytes"
	"errors"
	"fmt"
)

func Padding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := blockSize - srcLen%blockSize
	//rect := src
	//for i := 0; i < padLen; i++ {
	//rect = append(rect, byte(padLen))
	//}
	pad := bytes.Repeat([]byte{byte(padLen)}, padLen)
	/*
		[]int{10, 20, 30}：创建一个包含 3 个元素的 []int 切片，内容是 [10, 20, 30]
		[]byte{byte(padLen)} 就变成了 []byte{7}
		bytes.Repeat(pattern []byte, count int) []byte 这个函数就像一个“复印机”，
		它会把 pattern（模板）复印 count 次，然后把所有复印件连在一起，给你一个全新的切片。
		pad... 中的三个点 ... 叫做“展开操作符”。它的作用是把 pad 这个切片“打开”，
		把里面的所有元素一个个地拿出来。
	*/
	rect := append(src, pad...)
	return rect
}

func Unpadding(src []byte, blockSize int) ([]byte, error) {
	srcLen := len(src) // 获取填充后数据的总长度
	if srcLen%blockSize != 0 || srcLen <= blockSize {
		// 如果填充是正确的，原始数据经过填充后，总长度一定是 blockSize 的整数倍。
		// 检查长度是否大于块大小
		return nil, errors.New("参数问题")
	}
	padLen := int(src[srcLen-1])       // 读取最后一个字节的值，该值表示填充了多少个字节
	return src[0 : srcLen-padLen], nil // 截取从开头到去除填充部分的字节数组
}

func main36() {
	src := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Padding(src, 8))
	fmt.Println(Unpadding(Padding(src, 8), 8))
	src = []byte{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Padding(src, 8))
	fmt.Println(Unpadding(Padding(src, 8), 8))
}
