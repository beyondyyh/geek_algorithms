package offerII

// 剑指 Offer II 026. 重排链表
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//  L0 → L1 → … → Ln-1 → Ln
// 请将其重新排列后变为：
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 示例 1:
// 输入: head = [1,2,3,4]
// 输出: [1,4,2,3]
// 示例 2:
// 输入: head = [1,2,3,4,5]
// 输出: [1,5,2,4,3]

// 方法一：用数组存储链表节点，然后头尾遍历两端加逼
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	// 双指针两端加逼
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j { // 偶数个节点
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 方法二：
// 1.利用快慢指针找到中间节点，断开
// 2.反转后半段
// 3.依次连接两个链表
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// 1.快慢指针找中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 从中间断开
	newHead := slow.Next
	slow.Next = nil

	// 2. 反转后半段
	newHead = reverseList(newHead)

	// 3. 依次连接两个链表
	for newHead != nil {
		temp := newHead.Next
		newHead.Next = head.Next
		head.Next = newHead
		head = newHead.Next
		newHead = temp
	}
}
