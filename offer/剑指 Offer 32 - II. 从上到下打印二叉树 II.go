package offer

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
// @lc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// queue
func levelOrderII(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root} // Push
	for len(q) > 0 {
		n := len(q)
		level := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:] // Pop
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
