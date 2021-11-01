package offerII

// 剑指 Offer II 046. 二叉树的右侧视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 示例 1:
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
// 示例 2:
// 输入: [1,null,3]
// 输出: [1,3]
// 示例 3:
// 输入: []
// 输出: []

func rightSideView(root *TreeNode) []int {
	q, res := []*TreeNode{root}, []int{}
	if root == nil {
		return res
	}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:] // Pop
			if i == n-1 {
				res = append(res, node.Val)
			}
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
