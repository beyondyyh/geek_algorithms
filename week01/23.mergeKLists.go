package week01

import "container/heap"

// 23. 合并K个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
// @lc: https://leetcode-cn.com/problems/merge-k-sorted-lists/

// 方法一：小顶堆（优先队列），维护长度为k的小根堆，将每个链表的头结点入堆，弹出堆顶并把堆顶链表的下一个节点再次入堆
// 时间复杂度：O(N log(k)) n为链表个数，遍历n次，每层遍历中堆的调整时间复杂度为：log(k)
// 空间复杂度：O(n)  堆的空间
func mergeKLists(lists []*ListNode) *ListNode {
	// 遍历链表个数，将每个链表的头结点入小根堆，此时堆顶元素最小
	h := &IHeap{}
	heap.Init(h)
	for _, ln := range lists {
		if ln != nil {
			heap.Push(h, ln)
		}
	}

	// 哑结点方便操作，cur总是在后移，哑结点没动
	dummy := &ListNode{Val: -1}
	cur := dummy
	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*ListNode) // 堆顶元素最小
		cur = cur.Next                     // 自身指针后移
		if cur.Next != nil {               // 堆顶节点的下一个节点不为空就继续入堆
			heap.Push(h, cur.Next)
		}
	}
	return dummy.Next
}

// ----------------------------------------
// 小顶堆实现
type IHeap []*ListNode

// 实现 "container/Heap" 接口
func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push, Pop 操作不仅改变堆元素，而且改变堆长度，因此 receiver 为指针
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Pop返回数组最后一个元素，并把长度减一，heap接口通过heapifyDown()自上而下调整，维持堆的结构
func (h *IHeap) Pop() interface{} {
	old := *h
	x := old[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}
