package week01

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
// @leetcode: https://leetcode-cn.com/problems/min-stack

// 最简单的用2个栈，1个栈正常的先入后出，1个辅助栈用来存当前最小值
type MinStack struct {
	A *Stack // 正常的数据栈
	B *Stack // 辅助栈用来存当前最小值
}

func NewMinStack() *MinStack {
	return &MinStack{
		A: &Stack{},
		B: &Stack{},
	}
}

// push入栈，A正常入栈，B栈为空（第一个元素）或x比B栈的栈顶元素还小，说明x为此时的最小值，则入B栈
func (s *MinStack) push(x int) {
	s.A.Push(x)
	if s.B.IsEmpty() || x < s.B.Top().(int) {
		s.B.Push(x)
	}
}

// pop出栈，A正常出栈，如果A出栈的元素等于B栈顶的元素，则说明此时最小值要变了，则B也出栈
func (s *MinStack) pop() {
	if s.A.Pop().(int) == s.B.Top().(int) {
		s.B.Pop()
	}
}

// top返回A栈顶元素
func (s *MinStack) top() int {
	return s.A.Top().(int)
}

func (s *MinStack) getMin() int {
	return s.B.Top().(int)
}
