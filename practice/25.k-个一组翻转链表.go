package practice

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 链表分区为：已翻转部分 + 待翻转部分 + 未翻转部分
// 1、初始两个变量 pre 和 end，pre 代表待翻转链表的前驱，end 代表待翻转链表的末尾
// 2、经过k次循环，end 到达末尾，此时记录待翻转链表的后继 next = end.Next
// 3、待翻转链表的区间为：start, end = pre.Next, end，翻转区间链表[start, end]
// 4、然后将三部分再连接起来，然后重置 pre 和 next 指针，进入下一层循环
// 5、特殊情况，当翻转部分长度不足 k 时，在定位 end 完成后，end==null，已经到达末尾，说明题目已完成，直接返回即可
// 时间复杂度：O(n*K)O(n∗K) 最好的情况为 O(n)O 最差的情况为 O(n^2)
// 空间复杂度：O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 哑结点
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	// 1、初始化pre 和 end，分别代表待翻转链表的前驱 和 末尾
	pre, end := dummy, dummy
	// 只要下一个待翻转区间的前期节点不为空，就继续
	for end.Next != nil {
		// 2、经过k次循环，end 到达待翻转链表的末尾
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果end == nil ，说明需要翻转的链表的节点数小于k，不执行翻转。
		if end == nil {
			break
		}

		// 3、生成待翻转区间 [start:end]，前闭后闭区间
		next := end.Next  // 将下一个的节点先保存起来，切断之后就找不到了
		start := pre.Next // 因为 pre 是待翻转链表的前驱，所以pre.Next才不是待翻转区间的头结点，
		end.Next = nil    // 从 end 后面切断，这样才能形成区间

		// 4、反转区间，pre.Next(待翻转链表的前驱) 连接翻转后的链表头部。如：1->2 变成2->1。 dummy->2->1
		pre.Next = reverseHelper(start)
		// 翻转后头节点变到最后了。通过.Next把断开的链表重新链接。
		start.Next = next

		// 翻转结束，将pre换成下次要翻转的链表的前驱。即start
		pre = start
		// 翻转结束，将end置为下次要翻转的链表的前驱。即start
		end = start
	}
	return dummy.Next
}

// reverseHelper 翻转单链表
func reverseHelper(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next // 临时保存下一个节点
		cur.Next = pre   // 当前节点反向指向pre
		pre = cur        // pre指针后移一位
		cur = next       // cur指针后移一位
	}
	return pre
}

// @lc code=end
