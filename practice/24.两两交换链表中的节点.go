package practice

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ---------递归三部曲：
// 1、terminator 终止条件
// 2、找返回值，应该给上一级返回什么信息？
// 3、本级递归应该做什么：在这一级递归中，应该完成什么任务？

// 方法一：递归法
// 时间复杂度：O(n)
func swapPairs(head *ListNode) *ListNode {
	// terminator，链表只剩下一个节点或没有节点，没得交换了，递归结束。返回交换好的链表
	if head == nil || head.Next == nil {
		return head
	}

	// 一共三个节点: head, next, swapPairs(next.next)
	// 只需要交换前两个节点，以下三行代码即是
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	// 根据第二步：返回给上一级的是当前已经完成交换后，即处理好了的链表部分
	return next
}

// @lc code=end
