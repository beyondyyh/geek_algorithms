package week02

// 589. N 叉树的前序遍历
// 给定一个 N 叉树，返回其节点值的 前序遍历 。
// N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
// 进阶：
// 递归法很简单，你可以使用迭代法完成此题吗?
// @lc: https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// 方法一：递归大法
// 根->Children....
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	// travel children
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}
