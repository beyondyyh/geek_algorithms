package offerII

// 剑指 Offer II 056. 二叉搜索树中两个节点之和
// 给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。
// 示例 1：
// 输入: root = [8,6,10,5,7,9,11], k = 12
// 输出: true
// 解释: 节点 5 和节点 7 之和等于 12
// 示例 2：
// 输入: root = [8,6,10,5,7,9,11], k = 22
// 输出: false
// 解释: 不存在两个节点值之和为 22 的节点

// 二叉搜索树BST满足以下3个特点：
// 1. nil是bst
// 2. 左子树小于根节点，右子树大于根节点，同理可的所有子树都满足此特性
// 3. 鉴于第2条特点，满足inorder是严格递增的数组

func findTarget(root *TreeNode, k int) bool {
	nums := inorder(root)
	// 双指针两端加逼，查找target
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if k == sum {
			return true
		} else if k > sum {
			lo++
		} else {
			hi--
		}
	}
	return false
}
