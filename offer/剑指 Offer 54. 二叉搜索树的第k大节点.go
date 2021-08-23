package offer

// 剑指 Offer 54. 二叉搜索树的第k大节点
// 给定一棵二叉搜索树，请找出其中第k大的节点。

// 示例 1:
// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 4

// 示例 2:
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// 输出: 4
// @lc: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

// BST满足以下特点：
// 1.空树是BST
// 2.左子树小于根节点，右子树大于根节点
// 3.inorder为升序数组 左->根->右
// 第k大，可以inorder的逆序遍历（右->根->左）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func kthLargest(root *TreeNode, k int) int {
	res := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// inorder 的逆序：右->根->左
		dfs(root.Right)
		if k == 0 {
			return
		}
		k--
		res = root.Val
		dfs(root.Left)
	}
	dfs(root)
	return res
}
