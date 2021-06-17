package practice

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// ListNode的定义放到 `@lc code=start` 的外面，否则提交到LeetCode会报重复定义的错误
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1、链表通用解法，哑结点
// 2、当前节点的值 l1.Val+l2.Val+carry 要考虑进位的情况
// 3、两个链表可能不等长，短的补0
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 由于在遍历的过程中指针会移动，遍历结束后指针已经不是头结点了，所以申明哑结点便于返回
	dummy := &ListNode{Val: -1}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry         // 两数之和
		sum, carry = sum%10, sum/10    // 余数是该位上的值，商表示是否需要进位
		cur.Next = &ListNode{Val: sum} // 相加之和新节点
		cur = cur.Next                 // 新链表指针也后移
	}

	// 处理最后一位进位问题
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// @lc code=end
