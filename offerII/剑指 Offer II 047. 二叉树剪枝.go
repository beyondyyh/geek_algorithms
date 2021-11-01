package offerII

// 剑指 Offer II 047. 二叉树剪枝
// 给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
// 节点 node 的子树为 node 本身，以及所有 node 的后代。
// 示例 1:
// 输入: [1,null,0,0,1]
// 输出: [1,null,0,null,1]
// 解释:
// 只有红色节点满足条件“所有不包含 1 的子树”。
// 右图为返回的答案。
// 示例 2:
// 输入: [1,0,1,0,0,0,1]
// 输出: [1,null,1,null,1]
// 解释:

// postorder
// 左子树为空、右子树为空、node.val == 0，同时满足3个条件，将改节点删除（root = nil）
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
