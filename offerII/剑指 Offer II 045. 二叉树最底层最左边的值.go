package offerII

// 剑指 Offer II 045. 二叉树最底层最左边的值
// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
// 假设二叉树中至少有一个节点。
// 示例 1:
// 输入: root = [2,1,3]
// 输出: 1
// 示例 2:
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7

// 此题考查的是层序遍历，只需要记住每层的第一个元素即可
func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}
	var res int
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				res = node.Val
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
