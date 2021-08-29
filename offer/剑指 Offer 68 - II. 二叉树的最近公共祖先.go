package offer

// 剑指 Offer 68 - II. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

// 示例 1:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
// @lc: https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 左子树中找p、q，先找到谁就返回谁，或者都没找到返回nil
	left := lowestCommonAncestor(root.Left, p, q)
	// 同理在右子树中找p、q
	right := lowestCommonAncestor(root.Right, p, q)

	// 在左子树没找到，返回right
	if left == nil {
		return right
	}
	// 在右子树没找到，返回left
	if right == nil {
		return left
	}
	// 否则说明p、q在root的两侧，返回root
	return root
}
