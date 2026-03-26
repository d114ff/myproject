package main

import (
	"fmt"
	"math/rand"
)

// ============================================================
// 项目二：RPG 游戏角色战斗系统
// 涉及：嵌套结构体 / 方法集 / 接口组合 / 构造函数链
// 源码思路：类似 Go 中 http.Handler 接口的分层设计
//
//	http.Handler 只定义 ServeHTTP，各种 Handler 自己实现细节
//	这里 Fighter 接口让不同角色用同一套战斗框架
//
// ============================================================
type Describer interface {
	Descibe() string //自我介绍
}

type Fighter interface {
	Describer            //嵌入 Describer 接口
	Attack() int         // 计算攻击力
	TakeDameage(dmg int) //承受伤害
	IsAlive() bool       //是否存活
}

// ------ 基础结构体（会被其他结构体"嵌入"）------
// 这是 Go 实现"继承"效果的方式，叫"结构体嵌入"
// 嵌入后，外层结构体直接拥有内层的字段和方法

type BaseStast struct {
	HP    int // 当前血量
	MaxHP int // 最大血量
	ATK   int // 基础攻击力
	DEF   int // 防御力
}

func (b *BaseStast) TakeDameage(dmg int) {
	// 先减去防御力，最少造成 1 点伤害
	actual := dmg - b.DEF
	if actual < 1 {
		actual = 1
	}
	b.HP -= actual
	if b.HP < 0 {
		b.HP = 0 // 血量不能为负
	}
}

func (b *BaseStast) IsAlive() bool {
	return b.HP > 0
}

func (b *BaseStast) HPBar() string {
	pct := b.HP * 20 / b.MaxHP // 把血量转换成 20 格
	bar := ""
	for i := 0; i < 20; i++ {
		if i < pct {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return fmt.Sprintf("[%s]%d/%d", bar, b.HP, b.MaxHP)

}

// ------ 具体角色结构体 ------
// Warrior 战士：高防御，攻击稳定

type Warrior struct {
	Name        string
	BaseStast       //嵌入 BaseStats，直接继承 TakeDamage / IsAlive / HPBar
	ShieldBlock int // 额外盾牌格挡值
}

// Knight 骑士：均衡型，有被动加成
type Kinght struct {
	Name string
	BaseStast
	MountBonus int // 骑乘加成（攻击力额外加成）
}

// Mage 法师：高攻击，低防御
type Mage struct {
	Name string
	BaseStast
	Mana    int // 法力值，影响攻击
	MaxMana int
}

// ------ 构造函数（工厂模式）------
// 这种"根据类型创建对象"的模式在 Go 源码里非常常见
// 比如 net/http 包的 NewServeMux()、NewRequest() 等

func NewWarrior(name string) *Warrior {
	return &Warrior{
		Name: name,
		BaseStast: BaseStast{
			HP:    120, // 当前血量
			MaxHP: 120, // 最大血量
			ATK:   18,  // 基础攻击力
			DEF:   10,  // 防御力
		},
		ShieldBlock: 5, // 额外盾牌格挡值
	}
}
func NewKnight(name string) *Kinght {
	return &Kinght{
		Name: name,
		BaseStast: BaseStast{
			HP:    100,
			MaxHP: 100,
			ATK:   22,
			DEF:   7,
		},
		MountBonus: 5, // 骑乘加成（攻击力额外加成）
	}
}

func NewMage(name string) *Mage {
	return &Mage{
		Name: name,
		BaseStast: BaseStast{
			HP:    70,
			MaxHP: 70,
			ATK:   35,
			DEF:   2, // 法师防御低
		},
		Mana:    100,
		MaxMana: 100,
	}
}

// ------ 方法：实现 Fighter 接口 ------
// 每个角色用不同方式计算攻击，体现多态
func (w *Warrior) Attack() int {
	// 战士：基础攻击 + 随机浮动 ± 3
	return w.ATK + rand.Intn(7) - 3
}

func (w *Warrior) TakeDamage(dmg int) {
	// 战士的 TakeDamage 覆盖了 BaseStats 的同名方法
	// 额外加上盾牌格挡
	actual := dmg - w.DEF - w.ShieldBlock
	if actual < 1 {
		actual = 1
	}
	w.HP -= actual
	if w.HP < 0 {
		w.HP = 0
	}
}

func (w *Warrior) Descibe() string {
	return fmt.Sprintf("⚔  战士【%s】HP:%s ATK:%d DEF:%d+%d(盾)", w.Name, w.HPBar(), w.ATK, w.DEF, w.ShieldBlock)
}

func (k *Kinght) Attack() int {
	// 骑士：基础攻击 + 骑乘加成 + 随机浮动
	return k.ATK + k.MountBonus + rand.Intn(5) - 2
}

func (k *Kinght) Descibe() string {
	return fmt.Sprintf("🐴 骑士【%s】HP:%s ATK:%d+%d(骑乘)", k.Name, k.HPBar(), k.ATK, k.MountBonus)
}

func (m *Mage) Attack() int {
	// 法师：法力越高，攻击越强
	// 这里演示方法内可以修改结构体状态（消耗法力）
	if m.Mana >= 20 {
		m.Mana -= 20 // 每次攻击消耗 20 法力
		return m.ATK + m.Mana/5 + rand.Intn(10)
	}
	// 法力不足时用普通攻击
	return m.ATK/2 + rand.Intn(5)
}
func (m *Mage) Descibe() string {
	return fmt.Sprintf("✨ 法师【%s】HP:%s ATK:%d Mana:%d/%d",
		m.Name, m.HPBar(), m.ATK, m.Mana, m.MaxMana)
}

// ------ 普通函数：接受接口参数，实现多态调用 ------
// 这个函数不知道、也不关心 a 和 b 具体是什么类型
// 只要它们满足 Fighter 接口，就能参与战斗
// 这和 Go 标准库 sort.Sort(data Interface) 的思路完全一样

func Battle(a, b Fighter) {
	fmt.Println("\n========== 战斗开始 ==========")
	fmt.Println(a.Descibe())
	fmt.Println(b.Descibe())

	round := 1
	for a.IsAlive() && b.IsAlive() {
		fmt.Printf("\n--- 第 %d 回合 ---\n", round)
		// a 攻击 b
		dmg := a.Attack()
		b.TakeDameage(dmg)
		fmt.Printf("%s 攻击 -> %d 伤害\n", a.Descibe(), dmg)

		if !b.IsAlive() {
			break
		}
		// b 反击 a
		dmg = b.Attack()
		a.TakeDameage(dmg)
		fmt.Printf("%s 反击 -> %d 伤害\n", b.Descibe(), dmg)

		round++
		if round > 20 {
			fmt.Println("战斗超时，平局！")
			return
		}
	}
	fmt.Println("\n========== 战斗结束 ==========")
	if a.IsAlive() {
		fmt.Printf("胜者：%s\n", a.Descibe())
	} else {
		fmt.Printf("胜者：%s\n", b.Descibe())
	}

}

func main2() {
	warrior := NewWarrior("铁甲虎")
	mage := NewMage("闪电法师")
	// Battle 接收 Fighter 接口，两种类型都能传入
	Battle(warrior, mage)
	/*
		fmt.Println("\n--- 再来一场：骑士 vs 战士 ---")
		knight := NewKnight("银月骑士")
		warrior2 := NewWarrior("钢铁龙")
		Battle(knight, warrior2)
	*/

}
