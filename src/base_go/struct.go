package main

import "fmt"

// Residence 结构体定义居住地信息
// 将作为User结构体的嵌套组件，展示结构体组合
/*
type Residence struct {
	Province string // 省份
	City     string // 城市
}
*/

// User 结构体定义用户信息
// 包含基本信息和嵌套的居住地信息，展示复杂结构体的定义
type User struct {
	Id            int     // 用户ID
	Score         float64 // 用户分数
	Name, address string  // 用户姓名和地址（两个字段在同一行声明）
	//Residence     Residence // 居住地信息，嵌套Residence结构体
	Residence struct {
		Province string
		City     string
	}
	Father *User // 父亲引用，指向另一个User实例，形成链式层级关系
}

// hello 方法是User类型的方法，用于打印用户姓名
// (mi User) 表示这个方法属于User类型，mi是接收者
func (mi User) hello() {
	fmt.Println("My Name is", mi.Name)
}

// main 程序入口函数
func main7() {
	// 声明一个User类型的变量u（零值初始化）
	var u User

	// 方式1：使用字段名初始化结构体（推荐方式）
	u = User{Id: 32, address: "中国", Name: "大铭铭"}
	u.Residence.Province = "河南"
	u.Residence.City = "洛阳"
	fmt.Println(u.Name)

	// 方式2：按字段顺序初始化结构体（不推荐，容易出错）
	// 必须按结构体字段定义顺序提供所有字段的值，包括嵌套的Residence
	/*
		u = User{32, 32.4, "大铭铭", "中国", struct {
			Province string
			City     string
		}{"河南", "洛阳"},
		}
	*/
	u = User{32, 32.4, "大铭铭", "中国", u.Residence, &User{}} // 最后一个参数&User{}是Father字段的值，创建一个零值User的指针
	// Go语言的三种格式化输出方式，展示结构体的不同显示格式
	fmt.Printf("%v\n", u)  // %v: 基本格式，紧凑显示
	fmt.Printf("%+v\n", u) // %+v: 包含字段名的详细格式
	fmt.Printf("%#v\n", u) // %#v: Go语法格式的完整表示，可用于代码生成

	// 调用User结构体的hello方法
	u.hello()

	// 匿名结构体：直接定义结构体而不命名类型
	// 适用于只需要一次使用的临时数据结构
	var student struct {
		Name string // 学生姓名
		Age  int    // 学生年龄
	}

	// 给匿名结构体的字段赋值
	student.Age = 18
	student.Name = "芳芳"
	fmt.Println(student.Name)

	// 将student的Name赋值给u的Name字段
	// 演示不同结构体变量间的数据传递
	u.Name = student.Name
	u.hello()

	// 方式3：零值初始化结构体
	// User{} 创建一个所有字段都为零值的User结构体
	u2 := User{}
	u2.address = "杭州" // 给address字段赋值

	// 方式4：指针操作
	// &u2 取u2的内存地址，u3是*User类型的指针
	// 通过指针访问字段时，Go会自动解引用
	u3 := &u2
	fmt.Println(u3.address)

	// 方式5：使用new()函数创建指针
	// new(User) 分配User大小的内存，返回指向该内存的指针
	// 内存初始化为零值，等价于 &User{}
	u4 := new(User)
	u4.Name = "张霞白"
	fmt.Println(u4.Name)
}
