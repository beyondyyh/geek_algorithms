package week01

// 由于不需要查询某个节点，只需要操作头尾，因此基于链表是最佳选择
// 而且链表的效率更高

// NewDListNode returns *DListNode
func NewDListNode(x int) *DListNode {
	return &DListNode{Val: x}
}

// MyCircularDeque MyCircularDeque based on double-linked-list structure delcare
type MyCircularDeque struct {
	len, cap   int        // len 链表实际长度，cap 开辟的空间长度
	head, tail *DListNode // head,tail 其实是哑结点
}

// NewMyCircularDeque struct initialize, Set the size of the deque to be k.
func NewMyCircularDeque(k int) *MyCircularDeque {
	head, tail := NewDListNode(-1), NewDListNode(-1)
	head.Next = tail
	tail.Prev = head
	return &MyCircularDeque{
		cap:  k,
		head: head,
		tail: tail,
	}
}

// InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) InsertFront(x int) bool {
	if q.IsFull() {
		return false
	}

	node := NewDListNode(x)
	node.Prev = q.head
	node.Next = q.head.Next
	node.Next.Prev = node
	node.Prev.Next = node
	q.len++
	return true
}

// InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) InsertLast(x int) bool {
	if q.IsFull() {
		return false
	}

	node := NewDListNode(x)
	node.Next = q.tail
	node.Prev = q.tail.Prev
	node.Prev.Next = node
	node.Next.Prev = node
	q.len++
	return true
}

// DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}

	// 待移除节点
	node := q.head.Next
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	q.len--
	return true
}

// DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}

	// 待移除节点
	node := q.tail.Prev
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	q.len--
	return true
}

// GetFront Get the front item from the deque.
func (q *MyCircularDeque) GetFront() int {
	return q.head.Next.Val
}

// GetRear Get the last item from the deque.
func (q *MyCircularDeque) GetRear() int {
	return q.tail.Prev.Val
}

// IsEmpty Checks whether the circular deque is empty or not.
func (q *MyCircularDeque) IsEmpty() bool {
	return q.len == 0
}

// IsFull Checks whether the circular deque is full or not.
func (q *MyCircularDeque) IsFull() bool {
	return q.len == q.cap
}
