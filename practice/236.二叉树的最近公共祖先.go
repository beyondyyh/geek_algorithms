package practice

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
// 二叉树的定义本身就符合递归特性，最近重复子问题，所以用递归比较好解决问题
// 思路：
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// terminator 当前节点为空 或 找到了p, q，也就是说如果p或q等于root，那么root就是最近公共祖先
	if root == nil || p == root || q == root {
		return root
	}

	// 递归遍历左子树，只要在左子树中找到了p或q，则先找到谁就返回谁
	left := lowestCommonAncestor(root.Left, p, q)
	// 递归遍历右子树，只要在右子树中找到了p或q，则先找到谁就返回谁
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果left和right都不为空，说明p、q节点分别在root异侧, 最近公共祖先即为root
	if left != nil && right != nil {
		return root
	}
	// 左子树为空，说明在左子树没找到，则p、q祖先在右子树
	if left == nil {
		return right
	}
	// 否则右子树为空，则p、q祖先在左子树
	return left
}

// @lc code=end
