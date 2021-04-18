package week03

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
// 假设一个二叉搜索树具有如下特征：
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:
// 输入:
//     2
//    / \
//   1   3
// 输出: true
// @leetcode: https://leetcode-cn.com/problems/validate-binary-search-tree

func isValidBST(root *TreeNode) bool {
	var isValid func(*TreeNode, int) bool
	isValid = func(root *TreeNode, preVal int) bool {
		// terminator
		if root == nil {
			return true
		}
		// fmt.Printf("val:%d, preVal:%d\n", root.Val, preVal)
		// 访问左子树
		if !isValid(root.Left, preVal) {
			return false
		}
		// 访问当前节点：如果当前节点小于等于中序遍历的前一个节点，说明不满足BST，返回 false；否则继续遍历。
		if root.Val <= preVal {
			return false
		}
		preVal = root.Val
		// 访问右子树
		return isValid(root.Right, preVal)
	}

	return isValid(root, null)
}
