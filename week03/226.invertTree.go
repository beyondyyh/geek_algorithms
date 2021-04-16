package week03

// 翻转一棵二叉树。
// 输入：
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9

// 输出：
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1
// @leetcode: https://leetcode-cn.com/problems/invert-binary-tree

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
