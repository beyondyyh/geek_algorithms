package offer

// 剑指 Offer 26. 树的子结构
// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
// 例如:
// 给定的树 A:
//      3
//     / \
//    4   5
//   / \
//  1   2
// 给定的树 B：
//    4
//   /
//  1
// 返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

// 示例 1：
// 输入：A = [1,2,3], B = [3,1]
// 输出：false

// 示例 2：
// 输入：A = [3,4,5,1,2], B = [4,1]
// 输出：true
// @lc: https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

// 时间复杂度：O(m*n) 	m,n分别为树 A和 树 B 的节点数量；前序遍历树 A 占用 O(m) ，每次调用 contains(A, B) 判断占用 O(n) 。
// 空间复杂度：O(m) ： 当树 A 和树 B 都退化为链表时，递归调用深度最大。
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 以A为根的树是否包含B树，（必须是从A开始）
	// 1. 当节点 B 为空：说明树 B 已匹配完成（越过叶子节点），因此返回 true；
	// 2. 当节点 A 为空：说明已经越过树 A 叶子节点，即匹配失败，返回 false；
	// 3. 当节点 A 和 B 的值不同：说明匹配失败，返回 false；
	var contains func(*TreeNode, *TreeNode) bool
	contains = func(A, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return contains(A.Left, B.Left) && contains(A.Right, B.Right)
	}

	// 题目约定：空树不是任意一个树的子结构
	if B == nil || A == nil {
		return false
	}
	// 前序遍历A树（根->左->右），A树中的任意节点包含B树就返回true
	return contains(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
