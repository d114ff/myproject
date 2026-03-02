package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func getBigger(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// 下面是泛型,[T 类型](参数，T) 返回值T，T等于包含int32 | int64 两个返回值
func getBigger2[T int32 | int64](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// 接口泛型,类型写在接口里面
type Comparable interface {
	int32 | int64
}

// 函数泛型类型直接写接口
func getBigger3[T Comparable](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// 自定义泛型函数 Ages
type Ages int32
type Comparables interface {
	//int32 | int64 | Ages //要把自定义泛型写入接口
	~int32 | int64 //加一个～可以把int32所有类型都包含进来包括age
}

func getBigger4[T Comparables](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// 标准库里面的泛型,Ordered包括所有泛型
func getBigger5[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// 结构体使用泛型，方法不能使用方法
type Apples[T cmp.Ordered] struct{}

func (Apples[T]) getBigger6(a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

type GetUserRequest struct{}
type GetBookRequest struct{}

func httpRPC[T GetUserRequest | GetBookRequest](request T) {
	url := "http://127.0.0.1/"
	tp := reflect.TypeOf(request)
	switch tp.Name() {
	case "GetUserRequest":
		url += "user"
	case "GetBookRequest":
		url += "book"
	default:
		panic("不支持的request类型")
	}
	fmt.Println(url)
	bs, _ := json.Marshal(request)
	http.Post(url, "application/json", bytes.NewReader(bs))
}

func main() {
	c := getBigger(3, 6)
	fmt.Println(c)
	getBigger2[int32](8, 9) // 调用函数的时候指定一个类型（int32或者int64）
	getBigger3[rune](20, 9) // 接口参数类型，rune也是int32
	getBigger4[Ages](9, 10)
	getBigger5[int32](8, 9)
	a := Apples[int8]{}
	a.getBigger6(9, 22)
	httpRPC(GetBookRequest{})
}
