package v24

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](n int) Set[T] {
	m := make(map[T]struct{}, n)
	return Set[T](m)
}

// 往Set里面添加元素
func (set Set[T]) Add(ele T) {
	set[ele] = struct{}{} //struct{}是类型，{}是空字段
}

func (set Set[T]) Len() int {
	return len(set)
}

func (set Set[T]) Remove(ele T) {
	delete(set, ele)
}

func (set Set[T]) Exists(ele T) bool {
	_, exists := set[ele]
	return exists
}

func (set Set[T]) Range(f func(ele T)) {
	for key := range set {
		f(key)
	}
}
