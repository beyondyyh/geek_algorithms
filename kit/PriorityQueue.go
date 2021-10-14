package kit

// type Entry interface{}

// type LessFunc func(p, q Entry) bool

// type priorityQueueImpl struct {
// 	data []Entry
// 	less LessFunc
// }

// type priorityQueue struct {
// 	priorityQueueImpl
// }

// // Len 实现 sort.Interface 的Len接口方法
// func (pq priorityQueueImpl) Len() int { return len(pq.data) }

// // Swap 实现 sort.Interface 的Swap接口方法
// func (pq priorityQueueImpl) Swap(i, j int) {
// 	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
// }

// // Less 实现 sort.Interface 的Less接口方法
// func (pq priorityQueueImpl) Less(i, j int) bool {
// 	p, q := pq.data[i], pq.data[j]
// 	return pq.less(p, q)
// }

// // Push 实现 Push 接口方法，receiver是指针
// func (pq *priorityQueueImpl) Push(x interface{}) {
// 	(*pq).data = append((*pq).data, x.(Entry))
// }

// // Pop 实现 Pop 接口方法，receiver是指针
// func (pq *priorityQueueImpl) Pop() interface{} {
// 	old := (*pq)
// 	n := old.Len()
// 	x := old.data[n-1]
// 	(*pq).data = (*pq).data[:n-1]
// 	return x
// }

// // NewPriorityQueue 创建一个优先队列，传入一个比较器
// func NewPriorityQueue(comparator LessFunc) *priorityQueue {
// 	pq := priorityQueue{
// 		priorityQueueImpl: priorityQueueImpl{
// 			data: make([]Entry, 0),
// 			less: comparator,
// 		},
// 	}
// 	heap.Init(&pq.priorityQueueImpl)
// 	return &pq
// }

// func (pq *priorityQueue) Push(x Entry) {
// 	heap.Push(&pq.priorityQueueImpl, x)
// }

// func (pq *priorityQueue) Pop() Entry {
// 	return heap.Pop(&pq.priorityQueueImpl).(Entry)
// }

// func (pq *priorityQueue) Peek() Entry {
// 	return pq.priorityQueueImpl.data[0]
// }

// func (pq *priorityQueue) Len() int {
// 	return pq.priorityQueueImpl.Len()
// }

// 基于heap接口实现一个二维数组的优先队列，大顶堆

type Item struct {
	Index int // 元素下标
	Value int // 元素值
}

type PriorityQueue []*Item

// 实现heap接口内嵌的sort接口的Len, Swap, Less方法
func (pq PriorityQueue) Len() int      { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Less 比较器，先比较Value，Value相等时比较Index
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Value == pq[j].Value {
		return pq[i].Index > pq[j].Index
	}
	return pq[i].Value > pq[j].Value
}

// Push, Pop方法receiver用指针，因为不仅要改变数组元素，还改变长度
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop 此处pop的并不是真正的堆顶元素，而是数组的最后一个元素用来替换堆顶的，Pop只是实现了标准库 `container/heap`的Pop接口，内部还有调整堆结构的一系列操作
// 参考堆的删除操作，尾部先替换头部，然后自上而下调整堆结构
// go源码：https://golang.org/src/container/heap/heap.go?s=2190:2223#L50
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := pq.Len()
	item := old[n-1]
	old[n-1] = nil // 避免内存泄露
	*pq = old[0 : n-1]
	return item
}

// 只查看堆顶元素，不更改堆的结构，数组头部元素即为堆顶
func (pq PriorityQueue) Peek() *Item {
	return pq[0]
}

// 只查看堆顶元素，不更改堆的结构，数组头部元素即为堆顶
func (pq PriorityQueue) Elements() []Item {
	items := make([]Item, 0, pq.Len())
	for _, item := range pq {
		items = append(items, *item)
	}
	return items
}
