package offer

// 剑指 Offer 34. 二叉树中和为某一值的路径
// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
// 示例:
// 给定如下二叉树，以及目标和 target = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]
// @lc: https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

// 时间复杂度：O(n)， n为二叉树的节点数，先序遍历需要遍历所有节点。
// 空间复杂度：O(n)， 最差情况下，即树退化为链表时，path 存储所有树节点，使用 O(n) 额外空间
func pathSum(root *TreeNode, target int) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, []int, int)
	dfs = func(root *TreeNode, path []int, target int) {
		// terminator
		if root == nil {
			return
		}

		// current logic，preorder
		path = append(path, root.Val)
		target -= root.Val
		if target == 0 && root.Left == nil && root.Right == nil {
			res = append(res, append([]int{}, path...))
		}
		dfs(root.Left, path, target)
		dfs(root.Right, path, target)

		// revert
		path = path[:len(path)-1]
	}

	dfs(root, []int{}, target)
	return res
}
