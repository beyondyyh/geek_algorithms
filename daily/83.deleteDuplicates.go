package daily

// 83. 删除排序链表中的重复元素
// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
// 返回同样按升序排列的结果链表。
// 示例 1：
// 输入：head = [1,1,2]
// 输出：[1,2]
// @lc: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil { // 当前节点不为nil 遍历
		for cur.Next != nil && cur.Next.Val == cur.Val { // 只要下一个节点不为nil 且 是重复的，就指针后移
			cur.Next = cur.Next.Next
		}
		cur = cur.Next // 新链表的指针正常后移
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
