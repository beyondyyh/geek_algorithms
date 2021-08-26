package offer

import "math"

// 剑指 Offer 55 - II. 平衡二叉树
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

// 示例 1:
// 给定二叉树 [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回 true。

// 示例 2:
// 给定二叉树 [1,2,2,3,3,null,null,4,4]

//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// 返回 false。
// @lc: https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

// 1. 计算某个子树node 的左右子节点的最大深度，left, right
// 2. 如果 abs(left-right) <= 1 说明该子树node为平衡二叉树，若所有子树都平衡，则此树平衡
func isBalanced(root *TreeNode) bool {
	// 求某个子树的最大深度
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(depth(root.Left), depth(root.Right)) + 1
	}

	// 空树是平衡树
	if root == nil {
		return true
	}
	// 如果：左右子树高度差小于等于1，并且所有子树都平衡，则此树平衡
	left, right := depth(root.Left), depth(root.Right)
	return int(math.Abs(float64(left)-float64(right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
