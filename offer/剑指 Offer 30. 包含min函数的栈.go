package offer

import "math"

// 剑指 Offer 30. 包含min函数的栈
// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
// 示例:
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.min();   --> 返回 -3.
// minStack.pop();
// minStack.top();   --> 返回 0.
// minStack.min();   --> 返回 -2.
// @lc: https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

type MinStack struct {
	stack  []int // 正常出入元素的栈
	helper []int // 保存最小元素的辅助栈
}

/** initialize your data structure here. */
// func Constructor() MinStack {
func NewMinStack() MinStack {
	return MinStack{
		stack:  []int{},
		helper: []int{math.MaxInt32}, // 便于比较
	}
}

// 1.stack 正常将x入栈
// 2.helper 比较一下x和栈顶元素的大小，将二者较小者入栈
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	top := s.helper[len(s.helper)-1]
	s.helper = append(s.helper, min(top, x))
}

// stack, helper正常出栈
func (s *MinStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
		s.helper = s.helper[:len(s.helper)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	return s.helper[len(s.helper)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
