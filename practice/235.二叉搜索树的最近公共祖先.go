package practice

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
// 如果两个节点值都小于根节点，说明他们都在根节点的左子树上
// 如果两个节点值都大于根节点，说明他们都在根节点的右子树上
// 如果一个节点值大于根节点，一个节点值小于根节点，说明他们他们在根节点的两侧，那么root即为公共祖先

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	// 在root两侧
	if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
		return root
	}
	// 在root左侧
	if p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}

// @lc code=end
