package main

import (
	"fmt"
	"math"
	"strconv"
)

func main2() {
	var MyName int
	fmt.Print(MyName)
	fmt.Println(MyName) //使用变量

	var a int = 8
	var b = a // 自动推断b的类型
	_ = b
	c := b //第一次出现 ： 声明 相当于var
	a = c  // 第二次出现不能加：和var
	var (
		d uint16
		e int8
		f float32
		g float64
	)

	a = -5
	d = 05    //前缀0表示八进制
	a = 0o75  //前缀0o表示八进制
	a = 0xab3 //前置0x表示十六进制
	a = 5_8_7
	a = 13_000_000 //表示13兆
	f = 1.432665522366653555
	g = 34.
	m := 33.
	var n bool = true
	_, _, _, _ = d, e, f, m

	fmt.Printf("a = %d, g = %.2f, n = %t\n", a, g, n)
	fmt.Printf("f = %g, f = %e\n", f, f)
	fmt.Printf("f = %[1]f, g = %[2]f, g = %[2]g, f = %[1]e\n", f, g)

	cast()

}

// 强制类型转换

func cast() {
	// 高到低精度转换，数字小没问题
	var us uint64 = 1
	i8 := int8(us)
	fmt.Printf("i8 = %d\n", i8)
	//最高位的1变成了符号位
	ua := uint64(math.MaxUint64)
	binaryStr := strconv.FormatUint(ua, 2)
	fmt.Printf("ua的二进制是 %s\n", binaryStr)
	i64 := int64(ua)
	fmt.Printf("i64 = %d\n", i64)
	//位数丢失
	ui32 := uint32(ua)
	binaryStr = strconv.FormatUint(uint64(ui32), 2)
	fmt.Printf("ua的二进制是 %s\n", binaryStr)
	fmt.Printf("ui32 = %d\n", ui32)

	//单个字符可以转化为int
	var i int = int('中')
	fmt.Printf("i = %d\n", i)
	//bool和int不能相互转换

	//byte和int可以互相转换
	var by byte = byte(i)
	i = int(by)
	fmt.Printf("i = %d\n", i)

	//float和int可以互相转换，小数位会丢失
	var ft float32 = float32(i)
	i = int(ft)
	fmt.Printf("i = %d\n", i)

}
