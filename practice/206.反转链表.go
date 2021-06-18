package practice

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一：迭代法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseList1(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		nextTmp := cur.Next // 保存下一个节点
		cur.Next = pre      // 当前节点反向指向前一个
		pre = cur           // pre节点的指针后移一位
		cur = nextTmp       // cur节点的指针后移一位
	}
	return pre
}

// 方法二：递归法
// 时间复杂度；O(n)
// 空间复杂度：O(n)
func reverseList(head *ListNode) *ListNode {
	// terminator
	if head == nil || head.Next == nil {
		return head
	}
	// drill down
	cur := reverseList(head.Next) // 5
	head.Next.Next = head         // 第一次遍历：反向指，4.Next.Next=4 等价于 5.Next=4
	head.Next = nil               // 第一次遍历：4.Next=nil 从4这里断开防止死循环
	return cur                    // cur就是反转后的头结点，5
}

// @lc code=end
