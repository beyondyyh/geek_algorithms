package practice

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 二叉搜索树的特点：
// 1.空树是二叉搜索树
// 2.左子树小于根节点，右子树大于根节点，同理可得
// 3.inorder是升序数组（严格递增）
// 缺点：可能退化为一维的单链表，平衡是指左右子树的高度差不大于1，高度差也称为平衡因子

// 也就是inorder的逆过程，理论上选择任何一个元素作为root都可以，但是题目要求平衡的，所以我们选择数组的中间元素作为根节点，递归的构建
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var dfs func(nums []int, low, high int) *TreeNode
	dfs = func(nums []int, low, high int) *TreeNode {
		// terminator
		if low > high {
			return nil
		}
		mid := (low + high) >> 1
		root := &TreeNode{Val: nums[mid]}
		root.Left = dfs(nums, low, mid-1)
		root.Right = dfs(nums, mid+1, high)
		return root
	}

	return dfs(nums, 0, len(nums)-1)
}

// @lc code=end
