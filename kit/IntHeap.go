package kit

// MaxHeap int数组实现大顶堆
type MaxHeap []int

// Len(), Less(), Swap() 实现 container/heap 的 sort 接口
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push(), Pop() 实现container/heap 的接口，不仅要改变元素值，还要改变数组长度，所以receiver用指针
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
