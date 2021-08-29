package offer

// 剑指 Offer 07. 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

// 示例 1:
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]

// 示例 2:
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]
// @lc: https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解题思路：
// 1. 前序遍历preorder（根->左->右）preorder的第一个元素是根，设为A
// 2. 在中序遍历inorder（左->根->右）中找到A，设索引为i，则左子树的长度为i，那么：
//      inorder中左子树为：inorder[0:i]（golang的slice操作是左闭右开区间，i指向树根，故不包含i），右子树为：inorder[i+1:]
//      preorder中左子树为：preorder[1:i+1]（左子树部分长度为i），右子树为：preorder[i+1:]
// 3. 同理可得递归构建
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 树根的值
	i, rootval := 0, preorder[0]
	for ; i < len(inorder); i++ {
		if inorder[i] == rootval {
			break
		}
	}
	// 构建新树，golang的slice操作是左闭右开区间
	// 左子树部分为：inorder[0:i], preorder[1:i+1]，右子树部分为：inorder[i+1:], preorder[i+1:]
	root := &TreeNode{Val: rootval}
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
