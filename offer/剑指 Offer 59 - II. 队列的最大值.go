package offer

import "container/list"

// 剑指 Offer 59 - II. 队列的最大值
// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
// 若队列为空，pop_front 和 max_value 需要返回 -1

// 示例 1：
// 输入:
// ["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
// [[],[1],[2],[],[],[]]
// 输出: [null,null,null,2,1,2]

// 示例 2：
// 输入:
// ["MaxQueue","pop_front","max_value"]
// [[],[],[]]
// 输出: [null,-1,-1]
// @lc: https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

// 双端队列，维护2个队列，queue和辅助队列deque（单调递减队列）
// 1. push_pack时，queue正常入队，若入队的元素比deque最后一个元素大，则把最后一个元素移除（可以理解为：新加入的元素比较强大，把前面比他弱的都干掉了）；否则：直接入队；
// 2. pop_front()时，queue正常出队，若出队的元素是最大元素，则deque需要同时将首元素出队，以保持队列和双向队列的元素一致性。
// 3. max_value, deque的队首元素就是最大值。
// 时间复杂度：O(1)
// 空间复杂度：O(n) 当所有元素单调递减时，deque中保存了所有元素
type MaxQueue struct {
	queue *list.List
	deque *list.List
}

// func Constructor() MaxQueue {
func NewMaxQueue() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		deque: list.New(),
	}
}

func (q *MaxQueue) Max_value() int {
	if q.deque.Len() == 0 {
		return -1
	}
	return q.deque.Front().Value.(int)
}

func (q *MaxQueue) Push_back(value int) {
	q.queue.PushBack(value)
	for q.deque.Len() > 0 && q.deque.Back().Value.(int) < value {
		q.deque.Remove(q.deque.Back())
	}
	q.deque.PushBack(value)
}

func (q *MaxQueue) Pop_front() int {
	if q.queue.Len() == 0 {
		return -1
	}
	if q.queue.Front().Value.(int) == q.deque.Front().Value.(int) {
		q.deque.Remove(q.deque.Front())
	}
	return q.queue.Remove(q.queue.Front()).(int)
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
