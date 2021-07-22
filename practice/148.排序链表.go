package practice

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 归并排序思路：采用分治的思想，递归地将数组不断的切分成左右两部分，直到剩下最后一个元素可以任务是天然有序的，
// 然后合并2个有序数组，合并的过程称为 2-路归并
// 重点在于：
// 1. 如何找到链表中间节点，快慢指针
// 2. 合并2个有序链表（数组），双指针
func sortList(head *ListNode) *ListNode {
	// 空或1个节点可以认为是有序的
	if head == nil || head.Next == nil {
		return head
	}
	// pre指向slow指针的前驱
	var pre *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	// 此时slow指向的就是中间节点，从slow的前面断开。 left, right = head, slow
	pre.Next = nil

	// 2-路归并
	merge := func(left, right *ListNode) *ListNode {
		dummy := &ListNode{Val: -1}
		cur := dummy
		for left != nil && right != nil {
			if left.Val <= right.Val {
				cur.Next = left
				left = left.Next // 较小元素指针后移
			} else {
				cur.Next = right
				right = right.Next
			}
			cur = cur.Next // 自身指针后移
		}
		// 处理其中一个链表未遍历完的情况
		if left != nil {
			cur.Next = left
		}
		if right != nil {
			cur.Next = right
		}
		return dummy.Next
	}

	left, right := sortList(head), sortList(slow)
	return merge(left, right)
}

// @lc code=end
