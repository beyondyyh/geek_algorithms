package practice

import "math"

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32

	// 返回当前节点能为父亲提供的收益
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		// 如果当前节点为叶子节点，则对父节点提供的收益为0
		if root == nil {
			return 0
		}

		// 分别计算左右子节点，对自身的贡献
		left, right := dfs(root.Left), dfs(root.Right)
		// 更新最大值，就是当前节点的val 加上左右节点的贡献。
		res = max(res, root.Val+left+right)
		// 计算当前节点能为父亲提供的最大收益，必须是把 val 加上！
		maxVal := max(left, right) + root.Val
		// 如果收益为0，则返回0
		return max(maxVal, 0)
	}

	dfs(root)
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
