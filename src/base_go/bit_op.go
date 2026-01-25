package main

import "fmt"

func main12() {
	// 定义一个uint64变量a，值为20
	var a uint64 = 20
	// 使用Go的格式化输出直接打印20的二进制形式
	fmt.Printf("a=%b\n", a)
	// 调用自定义函数手动显示20的8位二进制格式
	binaryFormat(20)
}

// binaryFormat 手动实现8位二进制数的显示
// 参数a: 要转换为二进制的uint8数字
func binaryFormat(a uint8) {
	// c初始化为10000000(二进制)，即最高位为1的掩码
	var c uint8 = 1 << 7
	// 循环8次，检查每一位
	for i := 0; i < 8; i++ {
		// 位与运算：如果c&a == c，说明a的当前位是1
		if c&a == c {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		// 右移一位，检查下一个位
		c = c >> 1
	}
	// 打印换行符
	fmt.Println()
}
