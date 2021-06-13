package kit

// 队列 queue---------------------------------------------------------
type Queue []interface{}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
