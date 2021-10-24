package offerII

// 剑指 Offer II 021. 删除链表的倒数第 n 个结点
// 给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]

// 双指针大法
// 初始化快慢指针都指向哑结点，然后fast指针先走n步，此时fast和slow之间的间距为n；
// 然后fast和slow一起走，直到fast走到尾部（fast.Next=nil），此时slow.Next就是倒数第n个节点，删除之
// 返回dummy.Next
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	slow, fast := dummy, dummy
	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
	}
	// 此时 fast和slow 之间的距离为n，以实例1为例：
	// slow, fast = -1, 2
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	// 此时：slow, fast = 3, 5
	// 删除slow.Next节点
	slow.Next = slow.Next.Next
	return dummy.Next
}
