package main

import "fmt"

type Age int    //Age 和 int 可以互相做强制类型转换。存储的数据类型（成员变量）是一样的，但行为（成员方法）是不一样
type Tall = int //Tall 和 int完全等价，不需要显式做类型转换
type Odt Shiper //Shiper实现了Steamer接口

func (Age) move(src, dest string) (int, error) {
	return 0, nil
}

// Age实现了Transporter接口
func (Age) whistle(n int) int {
	return 0
}

func main23() {
	var a int
	var b Age
	var c Tall

	transport("BJ", "SH", b)
	c = a + 10
	b = Age(a) + 10 //字面量不需要显式转为Age类型
	fmt.Println(a, b, c)

	var o Odt
	o.tonage = 1000
	o.name = "ddq"
	o.price++
	seaTrTransport("BJ", "SH", Shiper(o)) //Odt没有实现Steamer接口

}
