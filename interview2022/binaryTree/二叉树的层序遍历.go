package binaryTree

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	res := [][]int{}
	for len(queue) != 0 {
		n, vals := len(queue), []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, vals)
	}
	return res
}
