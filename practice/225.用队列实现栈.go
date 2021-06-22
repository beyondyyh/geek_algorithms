/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

/** Push element x onto stack. */
// 入栈时：把队列想象成一个环，首先获得入栈前的元素个数 n，然后将元素入队到队列，
// 再将队列中的前 n 个元素（即除了新入栈的元素之外的全部元素）依次出队并入队到对队尾
func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- { // 注意：n>0，不包括刚入队的元素x
		s.queue = append(s.queue, s.queue[0]) // 将队首元素依次加入队尾
		s.queue = s.queue[1:]                 // 删除队首，即Pop
	}
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	x := s.queue[0]
	s.queue = s.queue[1:]
	return x
}

/** Get the top element. */
func (s *MyStack) Top() int {
	return s.queue[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
