package week01

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// @leetcode: https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// 方法一：
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1} // 哑结点，链表中通用解决方案
	dummy.Next = head
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		start, end := curr.Next, curr.Next.Next
		curr.Next = end
		start.Next = end.Next
		end.Next = start
		curr = start
	}
	return dummy.Next
}
