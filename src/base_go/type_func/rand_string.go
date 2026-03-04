package main

import (
	"fmt"
	"math/rand/v2"
)

var letterCollection = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY/_+*中国")

// 长度生成随机数
func RandString(n int) string {
	rect := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		index := rand.IntN(len(letterCollection))
		rect = append(rect, letterCollection[index])
	}
	return string(rect)
}
func main() {
	fmt.Println(RandString(6))
	fmt.Println(RandString(6))
	fmt.Println(RandString(6))
	fmt.Println(RandString(6))
}
