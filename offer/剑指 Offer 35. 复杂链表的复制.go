package offer

// 剑指 Offer 35. 复杂链表的复制
// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
// 示例 1：
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
// @lc: https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 方法一：hashmap，map中存的是（原节点->新节点）的映射关系
// 时间复杂度：2次遍历，O(n)
// 空间复杂度：O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// map中存的是（原节点->新节点）的映射关系
	m := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	// 将新节点组成一个链表
	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}
	return m[head]
}

// 方法二：拼接+拆分
// 时间复杂度：3次遍历，O(n)
// 空间复杂度：O(1)
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return head
	}

	// 1. 将拷贝节点放到原节点后面，例如1->2->3这样的链表就变成了这样1->1'->2->2'->3->3'
	// 此时拷贝节点的Random指针指向尚未安排，都是nil
	var copy *Node = nil
	for cur := head; cur != nil; cur = cur.Next.Next {
		copy = &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
	}
	// 2. 把拷贝节点的Random指针安排上
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			// cur的Next指针指向的节点就是copy，cur的Random指针指向的节点是原节点，原节点的Next就是copy，也就是新节点的Random指向嘛
			// 这步是关键，贼拉巧妙
			cur.Next.Random = cur.Random.Next
		}
	}
	// 3. 拆分原节点 和 copy节点
	newHead := head.Next
	var tmp *Node = nil
	for cur := head; cur != nil && cur.Next != nil; {
		tmp = cur.Next // copy
		cur.Next = tmp.Next
		cur = tmp
	}
	return newHead
}
