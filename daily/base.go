package daily

import "beyondyyh/geek_algorithms/kit"

type (
	ListNode  = kit.ListNode
	DListNode = kit.DListNode

	Stack = kit.Stack
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
