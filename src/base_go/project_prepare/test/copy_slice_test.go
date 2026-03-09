package projectprepare_test

import (
	projectprepare "myproject/src/base_go/project_prepare/test"
	"testing"
)

func TestCopySlice(t *testing.T) {
	//sonic.Marshal(2) // sonic是sonic目录下api的包名，重名是巧合，根目录下函数可以之间调用
	// go get github.com/bytedance/sonic 把文件源代码拉到本地，放在gonpasth/pkg/mod 目录下
	var src, dest []int16
	src = []int16{1, 2, 3, 4}
	var c, n int

	c = len(src) - 1
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != c {
		t.Errorf("c=%d,n=%d", c, n)
	}
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Errorf("c=%d,i=%d,dest%d src%d", c, i, dest[i], src[i])
		}

	}
	c = len(src)
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != c {
		t.Errorf("c=%d,n=%d", c, n)
	}
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Errorf("c=%d,i=%d,dest%d src%d", c, i, dest[i], src[i])
		}

	}

	c = len(src) + 1 - 1
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != c {
		t.Errorf("c=%d,n=%d", c, n)
	}
	//dest[0]--
	//dest[1]--
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Fatalf("c=%d,i=%d,dest%d src%d", c, i, dest[i], src[i]) //Fatalf 发现一点问题就推出for循环和函数
		}

	}
}
func BenchmarkCopySlice(b *testing.B) {
	src := make([]int8, 10000)
	dest := make([]int8, 10000)
	b.ResetTimer() //开始记时
	for i := 0; i < b.N; i++ {
		projectprepare.CopySlice(dest, src)
	}
}

func BenchmarkStdCopySlice(b *testing.B) {
	src := make([]int8, 10000)
	dest := make([]int8, 10000)
	b.ResetTimer() //开始记时
	for i := 0; i < b.N; i++ {
		copy(dest, src)
	}
}
