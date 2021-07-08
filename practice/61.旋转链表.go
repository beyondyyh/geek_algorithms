package practice

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 方法一：变成环形，然后在 (n - k%n - 1) 处断开
// 时间复杂度：O(n) 最坏情况下需要遍历2次
// 空间复杂度：O(1)
func rotateRight1(head *ListNode, k int) *ListNode {
	// k=0 或 空链表、链表只有一个节点，不需要旋转
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 1. 迭代求链表长度n，迭代结束tail所指的位置即是原链表的尾节点
	n, tail := 1, head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	// 2. 将尾节点指向头结点，形成一个环环
	tail.Next = head

	// 3. 找到需要断开的位置，n-k%n-1，以【1->2->3->4-5, k=7举例，输出：3->4->5->1->2】
	newTail := head
	for i := 0; i < n-k%n-1; i++ { // i->[0..2)
		newTail = newTail.Next
	}
	// 新的头结点指向断开环的位置，此时 newtail 指向2，正好是新链表的尾节点
	newHead := newTail.Next // 3
	newTail.Next = nil      // 断开

	return newHead
}

// 方法二：快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	// k=0 或 空链表、链表只有一个节点，不需要旋转
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	// 1. 迭代求链表长度n
	n, cur := 1, head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	// 对k取模，当k=n时即便旋转后也是原链表，因此防止k>n多出无用的旋转操作
	k %= n
	if k == 0 {
		return head
	}

	// 2. 快慢指针，fast先走 k 步
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	// [1->2->3->4->5, k=2] slow=1, fast=3
	// 此时 slow 和 fast 之间的距离是 k；fast 指向第 k+1 个节点，因为索引从0开始

	// 3. 快慢指针同时走，直到fast走到尾，
	// 当 fast.next 为空时，fast 指向链表最后一个节点，slow 指向倒数第 k + 1 个节点
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	// newHead 是倒数第 k 个节点，即新链表的头
	newHead := slow.Next

	// 4. 让倒数第 k+1 个节点 和 倒数第 k 个节点断开，最后一个节点指向原始链表的头
	slow.Next, fast.Next = nil, head
	return newHead
}

// @lc code=end
