package offer

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
