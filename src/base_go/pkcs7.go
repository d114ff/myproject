package main

import "fmt"

func Padding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := blockSize - srcLen%blockSize
	rect := src
	for i := 0; i < padLen; i++ {
		rect = append(rect, byte(padLen))
	}
	return rect
}

func Unpadding(src []byte, blockSize int) []byte {
	srcLen := len(src)
	padLen := int(src[srcLen-1])
	return src[0 : srcLen-padLen]
}

func main() {
	src := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Padding(src, 8))
	fmt.Println(Unpadding(Padding(src, 8), 8))
}
