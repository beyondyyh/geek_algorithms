package offer

// 剑指 Offer 52. 两个链表的第一个公共节点
// 输入两个链表，找出它们的第一个公共节点。
// 如下面的两个链表：
// 在节点 c1 开始相交。
// @lc: https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

// A+B的长度 = B+A的长度
// 时间复杂度：O(A+B)
// 空间复杂度：O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}
