package week03

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树T的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
// 示例 1:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
// @leetcode: https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof

// DFS，前序遍历
// 时间复杂度： O(n) 其中n为二叉树节点数；最差情况下，需要递归遍历树的所有节点。
// 空间复杂度： O(n) 最差情况下，递归深度达到n ，系统使用o(n)大小的额外空间。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// terminator 递归终止条件，当前节点为空或 找到了p或q
	// 如果树为空，直接返回
	// 如果p和q中有等于root的，那么它们的最近公共祖先即为root（一个节点也可以是它自己的祖先）
	if root == nil || p == root || q == root {
		return root
	}

	// 递推逻辑：drill down
	// 递归遍历左子树，只要在左子树中找到了p或q，则先找到谁就返回谁
	left := lowestCommonAncestor(root.Left, p, q)
	// 递归遍历右子树，只要在右子树中找到了p或q，则先找到谁就返回谁
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果left和right均不为空，说明p、q节点分别在root异侧, 最近公共祖先即为root
	if left != nil && right != nil {
		return root
	}
	// 否则在左子树中p和q都找不到，则p和q一定都在右子树中，右子树中先遍历到的那个就是最近公共祖先（一个节点也可以是它自己的祖先）
	if left == nil {
		return right
	}
	// 否则，如果left不为空，则p和q一定都在左子树中，左子树中先遍历到的那个就是最近公共祖先（一个节点也可以是它自己的祖先）
	return left
}
