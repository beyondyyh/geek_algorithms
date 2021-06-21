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
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}

// @lc code=end
