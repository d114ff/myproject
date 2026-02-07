package main

type Selector interface {
	Select([]string) int
}

type ConnectionPool struct {
	Servers      []string
	LoadBalancer Selector // 成员变量是接口函数类型
}

func f1([]string) int {
	return 0
}

func f2([]string) int {
	return 0
}

type RoundRobin struct{}

func (RoundRobin) Select(s []string) int {
	return f1(s)
}

type interleave struct{}

func (interleave) Select(s []string) int {
	return f2(s)

}

type ConnectionPool2 struct {
	Servers      []string
	LoadBalancer func([]string) int
}

func main27() {
	cp := ConnectionPool{
		Servers:      []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		LoadBalancer: RoundRobin{},
	}
	_ = cp.LoadBalancer.Select(cp.Servers)

	cp2 := ConnectionPool2{
		Servers: []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		//LoadBalancer: f1,
		LoadBalancer: f2,
	}
	_ = cp2.LoadBalancer(cp2.Servers)
}
