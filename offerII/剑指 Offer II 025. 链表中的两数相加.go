package offerII

// 剑指 Offer II 025. 链表中的两数相加
// 给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
// 可以假设除了数字 0 之外，这两个数字都不会以零开头。
// 示例1：
// 输入：l1 = [7,2,4,3], l2 = [5,6,4]
// 输出：[7,8,0,7]
// 示例2：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[8,0,7]

// 此题是第2题的进阶版，所以可以先反转之后，从低位到高位进行相加，最后再反转一下

// 方法二：利用stack保存链表节点，取出时就是从低位到高位了，然后利用头插法不断移动head
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := &Stack{}, &Stack{}
	for ; l1 != nil; l1 = l1.Next {
		stack1.Push(l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		stack2.Push(l2.Val)
	}

	carry := 0               // carry表示进位
	var head *ListNode = nil // 头结点
	// carry > 0 表示：处理最后一位进位的情况，需要多生成一个节点
	for !stack1.IsEmpty() || !stack2.IsEmpty() || carry > 0 {
		n1, n2 := 0, 0
		if !stack1.IsEmpty() {
			n1 = stack1.Pop().(int)
		}
		if !stack2.IsEmpty() {
			n2 = stack2.Pop().(int)
		}
		sum := n1 + n2 + carry      // 两数之和，加上进位的值
		sum, carry = sum%10, sum/10 // 余数是该位上的值，商表示是否需要进位
		node := &ListNode{Val: sum} // 新节点
		node.Next = head            // 头插法
		head = node                 // head前移一位
	}
	return head
}
