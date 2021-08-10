package offer

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 示例 1：
// 输入：head = [1,3,2]
// 输出：[2,3,1]
// @lc: https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

func reversePrint(head *ListNode) []int {
	res := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}

	// 双指针两端夹逼反转数组
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
