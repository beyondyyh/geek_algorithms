package offer

// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
// 示例 1:
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。

// BST满足：
// 1. 空树是BST
// 2. 左子树小于根节点，右子树大于根节点，所有子节点都满足此属性
// 3. inorder中序遍历为升序

// 解题思路：
// 如果两个节点值都小于根节点，说明他们都在根节点的左子树上
// 如果两个节点值都大于根节点，说明他们都在根节点的右子树上
// 如果一个节点值大于根节点，一个节点值小于根节点，说明他们他们在根节点的两侧，那么root即为公共祖先
func lowestCommonAncestorI(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
