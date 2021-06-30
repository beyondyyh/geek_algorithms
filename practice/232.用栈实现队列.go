package practice

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// 最简洁的代码
// @lc code=start
type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func NewMyQueue() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
// 出队时：
// 1.如果out为空，则一次性将in里的元素转移到out中
// 2.如果out不为空，则正常Pop
func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.shift()
	}
	top := q.outStack[len(q.outStack)-1]        // outStack的栈顶元素，也就是队列的队首
	q.outStack = q.outStack[:len(q.outStack)-1] // outStack.Pop
	return top
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.shift()
	}
	return q.outStack[len(q.outStack)-1] // outStack的栈顶元素，也就是队列的队首
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// 一次性把inStack里的元素 全部转移 到outStack中
func (q *MyQueue) shift() {
	for len(q.inStack) != 0 {
		top := q.inStack[len(q.inStack)-1]       // 栈顶元素
		q.inStack = q.inStack[:len(q.inStack)-1] // inStack.Pop
		q.outStack = append(q.outStack, top)     // outStack.Push
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
