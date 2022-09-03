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
// slow, fast起始位置都是head，然后slow每次走一步，fast每次走2步，直到相遇记为x
// 相遇时，快指针一定是比慢指针多走了n圈，此时再将fast指针指向head，与slow指针一起走再次相遇时，就是入环的交点
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		// 无环特殊处理
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 快指针重新指向头部
	fast = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return fast
}

// @lc code=end
