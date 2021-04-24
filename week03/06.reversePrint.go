package week03

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 示例 1：
// 输入：head = [1,3,2]
// 输出：[2,3,1]
// @leetcode: https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

// 方法一：借助栈，最简单方法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reversePrint1(head *ListNode) []int {
	stack := &Istack{}
	for head != nil {
		stack.Push(head.Val)
		head = head.Next
	}

	res := make([]int, 0, stack.Len())
	for !stack.IsEmpty() {
		res = append(res, stack.Pop().(int))
	}
	return res
}

// 方法二：递归
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reversePrint2(head *ListNode) []int {
	size, i := 0, 0
	var res []int
	var recursion func(*ListNode)
	recursion = func(ln *ListNode) {
		// terminator 递归终止条件，走到链表尾部count即是链表的长度
		if ln == nil {
			// fmt.Printf("--size:%d\n", size)
			res = make([]int, size)
			return
		}
		size++
		recursion(ln.Next)
		res[i] = ln.Val
		// fmt.Printf("i:%d size:%d res:%+v\n", i, size, res)
		i++
	}
	recursion(head)

	return res
}
