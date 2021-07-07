package practice

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
// 二叉搜索树特点：
// 1.空树是二叉搜索树
// 2.左子树小于根节点，右子树大于根节点，以此类推
// 3.inorder是升序数组（严格递增）
func isValidBST(root *TreeNode) bool {
	preVal := math.MinInt64
	var isValid func(*TreeNode) bool
	isValid = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		// 访问左子树
		if !isValid(root.Left) {
			return false
		}
		// 访问根，不满足inorder严格递增特点
		if root.Val <= preVal {
			return false
		}
		preVal = root.Val
		// 访问右子树
		return isValid(root.Right)
	}

	return isValid(root)
}

// @lc code=end
