package practice

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	res := [][]int{}
	for len(queue) != 0 {
		size, level := len(queue), []int{}
		for i := 0; i < size; i++ {
			// queue Pop
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			// queue Push
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// @lc code=end
