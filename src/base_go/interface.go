package main

import "fmt"

// 接口是一组行为规范的集合
type Transporter interface { //定义接口。通常接口名以er结尾
	move(src string, dest string) (int, error) //方法名 (参数列表) 返回值列表
	whistle(int) int                           //方法名 (参数列表) 返回值列表
}

type Steamer interface {
	//move(string, string)(int, error)
	//whistle(int) int
	Transporter // 嵌入Transporter接口，Steamer接口也有move和whistle方法.通过匿名成员变量实现“接口继承”
	displacement() int
}

type Car struct { // 定义结构体时无需要显式声明实现了哪些接口
	name  string
	price int
}

// 只要结构体拥有接口里声明的所有方法，就称该结构体“实现了接口”
func (car *Car) move(src string, dest string) (int, error) {
	car.price++
	fmt.Printf("从%s到%s用%s运输需要%d元\n", src, dest, car.name, car.price)
	return car.price, nil

}

func move(src string, dest string, car *Car) (int, error) {
	car.price++ // 增加汽车的运输价格
	// 打印运输信息
	fmt.Printf("从%s到%s用%s运输需要%d元\n", src, dest, car.name, car.price)
	return car.price, nil // 返回更新后的价格和nil错误
}

func (car Car) whistle(n int) int {
	for i := 0; i < n; i++ {
		fmt.Printf("滴")
	}
	fmt.Printf("\n")
	return n
}

// 结构体可以拥有接口之外的方法
func (car Car) GetName() string {
	return car.name
}

type Shiper struct {
	name   string
	price  int
	tonage int
}

func (ship *Shiper) move(src string, dest string) (int, error) {
	ship.price++
	fmt.Printf("从%s到%s用%s运输需要%d元\n", src, dest, ship.name, ship.price)
	return ship.price, nil
}

func (ship Shiper) whistle(n int) int {
	for i := 0; i < n; i++ {
		fmt.Printf("呜")
	}
	fmt.Printf("\n")
	return n
}

// 一个struct可以同时实现多个接口
func (ship Shiper) displacement() int {
	return ship.tonage
}

// 接口可以作为函数参数，也可以作为结构体的成员变量（策略模式）
func transport(src, dest string, transporter Transporter) error {
	if _, err := transporter.move(src, dest); err != nil {
		return err
	} else {
		transporter.whistle(3)
		return nil
	}
}

func seaTrTransport(src, dest string, steamer Steamer) error {
	fmt.Printf("排水量%d\n", steamer.displacement())
	if _, err := steamer.move(src, dest); err != nil { //不需要价格的返回值 (ship *Shiper) move(src string, dest string) (int, error)
		return err
	} else {
		steamer.whistle(3)
		return nil
	}
}

func main22() {
	car := &Car{"宝马", 100}
	ship := &Shiper{"哈哈号", 220, 200000}

	/*
		car.name = "hhhh"
		car.price++
		car.move("北京", "上海")

		ship.name = "和平号"
		ship.price++
		ship.move("北京", "广州")
	*/
	var transporter Transporter
	transporter = car //接口赋值能否成功，取决于你赋的“那个值”的类型是否实现了接口。
	transporter.whistle(1)
	transporter = ship
	transporter.whistle(1)

	/*
		transport("北京", "上海", car)
		transport("北京", "宁波", ship)
		fmt.Println()
		seaTrTransport("北京", "天津", ship)
	*/

}
