package main

import (
	"fmt"
	"time"
)

// ============================================================
// 项目三：外卖订单系统
// 涉及：接口 / 结构体嵌套 / 构造函数 / 方法 / goroutine / channel
// 源码思路：类似 Go 标准库 net/http 的 Handler 流水线设计
//   每个 Handler 处理自己的一段逻辑，通过 channel 传递数据
//   就像流水线：接单 -> 厨房 -> 配送，各环节并发运行
// ============================================================

// ------ 状态常量 ------
// Go 没有 enum，用 const + iota 模拟
// iota 从 0 开始自动递增，这是 Go 源码里极常见的写法

// OrderStatus 订单状态枚举类型
// 使用 int 作为底层类型，方便序列化和比较
type OrderStatus int

// 订单状态常量定义
// iota 是 Go 的常量计数器，从 0 开始，每行使用 iota 时自动 +1
const (
	StatusPending    OrderStatus = iota // 0: 待接单 - 订单刚创建，等待商家确认
	StatusCooking                       // 1: 制作中 - 商家已接单，正在准备菜品
	StatusDelivering                    // 2: 配送中 - 菜品制作完成，骑手正在配送
	StatusDone                          // 3: 已完成 - 顾客已签收，订单结束
)

// String 将订单状态转换为可读的中文字符串
// 这是 Go 中为自定义类型添加 String() 方法的常用模式
// 类似于 Java 的 toString() 或 Python 的 __str__()
func (s OrderStatus) String() string {
	switch s {
	case StatusPending:
		return "待接单"
	case StatusCooking:
		return "制作中"
	case StatusDelivering:
		return "配送中"
	case StatusDone:
		return "已完成"
	default:
		return "未知状态"
	}
}

// ------ 接口定义 ------
// Processable 可处理的订单接口
// 接口定义了订单必须实现的方法，任何实现了这两个方法的类型都可以作为订单处理
// 这体现了 Go 的"隐式接口"特性，不需要显式声明 implements
type Processable interface {
	// Process 处理订单，推进订单状态并返回处理结果描述
	// 返回值描述了本次处理的详细情况，可用于日志或通知
	Process() string

	// GetID 获取订单的唯一标识符
	// 用于在并发处理时识别不同订单
	GetID() string
}

// ------ 结构体 ------

// MenuItem 菜品信息结构体
// 包含菜品的名称和价格
type MenuItem struct {
	Name  string  // 菜品名称，如"宫保鸡丁"
	Price float64 // 菜品价格，单位：元
}

// Order 订单基础结构体（核心数据结构）
// 这是所有订单类型的基础，包含通用的订单信息
// Order 被设计为可嵌入到其他结构体中，类似于面向对象中的基类
type Order struct {
	ID       string      // 订单唯一标识符，格式如 "F001"、"E001"
	Customer string      // 顾客姓名，用于联系和识别
	Items    []MenuItem  // 订单中的菜品列表，可包含多个菜品
	Status   OrderStatus // 当前订单状态，追踪订单进度
	Address  string      // 配送地址，精确到门牌号
}

// FoodOrder 普通外卖订单（通过结构体嵌套扩展 Order）
// 使用嵌入（匿名字段）的方式继承 Order 的所有字段和方法
// 类似于面向对象中的继承，但 Go 使用组合而非继承
type FoodOrder struct {
	Order                 // 嵌入基础订单，自动继承 ID, Customer, Items, Status, Address 字段
	RestaurantName string // 餐厅名称，如"老王烧烤"
	EstimatedMins  int    // 预计配送时间（分钟），用于向顾客展示送达时间预估
}

// ExpressOrder 急速/紧急订单（同样通过嵌入扩展 Order）
// 急速订单跳过了制作环节，直接进行配送
// 使用 Priority 字段区分紧急程度
type ExpressOrder struct {
	Order        // 嵌入基础订单
	Priority int // 优先级 1-5，数字越大越紧急，优先处理
}

// ------ 构造函数 ------
// Go 的构造函数是普通的函数，通常命名为 New[TypeName]
// 构造函数负责初始化结构体并返回指针或值

// NewOrder 创建基础订单实例
// 返回值是 Order 类型（值而非指针），因为 Order 会被嵌入到其他结构体中
// 当嵌入到其他结构体时，外层结构体会持有 Order 的副本
func NewOrder(id, customer, address string) Order {
	return Order{
		ID:       id,            // 订单号
		Customer: customer,      // 顾客姓名
		Address:  address,       // 配送地址
		Status:   StatusPending, // 初始状态为待接单
		Items:    []MenuItem{},  // 空菜品列表，稍后通过 AddItem 添加
	}
}

