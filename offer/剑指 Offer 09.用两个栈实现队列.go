package offer

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
// 示例 1：
// 输入：
// ["CQueue","appendTail","deleteHead","deleteHead"]
// [[],[3],[],[]]
// 输出：[null,null,3,-1]
// @lc: https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

// 复杂度分析：
// 针对每一个元素都入栈和出栈一次，所以时间复杂度为：O(1)
// 空间复杂度：O(n) 栈中有n个元素
type CQueue struct {
	inStack  []int
	outStack []int
}

// func Constructor() CQueue {
func NewCQueue() CQueue {
	return CQueue{
		inStack:  []int{}, // 入栈
		outStack: []int{}, // 出栈
	}
}

// inStack 将value正常入栈
func (q *CQueue) AppendTail(value int) {
	q.inStack = append(q.inStack, value) // Push
}

// 如果outStack非空，直接Pop即可
// 如果outStack为空，则将in一次性转移到out中，然后再Pop
func (q *CQueue) DeleteHead() int {
	if len(q.outStack) == 0 {
		q.shift()
	}

	x := -1
	if len(q.outStack) > 0 {
		x = q.outStack[len(q.outStack)-1]           // Top
		q.outStack = q.outStack[:len(q.outStack)-1] // Pop
	}
	return x
}

// 一次性将inStack中的元素转移到outStack中
func (q *CQueue) shift() {
	for len(q.inStack) > 0 {
		x := q.inStack[len(q.inStack)-1]         // Top
		q.inStack = q.inStack[:len(q.inStack)-1] // Pop
		q.outStack = append(q.outStack, x)       // Push
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
