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
