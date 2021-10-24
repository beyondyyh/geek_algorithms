package offerII

type ListNode struct {
	Val  int
	Next *ListNode
}

type DListNode struct {
	Val        int
	Next, Prev *ListNode
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
