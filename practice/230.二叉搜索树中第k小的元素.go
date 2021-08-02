package practice

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
// BST中序遍历是升序数组
func kthSmallest(root *TreeNode, k int) int {
	var mink, cnt int
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		cnt++
		if cnt == k {
			mink = root.Val
			return
		}
		inorder(root.Right)
	}

	inorder(root)
	return mink
}

// @lc code=end
