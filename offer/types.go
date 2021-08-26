package offer

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

type DListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// int类型小顶堆
type iMinHeap struct{ sort.IntSlice }

// 实现 container/Heap 接口的 Push,Pop 方法
func (h *iMinHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *iMinHeap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

// 查看堆顶元素，不改变结构
func (h iMinHeap) Peek() int {
	return h.IntSlice[0]
}

// int类型大顶堆，sort.IntSlice 默认升序，即小顶堆
type iMaxHeap struct{ sort.IntSlice }

// 实现 container/Heap 接口的 Push,Pop 方法
func (h *iMaxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *iMaxHeap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

func (h iMaxHeap) Peek() int {
	return h.IntSlice[0]
}
