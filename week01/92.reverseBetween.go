package week01

// 92. 反转链表 II
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表 。
// 示例 1：
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
// @lc: https://leetcode-cn.com/problems/reverse-linked-list-ii/

// 方法一：找到[left,right]子区间，反转子链表后再拼接
//	1. 从虚拟头节点走 left-1 步，来到 left 节点的前一个节点
// 	2. 从 pre 再走 right-left+1 步，来到 right 节点（注意此时是right节点，而不是right的前一个 或 后一个节点）
// 	3. 切断出一个子链表（截取链表），同第 206 题，反转链表的子区间
//	4. 接回到原来的链表中
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseBetween1(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Val: -1} // 哑结点，链表通用解法
	dummy.Next = head
	// 1. 从虚拟头节点走 left-1 步，来到 left 节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 2. 从 pre 再走 right-left+1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 3. 切断出一个子链表（截取链表），并反转
	leftNode := pre.Next  // 子链表的头结点
	cur := rightNode.Next // right的下一个节点，先保存起来，便于第4步拼接

	// 切断
	pre.Next = nil
	rightNode.Next = nil
	// 反转子链表
	reverseList(leftNode)

	// 4. 接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = cur

	return dummy.Next
}

// reverseList 反转单链表
// 双指针，先保存下一个节点为tmp，然后把cur反转（cur.Next指向前一个），最后再同时更新pre和cur
func reverseList(head *ListNode) {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tmp := cur.Next // 先保存下一个节点，因为马上要断开
		cur.Next = pre  // 反转操作
		pre = cur       // pre指针后移
		cur = tmp       // cur指针后移
	}
}

// 方法二：头部插入法，方法一的问题在于：如果left=1，right=n(链表长度)时，会遍历2次链表
// 	1. 从虚拟头节点走 left-1 步，来到 left 节点的前一个节点
//	2. 从 left-1..right 区间依次将当前节点插入到 子区间的头部，也就是拼接到 pre后面
//	第二步的具体步骤：
//	2.1. pre指针永远不动，指向的是left 的前一个节点
// 	2.2. cur指针指向 待反转区域的第一个节点 left
//	2.3. next指针永远指向 cur的下一个节点，循环过程中，cur 变化后 next 会随着变化
// @ref: https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Val: -1} // 哑结点，链表通用解法
	dummy.Next = head

	// 	1. 从虚拟头节点走 left-1 步，来到 left 节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
