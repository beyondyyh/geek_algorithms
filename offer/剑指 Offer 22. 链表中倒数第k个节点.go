package offer

// 剑指 Offer 22. 链表中倒数第k个节点
// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

// 示例：
// 给定一个链表: 1->2->3->4->5, 和 k = 2.
// 返回链表 4->5.
// @lc: https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

// 1. 初始化快慢指针指向头head​；
// 2. 快指针先向前走k步，结束后，双指针fast、slow 间距离为k；
// 3. 双指针共同移动，每次走一步，直到fast走到尾部时跳出（跳出后，slow与尾节点距离为k-1，即slow指向倒数第k个节点）。
// 返回slow即可
func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for ; k > 0 && fast != nil; k-- {
		fast = fast.Next
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
