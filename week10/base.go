package week10

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

// ----------------------------------------------------------
type Stack []interface{}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) Pop() interface{} {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}

// ----------------------------------------------------------
type Queue []interface{}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *Queue) Top() interface{} {
	return (*q)[0]
}

// ----------------------------------------------------------
type kvpair struct {
	key, val int
}

type MinHeap []kvpair

// Implements sort Len,Less,Swap interface
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
