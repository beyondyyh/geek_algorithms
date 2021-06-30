package practice

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */
type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// // 方法一：递归大法
// // 根->Children....
// func preorder1(root *Node) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}
// 	res = append(res, root.Val)
// 	// travel children
// 	for _, node := range root.Children {
// 		res = append(res, preorder(node)...)
// 	}
// 	return res
// }

// 方法二：迭代，栈
//
func preorder(root *Node) []int {
	res := []int{}
	stack := []*Node{}
	stack = append(stack, root) // Push
	for len(stack) > 0 {
		node := stack[len(stack)-1]  // Top
		stack = stack[:len(stack)-1] // Pop
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}

// @lc code=end
