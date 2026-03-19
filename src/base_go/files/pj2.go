package main

import "fmt"

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
	MaxHP int // 当前血量
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