// NewFoodOrder 创建普通外卖订单实例
// 返回指针类型 *FoodOrder，因为需要调用者修改订单状态
func NewFoodOrder(id, customer, address, restaurant string) *FoodOrder {
	return &FoodOrder{
		Order:          NewOrder(id, customer, address), // 组合基础订单
		RestaurantName: restaurant,                      // 餐厅名称
		EstimatedMins:  30,                              // 默认预计配送时间 30 分钟
	}
}

// NewExpressOrder 创建急速订单实例
// 急速订单不需要餐厅名称，但需要指定优先级
func NewExpressOrder(id, customer, address string, priority int) *ExpressOrder {
	return &ExpressOrder{
		Order:    NewOrder(id, customer, address), // 组合基础订单
		Priority: priority,                        // 优先级 1-5
	}
}

// ------ 方法 ------
// Go 的方法是绑定到特定类型的函数
// 分为值接收者 (func (o Order)) 和指针接收者 (func (o *Order))
// 指针接收者可以修改原始数据，值接收者只能读取

// AddItem 向订单添加一个菜品
// 使用指针接收者 *Order，因为需要修改原始 Order 的 Items 字段
// 如果使用值接收者，修改不会影响原始数据
func (o *Order) AddItem(name string, price float64) {
	// append 会返回新的切片，需要重新赋值给 o.Items
	o.Items = append(o.Items, MenuItem{Name: name, Price: price})
}

// Total 计算订单的总价格
// 使用值接收者 Order，因为只读取数据不修改
// 对于简单的只读操作，值接收者是更好的选择（并发安全）
func (o Order) Total() float64 {
	var sum float64 // 使用 float64 存储总价，支持小数
	// 遍历所有菜品，累加价格
	for _, item := range o.Items {
		sum += item.Price
	}
	return sum
}

// Summary 返回订单的摘要信息
// 将订单的关键信息格式化为可读的字符串
// 包含订单号、顾客、地址、总价和当前状态
func (o Order) Summary() string {
	return fmt.Sprintf("订单[%s] 顾客：%s 地址：%s 金额：%.2f 状态:%s",
		o.ID, o.Customer, o.Address, o.Total(), o.Status)
}

// Process FoodOrder 的处理方法，实现 Processable 接口
// 模拟普通外卖的处理流程：接单 -> 制作 -> 配送 -> 完成
// 每次调用 Process() 会推进一个状态
func (f *FoodOrder) Process() string {
	switch f.Status {
	case StatusPending:
		// 待接单状态：商家确认接单，开始制作
		f.Status = StatusCooking
		itemName := ""
		if len(f.Items) > 0 {
			itemName = f.Items[0].Name
		}
		return fmt.Sprintf("[%s][%s]接单，%s开始制作，预计%d分钟",
			f.ID, f.RestaurantName, itemName, f.EstimatedMins)
	case StatusCooking:
		// 制作中状态：制作完成，等待骑手取餐
		f.Status = StatusDelivering
		return fmt.Sprintf("[%s] 制作完成，骑手取餐中", f.ID)
	case StatusDelivering:
		// 配送中状态：骑手已取餐，正在配送
		f.Status = StatusDone
		return fmt.Sprintf("[%s] 已送达 %s，请签收！", f.ID, f.Address)
	default:
		// 已完成或其他状态：订单已结束
		return fmt.Sprintf("[%s] 订单已完成", f.ID)
	}
}

// GetID 获取 FoodOrder 的订单号
// 实现 Processable 接口的 GetID() 方法
// 直接访问嵌入的 Order.ID 字段，因为嵌入字段可以像普通字段一样直接访问
func (f *FoodOrder) GetID() string {
	return f.ID
}

// Process ExpressOrder 的处理方法，实现 Processable 接口
// 急速订单的处理流程与普通订单不同：跳过制作环节，直接派单 -> 送达 -> 完成
func (e *ExpressOrder) Process() string {
	switch e.Status {
	case StatusPending:
		// 待接单状态：急速订单跳过商家制作，直接派给骑手
		e.Status = StatusDelivering
		return fmt.Sprintf("[%s] 【急速订单 P%d】直接派单给最近骑手！",
			e.ID, e.Priority)
	case StatusDelivering:
		// 配送中状态：急速送达
		e.Status = StatusDone
		return fmt.Sprintf("[%s] 紧急送达 %s!", e.ID, e.Address)
	default:
		// 已完成状态
		return fmt.Sprintf("[%s] 急速订单已完成", e.ID)
	}
}

