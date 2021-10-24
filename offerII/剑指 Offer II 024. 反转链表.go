package offerII

// 剑指 Offer II 024. 反转链表
// 给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]
// 示例 3：
// 输入：head = []
// 输出：[]

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for cur := head; cur != nil; {
		next := cur.Next     // 临时保存next节点
		cur.Next = pre       // cur反向指向pre
		pre, cur = cur, next // pre,cur分别后移
	}
	return pre
}
