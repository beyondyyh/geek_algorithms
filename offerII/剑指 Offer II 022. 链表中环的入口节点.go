package offerII

// 剑指 Offer II 022. 链表中环的入口节点
// 给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
// 说明：不允许修改给定的链表。
// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。

// 方法一：快慢指针
// slow, fast起始位置都是head，然后slow每次走一步，fast每次走2步，直到相遇记为x
// 相遇时，快指针一定是比慢指针多走了n圈，此时再将fast指针指向head，与slow指针一起走再次相遇时，就是入环的交点
// 时间复杂度：O(n)，整体呈线性时间复杂度
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
		slow, fast = fast.Next, slow.Next
	}
	return fast
}