// GetID 获取 ExpressOrder 的订单号
// 实现 Processable 接口的 GetID() 方法
func (e *ExpressOrder) GetID() string {
	return e.ID
}

// ------ goroutine + channel：并发处理订单 ------
// Go 的并发模型基于 CSP（通信顺序进程）
// 使用 goroutine 实现并发执行，使用 channel 实现 goroutine 之间的通信
//
// 流水线设计模式说明：
// 1. 每个订单在独立的 goroutine 中推进状态
// 2. 通过 channel 将处理结果发送给主 goroutine
// 3. 主 goroutine 收集所有结果并输出
//
// 对比源码中的常见模式：
//   go func() { result <- doWork() }()  // 启动 goroutine，结果放入 channel
//   for r := range result { ... }        // 主 goroutine 收集结果

// processOrderAsync 异步处理订单
// 启动独立的 goroutine 来处理订单的多个步骤
//
// 参数说明：
//   - oder: 实现 Processable 接口的订单实例
//   - step: 需要处理的步骤数（普通订单3步，急速订单2步）
//   - result: 用于接收处理结果的 channel
//
// channel 方向说明：
//
//	chan<- string 表示只能发送不能接收的 channel
//	<-chan string 表示只能接收不能发送的 channel
//	chan string 表示可以发送和接收
//
// 使用方向性 channel 可以防止误操作，提高代码安全性
func processOrderAsync(order Processable, step int, result chan<- string) {
	// 模拟订单处理，每个步骤之间有短暂延迟
	for i := 0; i < step; i++ {
		// 休眠 100 微秒，模拟处理耗时
		time.Sleep(time.Microsecond * 100)
		// 调用订单的 Process() 方法推进状态
		msg := order.Process()
		// 将处理结果发送到 channel，供主 goroutine 收集
		// 发送操作会阻塞，直到有 goroutine 接收
		result <- msg
	}
}

// main 主函数，程序入口
func main() {
	// ===== 创建订单实例 =====

	// 创建普通外卖订单 F001
	// 使用 NewFoodOrder 构造函数创建
	order1 := NewFoodOrder("F001", "小明", "幸福小区3栋", "老王烧烤")
	// 向订单添加菜品
	order1.AddItem("烤羊肉串", 25.0)
	order1.AddItem("烤玉米", 8.0)

	// 创建普通外卖订单 F002
	order2 := NewFoodOrder("F002", "小红", "阳光花园5栋", "兰州拉面")
	order2.AddItem("牛肉拉面", 18.0)

	// 创建急速订单 E001（用于紧急文件配送）
	// 优先级设为 5（最高）
	order3 := NewExpressOrder("E001", "张总", "CBD写字楼18F", 5)
	order3.AddItem("文件", 0)

	// ===== 输出订单摘要 =====
	fmt.Println("=== 外卖系统启动，并发处理订单 ===")
	fmt.Println(order1.Summary())
	fmt.Println(order2.Summary())
	fmt.Println(order3.Summary())
	fmt.Println()

	// ===== 并发处理订单 =====
	// 创建有缓冲的 channel，缓冲区大小为 10
	// 有缓冲 channel 不会阻塞发送方，直到缓冲区满
	// 这适合"多发一收"的场景（多个 goroutine 发送，一个 goroutine 接收）
	results := make(chan string, 10)

	// 启动三个 goroutine 并发处理订单
	// 使用 go 关键字启动 goroutine
	//
	// 处理步骤说明：
	// - 普通订单（F001, F002）：需要 3 步（接单 -> 制作 -> 配送）
	// - 急速订单（E001）：需要 2 步（派单 -> 送达），跳过了制作环节
	go processOrderAsync(order1, 3, results)
	go processOrderAsync(order2, 3, results)
	go processOrderAsync(order3, 2, results)

	// 总共需要接收的消息数：3 + 3 + 2 = 8
	total := 8

	// 阻塞等待并接收所有处理结果
	// 从 channel 接收数据，直到收够 total 条
	for i := 0; i < total; i++ {
		// <-results 是接收操作，会阻塞直到有数据可读
		msg := <-results
		fmt.Printf("[通知] %s\n", msg)
	}

	// ===== 输出最终状态 =====
	fmt.Println("\n=== 所有订单处理完毕 ===")
	fmt.Println(order1.Summary())
	fmt.Println(order2.Summary())
	fmt.Println(order3.Summary())
}
