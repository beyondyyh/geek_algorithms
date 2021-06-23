/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 本质上就是 LeverOrder(用队列)，只需要把奇数层的元素翻转即可
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder(root *TreeNode) [][]int {
	q := &Queue{}
	if root != nil {
		q.Push(root)
	}

	// 反转数组
	reverse := func(nums []int) {
		n := len(nums)
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}

	// level当前第几层
	level, res := 0, [][]int{}
	for !q.IsEmpty() {
		// n 当前层节点个数，vals当前层的值
		n, vals := q.Len(), []int{}
		for i := 0; i < n; i++ {
			node := q.Pop().(*TreeNode)
			vals = append(vals, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		// 如果当前层为奇数，则反转vals
		if level&1 == 1 {
			reverse(vals)
		}
		res = append(res, vals)
		level++
	}
	return res
}

// @lc code=end