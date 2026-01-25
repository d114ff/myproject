package main

import "fmt"

func main5() {
	// 输出字母 A 和 Z 的 ASCII 码值
	fmt.Printf("A = %d Z = %d\n", 'A', 'Z')
	var base int = 'Z' - 'A' + 1 // 计算 Excel 列号系统的进制数（26进制，但数字从1-26）
	fmt.Println(base, "进制")

	// 计算三字母列号 "DFX" 的十进制值（列索引）
	// Excel 列号系统：A=1, B=2, ..., Z=26, AA=27, AB=28, ..., XFD=16384
	var total int
	// 第3位（最右边）: D = 第4个字母 = 4 × 26⁰ = 4
	total += 'D' - 'A' + 1
	// 第2位: F = 第6个字母 = 6 × 26¹ = 156
	total += base * ('F' - 'A' + 1)
	// 第1位（最左边）: X = 第24个字母 = 24 × 26² = 16224
	total += base * base * ('X' - 'A' + 1)
	// 总计: 4 + 156 + 16224 = 16384
	fmt.Println("total", total)
}
