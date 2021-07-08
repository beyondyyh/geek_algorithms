package practice

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
// 方法三：颜色标记法
// 其核心思想如下：
// 使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色。
// 如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈。
// 如果遇到的节点为灰色，则将节点的值输出。
// 时间复杂度：O(n)
func inorderTraversal(root *TreeNode) []int {
	white, gray := 0, 1
	type item [2]interface{}

	res := []int{}
	stack := []item{}
	stack = append(stack, item{white, root})
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop弹栈
		color, node := top[0].(int), top[1].(*TreeNode)
		if node == nil {
			continue
		}
		if color == white { // Push入栈
			stack = append(stack, item{white, node.Right})
			stack = append(stack, item{gray, node})
			stack = append(stack, item{white, node.Left})
		} else {
			res = append(res, node.Val)
		}
	}
	return res
}

// @lc code=end

// 方法一：递归
func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}

// 方法二：二叉树DFS遍历用stack
func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// 只要栈不为空，就一直搞
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 左子树一插到底
		for root != nil {
			stack = append(stack, root) // Push入栈
			root = root.Left
		}

		// 此时到达最左侧倒数第二层根节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop弹栈
		res = append(res, node.Val)
		// 转向右节点
		root = node.Right
	}
	return res
}
