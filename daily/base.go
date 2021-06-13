package daily

import "beyondyyh/geek_algorithms/kit"

type (
	ListNode  = kit.ListNode
	DListNode = kit.DListNode

	Queue = kit.Queue
	Stack = kit.Stack
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 堆，优先队列----------------------------------------------------------
type kvpair struct {
	key, val int
}

type MinHeap []kvpair

// Implements sort Len,Less,Swap interface
func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].key > h[j].key
	}
	return h[i].val > h[j].val
}

// Implements container/Heap Push,Pop interface, receiver is ptr
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(kvpair))
}

func (h *MinHeap) Pop() interface{} {
	n := (*h).Len()
	x := (*h)[n-1] // 最后一个元素
	*h = (*h)[:n-1]
	return x
}

// 查看堆顶元素，数组第一个元素即是堆顶
func (h MinHeap) Peek() kvpair {
	return h[0]
}
