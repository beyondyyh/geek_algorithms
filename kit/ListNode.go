package kit

// ListNode 单链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// Ints2List convert []int to singly-linked-list.
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{Val: -1} // 哑结点，链表里通用做法
	curr := dummy
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return dummy.Next
}

// List2Ints convert singly-linked-list to []int，假设链表无环
func List2Ints(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
