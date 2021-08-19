package offer

// 剑指 Offer 18. 删除链表的节点
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点。
// 注意：此题对比原题有改动
// 示例 1:
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
// @lc: https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	cur := dummy
	for cur.Next.Val != val { // 不相等则继续遍历
		cur = cur.Next
	}
	cur.Next = cur.Next.Next // 找到之后直接删除
	return dummy.Next
}
