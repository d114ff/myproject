package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func asing_string() {
	//字符串里可以包含任意Unicode符号
	s1 := "My name is 王猛萌☻"
	//包含转义字符
	s2 := "He say:\"I'm fine.\"\n\\Thank\tyou\\"
	//反引号里的转移字符不会被转义。反引号里的原封不动地输出，包括空白符和换行符
	s3 := `He say:"I'm fine.

\Thank\you .
`
	fmt.Println("s1:")
	fmt.Println(s1)
	fmt.Println("s2:")
	fmt.Println(s2)
	fmt.Println("s3:")
	fmt.Println(s3)

}

func string_impl() {
	/*
		英文字符：M, y,  , n, a, m, e,  , i, s,   → 共 11 个 ASCII 字符（每个占 1 字节）
		中文字符：王、猛、萌 → 每个 UTF-8 编码占 3 字节
		特殊符号：☻（U+263B，BLACK SMILING FACE）→ UTF-8 编码为 3 字节
		英文、数字、空格 → 1 字节/字符
		中文、希腊字母、☻（U+263B）→ 3 字节/字符
		表情 😂（U+1F602）→ 4 字节/字符

	*/
	s1 := "My name is 王猛萌"
	arr := []byte(s1)
	brr := []rune(s1)
	fmt.Printf("last byte : %d\n", arr[len(arr)-1]) // string可以转换为[]byte切片或者[]rune切片
	fmt.Printf("last byte : %c\n", arr[len(arr)-1]) // %c以unicode字符格式输出
	//arr[len(arr)-1] 只拿到了最后一个字节，这不是一个完整的中文字符.
	fmt.Printf("last rune : %d\n", brr[len(brr)-1])
	fmt.Printf("last rune : %c\n", brr[len(brr)-1])
	L := len(s1)
	fmt.Printf("string len %d byte array len %d rune array len %d\n", L, len(arr), len(brr))

	for i := 0; i < L; i++ {
		fmt.Printf("%c", s1[i]) //[i]前面应该出现数组或切片，这里自动把string转成了[]byte
	}
	fmt.Println()

	for _, ele := range s1 {
		fmt.Printf("%c", ele) //按rune进行遍历输出
	}
	fmt.Println()
	arr[0] = 9
	//s[0] = 9 //字符串不能修改
	fmt.Println(utf8.RuneCountInString(s1), len([]rune(s1))) //查看string里有几个rune
}

// 字符串拼接
func string_join() {
	s1 := "Hello"
	s2 := "how"
	s3 := "are"
	s4 := "you"
	merged := s1 + " " + s2 + " " + s3 + " " + s4
	fmt.Println(merged)
	merged = fmt.Sprintf("%s %s %s %s", s1, s2, s3, s4)
	fmt.Println(merged)
	merged = strings.Join([]string{s1, s2, s3, s4}, " ")
	fmt.Println(merged)
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	sb.WriteString(" ")
	sb.WriteString(s3)
	sb.WriteString(" ")
	sb.WriteString(s4)
	sb.WriteString(" ")
	merged = sb.String()
	fmt.Println(merged)
}

func string_op() {
	s := "born to win, born to die."
	fmt.Printf("sentence length %d\n", len(s))
	fmt.Printf("\"s\" length %d \n", len("s")) //英文字母的长度为1
	fmt.Printf("\"中\" length %d \n", len("中")) //一个汉字占3个长度
	arr := strings.Split(s, " ")
	fmt.Printf("arr[3]=%s\n", arr[3])
	fmt.Printf("contain die %t\n", strings.Contains(s, "die")) //包含子串
	fmt.Printf("contain wine %t\n", strings.Contains(s, "wine"))
	fmt.Printf("first index of born %d\n", strings.Index(s, "born"))    //寻找子串第一次出现的位置
	fmt.Printf("last index of born %d\n", strings.LastIndex(s, "born")) //寻找子串最后一次出现的位置
	fmt.Printf("begin with born %t\n", strings.HasPrefix(s, "born"))    //以xxx开头
	fmt.Printf("end with die. %t\n", strings.HasSuffix(s, "die."))      //以xxx结尾
	fmt.Println(strings.Repeat("-", 50))                                //字符串重复N次
}

func string_other_convert() {
	/*
		整数转字符串：strconv.Itoa（适合 int）或 strconv.FormatInt（适合 int64）。
		字符串转整数：strconv.Atoi（转 int）或 strconv.ParseInt（转 int64，支持更多选项）。
	*/
	var err error
	var i int = 8
	var i64 int64 = int64(i)
	//int 转 string
	var s string = strconv.Itoa(i)
	fmt.Printf("int转string结果: 类型=%T, 值=%v\n", s, s)
	s = strconv.FormatInt(i64, 10) // 参数10表示使用十进制，需要几进制写几进制
	fmt.Printf("int64转string结果: 类型=%T, 值=%v\n", s, s)
	// string 转 int
	i, err = strconv.Atoi(s)
	if err != nil {
		fmt.Printf("转换失败 %s\n", err)
	} else {
		fmt.Printf("string转int结果: 类型=%T, 值=%v\n", i, i) // 自动显示int类型和值
	}
	// string 转 int64
	i64, err = strconv.ParseInt(s, 10, 64) // 10同上，参数64表示使用64位整数
	if err != nil {
		fmt.Printf("转换失败 %s\n", err)
	} else {
		fmt.Printf("string转int64结果: 类型=%T, 值=%v\n", i64, i64) // 自动显示int64类型和值
	}

	// float转string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) //保留2位小数  %.2f 'f'：固定小数点格式
	fmt.Printf("float64转string结果: 类型=%T, 值=%v\n", s, s)
	// string 转 float
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("转换失败 %s\n", err)
	} else {
		fmt.Printf("string转float64结果: 类型=%T, 值=%v\n", f, f)
	}
	//string<-->[]byte
	var arr []byte = []byte(s)
	fmt.Printf("[]byte: 类型=%T, 值=%v\n", arr, arr)
	s = string(arr)
	fmt.Printf("[]byte转string结果: 类型=%T, 值=%v\n", s, s)
	var brr []rune = []rune(s)
	fmt.Printf("[]rune: 类型=%T, 值=%v\n", brr, brr)
	s = string(brr)
	fmt.Printf("[]rune转string结果: 类型=%T, 值=%v\n", s, s)
}
func main18() {
	//asing_string()
	//string_impl()
	//string_join()
	//string_op()
	string_other_convert()
}
