package offer

// 剑指 Offer 24. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// @lc: https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next // 暂存下一个节点
		cur.Next = pre   // 反向指向前面节点
		pre = cur        // pre后移
		cur = next       // cur后移
	}
	return pre
}
