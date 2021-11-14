package offerII

// 剑指 Offer II 027. 回文链表
// 给定一个链表的 头节点 head ，请判断其是否为回文链表。
// 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
// 示例 1：
// 输入: head = [1,2,3,3,2,1]
// 输出: true
// 示例 2：
// 输入: head = [1,2]
// 输出: false

// 方法一：遍历将链表Vals放入数组，然后判断数组是否为回文
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isPalindromeLinkedList(head *ListNode) bool {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}

// 方法二：快慢指针找到中间节点，切断，翻转前半段，然后遍历反转的前半段和后半段判断是否为回文
// 处理技巧，在快慢指针运动的过程中可以顺手把前半段给翻转了
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func isPalindromeLinkedList2(head *ListNode) bool {
	var slow, fast, pre *ListNode = head, head, nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 反转慢指针走过的路径，以下4行代码与 翻转单链表 一模一样，此处的slow就是cur
		next := slow.Next // 临时保存下一个节点
		slow.Next = pre   // slow反指向前一个节点
		pre = slow        // pre后移一位
		slow = next       // slow后移一位
	}
	// 遍历结束后，slow，fast的位置：
	// 实例1 偶数：slow, fast = 第二个3, 1
	// 实例：1->2->3->2->1 奇数： slow, fast = 3, nil
	// 长度为偶数时，fast最终必为 null
	// 长度为奇数时，当fast.Next为null时，fast必不为 null
	// fmt.Printf("slow:%d, fast:%d\n", slow.Val, fast.Val)
	if fast != nil {
		slow = slow.Next // 此时slow指向的是后半段的头部
	}

	// 判断 pre 和 slow 是否为回文
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre, slow = pre.Next, slow.Next
	}
	return true
}
