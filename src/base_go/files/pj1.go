package main

import (
	"fmt"
)

// 接口只描述"能做什么"，不关心是谁做、怎么做
// 就像图书馆规定：任何能借出的东西，必须有"借"和"还"两个动作
type Borrowable interface {
	Borrow(user string) string // 借出，返回提示信息
	Return() string            // 归还，返回提示信息
	Info() string              // 展示自己的信息
}

// 结构体只存数据，不放逻辑
// Book 代表一本书
type Book struct {
	Title      string // 书名
	Author     string // 作者
	ISBN       string // 书号（唯一标识）
	Available  bool   // 是否可借
	BorrowedBy string // 当前借阅人（空字符串 = 无人借阅）
}

// Magazine 代表一本杂志（和 Book 是不同类型，但都能借）

type Magazine struct {
	Title      string
	Issue      int // 期号，比如第 42 期
	Available  bool
	BorrowedBy string
}

// Go 没有 class 的 constructor，约定用 NewXxx 函数代替
// 返回指针（*Book）而不是值（Book），原因：
//   1. 调用方修改书的状态（比如 Available）时，改的是同一份数据
//   2. 避免每次传参都复制整个结构体

func NewBook(title, author, isbn string) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}
}

func NewMagazine(title string, isuue int) *Magazine {
	return &Magazine{
		Title:     title,
		Issue:     isuue,
		Available: true,
	}
}

// 方法 = 函数 + 接收者
// (b *Book) 叫"指针接收者"：方法内可以修改 b 的字段
// 如果写 (b Book) 叫"值接收者"：方法拿到的是副本，修改不影响原始数据
// Book 实现 Borrowable 接口的三个方法
func (b *Book) Borrow(user string) string {
	// 先检查是否可借
	if !b.Available {
		return fmt.Sprintf("《%s》已被 %s 借走，暂时不可借", b.Title, b.BorrowedBy)
	}
	b.Available = false
	b.BorrowedBy = user
	return fmt.Sprintf("%s 成功借出《%s》", user, b.Title)
}

func (b *Book) Return() string {
	if b.Available {
		return fmt.Sprintf("《%s》没有被借出", b.Title)
	}
	user := b.BorrowedBy
	b.Available = true
	b.BorrowedBy = ""
	return fmt.Sprintf(" %s 归还了 《%s》，感谢", user, b.Title)
}

func (b *Book) Info() string {
	status := "可借"
	if !b.Available {
		status = "已借出（借阅人：" + b.BorrowedBy + "）"
	}
	return fmt.Sprintf("书名：《%s》作者：%s ISBN：%s 状态：%s", b.Title, b.Author, b.ISBN, status)
}

// Magazine 也实现同样的接口
// 注意：Go 不需要写 "implements Borrowable"，只要方法签名匹配就自动满足
func (m *Magazine) Borrow(user string) string {
	if !m.Available {
		return fmt.Sprintf("《%s》第%d期已被借走", m.Title, m.Issue)
	}
	m.Available = false
	m.BorrowedBy = user
	return fmt.Sprintf("%s 借出《%s》第%d期", user, m.Title, m.Issue)
}

func (m *Magazine) Return() string {
	m.Available = true
	m.BorrowedBy = ""
	return fmt.Sprintf("《%s》第%d期已归还", m.Title, m.Issue)
}

func (m *Magazine) Info() string {
	return fmt.Sprintf("[杂志] 《%s》第%d期  可借:%v", m.Title, m.Issue, m.Available)
}

// 图书馆结构体 + 管理方法
type Library struct {
	Name string
	// 这里用接口切片！不管是 *Book 还是 *Magazine，都能放进来
	// 这就是"面向接口编程"，和 Go 标准库的设计思路一样
	Stock []Borrowable
}

func NewLibrary(name string) *Library {
	return &Library{
		Name:  name,
		Stock: []Borrowable{},
	}
}

// 添加馆藏（参数是接口，所以 Book 和 Magazine 都能传入）

func (lib *Library) Add(item Borrowable) {
	lib.Stock = append(lib.Stock, item)
}

// 添加馆藏（参数是接口，所以 Book 和 Magazine 都能传入）
func (lib *Library) ListAll() {
	fmt.Printf("\n=== %s 馆藏列表 ===\n", lib.Name)
	for i, item := range lib.Stock {
		// item.Info() 会根据实际类型（Book 或 Magazine）调用对应的方法
		// 这就是"多态"：同一行代码，不同类型执行不同逻辑
		fmt.Printf("%d %s \n", i+1, item.Info())
	}
}

func main() {
	lib := NewLibrary("市立图书馆")
	book1 := NewBook("Go 程序设计语言", "Donovan", "978-0")
	book2 := NewBook("算法导论", "Cormen", "978-1")
	mag1 := NewMagazine("程序员", 42)
	lib.Add(book1)
	lib.Add(book2)
	lib.Add(mag1)
	lib.ListAll()
	fmt.Println("\n--- 借阅操作 ---")
	fmt.Println(book1.Borrow("小明"))
	fmt.Println(book1.Borrow("小红"))
	fmt.Println(book1.Borrow("小李"))
	fmt.Println(book2.Borrow("小李"))
	fmt.Println(mag1.Borrow("小李"))
	lib.ListAll()

	fmt.Println("\n--- 归还操作 ---")
	fmt.Println(book1.Return())
	fmt.Println(book2.Return())
	fmt.Println(mag1.Return())

	lib.ListAll()

}
