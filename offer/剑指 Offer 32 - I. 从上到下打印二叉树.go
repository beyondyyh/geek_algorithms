package offer

// 剑指 Offer 32 - I. 从上到下打印二叉树
// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回：
// [3,9,20,15,7]
// @lc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// queue
func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	// q = append(q, root) // Push
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:] // Pop
			res = append(res, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}
