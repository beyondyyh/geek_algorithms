package kit

import "sort"

// MinHeap int数组实现大顶堆
type MaxHeap []int

// Len(), Less(), Swap() 实现 container/heap 的 sort 接口
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push(), Pop() 实现container/heap 的接口，不仅要改变元素值，还要改变数组长度，所以receiver用指针
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop, 删除堆顶元素
// 注意：此处的x并不是真正的堆顶元素，而是数组的最后一个元素用来替换堆顶的，Pop只是实现了标准库 `container/heap`的Pop接口，内部还有调整堆结构的一系列操作
// 想一想二叉堆的删除操作：
//   - 将堆尾部元素替换到堆顶，即数组最后一个元素替换第一个元素
//   - 然后依次从堆顶向下调整整个堆的结构，（一直到堆尾即可），`heapifyDown`
// go标准库源码：https://golang.org/src/container/heap/heap.go?s=2190:2223#L50
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek 只查看堆顶元素，不改变堆的结构
func (h *MaxHeap) Peek() interface{} {
	x := (*h)[h.Len()-1]
	return x
}

// MinHeap int数组实现小顶堆
type MinHeap []int

// Len(), Less(), Swap() 实现 container/heap 的 sort 接口
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push(), Pop() 实现container/heap 的接口，不仅要改变元素值，还要改变数组长度，所以receiver用指针
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop, 删除堆顶元素
// 注意：此处的x并不是真正的堆顶元素，而是数组的最后一个元素用来替换堆顶的，Pop只是实现了标准库 `container/heap`的Pop接口，内部还有调整堆结构的一系列操作
// 想一想二叉堆的删除操作：
//   - 将堆尾部元素替换到堆顶，即数组最后一个元素替换第一个元素
//   - 然后依次从堆顶向下调整整个堆的结构，（一直到堆尾即可），`heapifyDown`
// go标准库源码：https://golang.org/src/container/heap/heap.go?s=2190:2223#L50
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek 只查看堆顶元素，不改变堆的结构
func (h *MinHeap) Peek() interface{} {
	x := (*h)[h.Len()-1]
	return x
}

// ------------------------------------------------------------------------------------------------------------
// 使用标准库的 sort.IntSlice 实现堆数据结构（优先队列），标准库已经实现了排序相关的 `Len,Less,Swap` 接口，不过默认是升序也就是小顶堆
// 因此只需要实现 container/heap 接口的 `Push,Pop` 方法即可
type IMinHeap struct{ sort.IntSlice }

// 注意：Push、Pop的receiver需要是指针
func (h *IMinHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IMinHeap) Pop() interface{} {
	old := h.IntSlice
	n := h.IntSlice.Len()
	x := old[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

// 大顶堆
type IMaxHeap struct{ sort.IntSlice }

// 默认是小顶堆，重写 Less 方法
func (h IMaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// 注意：Push、Pop的receiver需要是指针
func (h *IMaxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IMaxHeap) Pop() interface{} {
	old := h.IntSlice
	n := h.IntSlice.Len()
	x := old[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}
