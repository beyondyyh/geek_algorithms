package practice

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针法：A+B的长度 = B+A的长度
// 为啥不会死循环？如果A和B有交点，一定会返回交点，如果没有交点最后都nil
// 时间复杂度：O(n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		// 如果A没到尾部，则往后走，当A走到尾部时，指向B的头部
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}
		// 如果B没到尾部，则往后走，当B走到尾部时，指向A的头部
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}

// @lc code=end
