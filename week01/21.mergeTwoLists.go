package week01

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// @leetcode: https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 方法一：双指针法
// 与合并2个有序数组一样使用双指针法，l1和l2分别初始化i，j指针都从头部开始遍历，把l1和l2中较小值赋值给 sortedList，同时自身指针后移一位，
// 直到其中一个链表走完，然后将剩下一个还未走完的链表直接追加给 sortedList 即可
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1} // 哑结点，链表里通用解法，可以巧妙的避免各种边界值情况
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// 处理其中一个尚未遍历完的case
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next // 哑结点
}

// 方法二：递归
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}
