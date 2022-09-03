package linkedList

// 快慢指针：
// 快慢指针起始位置都在head，fast先走k步骤，此时slow和fast之间的距离为k，fast到达第k+1个节点。
// 然后快慢指针同时走，当fast走到最后一个节点的下一个节点时，slow指针所在的节点就是倒数第k个
// tips：链表的问题可以在纸上画一画
func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	// fast先走k步，slow和fast之间的距离为k，fast此时在第k+1个节点
	for ; k > 0 && fast != nil; k-- {
		fast = fast.Next
	}
	// fast在前，因此判断fast!=nil就行了
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
