package week01

// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// @leetcode: https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// 方法一：递归
// 递归终止条件：head 为空指针或者 next 为空指针，也就是当前无节点或者只有一个节点，无法进行交换
// 递归子逻辑：设需要交换的两个点为 head 和 next，head 连接后面交换完成的子链表，next 连接 head，完成交换
func swapPairs1(head *ListNode) *ListNode {
	// terminator
	if head == nil || head.Next == nil {
		return head
	}
	// head连接后面交换完成的子链表，next连接head
	next := head.Next
	head.Next = swapPairs1(next.Next)
	next.Next = head
	return next
}

// 方法二：迭代
func swapPairs2(head *ListNode) *ListNode {
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
