package kit

// 基于数组实现栈

type Stack []interface{}

// Len 返回栈长度
func (s *Stack) Len() int {
	return len(*s)
}

// Push 入栈，数组头部作为栈底
func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

// Pop 删除并返回栈顶元素
func (s *Stack) Pop() interface{} {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}

// Top 返回栈顶元素不改变结构
func (s *Stack) Top() interface{} {
	return (*s)[s.Len()-1]
}

// IsEmpty 返回s是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
