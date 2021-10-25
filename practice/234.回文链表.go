package practice

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法一：首先想到遍历一次链表把Val放到数组里，然后判断是否为回文
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isPalindrome1(head *ListNode) bool {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}

// 方法二：快慢指针，技巧处理：迭代的过程中可以把前半部分（slow指针走过的路径）反转
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func isPalindrome2(head *ListNode) bool {
	var slow, fast, pre *ListNode = head, head, nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针一次走2步
		// 反转慢指针走过的路径，以下4行代码与 反转单链表 一模一样
		next := slow.Next // 临时保存下一个节点
		slow.Next = pre   // slow反指向前一个节点
		pre = slow        // pre后移一位
		slow = next       // slow后移一位
	}
	// 长度为偶数时，fast最终必不为 null
	// 长度为奇数时，当fast.next为null时，fast必为null
	if fast != nil {
		slow = slow.Next
	}
	// 判断 pre 和 slow 是否为回文
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre, slow = pre.Next, slow.Next
	}
	return true
}

// @lc code=end

// 反转单链表，温故而知新
func reverseListx(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next // 临时保存下一个节点
		cur.Next = pre   // 当前节点反向指向前一个
		pre = cur        // pre后移
		cur = next       // cur后移
	}
	return pre
}
