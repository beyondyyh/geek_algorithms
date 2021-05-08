package week01

// 61. 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
// @lc: https://leetcode-cn.com/problems/rotate-list/

// 方法一：
// 环形迭代
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 迭代求得链表长度n，由于迭代结束后，头结点指向尾部了，所以用一个临时变量代替head，防止head被修改不便于后续操作
	n, iter := 1, head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	// 当k=n时即便旋转后也是原链表，因此防止k>n多出无用的旋转操作，用k%n
	add := n - k%n
	// 参数k是n的倍数，旋转之后还是原链表，不做无畏的牺牲了直接返回
	if add == n {
		return head
	}
	// 把原链表的尾部指向头部，形成一个环
	iter.Next = head
	// 从iter此时指向原head，从iter开始往后移动n-k处即为旋转后新链表的尾部，需要切断
	for add > 0 {
		iter = iter.Next
		add--
	}

	// 旋转后的新链表的头部即为 ite.Next
	ret := iter.Next
	// 从add出切断环
	iter.Next = nil
	return ret
}

// 方法二：快慢指针
// 题意：将链表每个节点向右移动 k 个位置，相当于把链表的后面 k % n 个节点移到链表的最前面。（n 为 链表长度）
// 步骤：
// 		1. 求链表长度n；
// 		2. 找出倒数第 k+1 个节点（第k+1的next就是k）；
// 		3. 链表重整：将链表的倒数第 k+1 个节点和倒数第 k 个节点断开，并把后半部分拼接到链表的头部
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func rotateRight2(head *ListNode, k int) *ListNode {
	// 如果链表为空 或 就一个节点 或 k为0，直接返回
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 1. 求链表长度n
	n, cur := 0, head
	for cur != nil {
		cur = cur.Next
		n++
	}
	// 对k取模，当k=n时即便旋转后也是原链表，因此防止k>n多出无用的旋转操作
	k %= n
	if k == 0 { // 说明k为n的倍数，此时旋转也是无用功
		return head
	}

	// 2. 快慢指针法找到倒数第 k+1 个节点
	slow, fast := head, head
	// a. 快指针先走 k 步，此时slow和fast之间距离是k，当k=0时退出，此时fast指向的是第 k+1 个节点
	for k > 0 {
		fast = fast.Next
		k--
	}
	// b. 双指环同时走，当fast.Next=nil时，fast指向链表最后一个节点，slow指向倒数第 k+1 个节点
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 3. 链表重建
	// a. newHead 是倒数第 k 个节点，即新链表的头
	newHead := slow.Next
	// b. 让倒数第 k+1 个节点 和 倒数第 k 个节点断开
	slow.Next = nil
	// c. 让最后一个节点指向原始链表的头
	fast.Next = head

	return newHead
}
