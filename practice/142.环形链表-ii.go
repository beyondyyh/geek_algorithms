package practice

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法一：快慢指针
// slow, fast起始位置都是head，然后slow每次走一步，fast每次走2步，直到相遇即为x
// 相遇时，快指针一定是比慢指针多走了n圈，此时再开一个指针指向head，与slow指针一起走再次相遇时，就是入环的交点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			break
		}
	}

	// 快指针返回头部
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}

// @lc code=end
