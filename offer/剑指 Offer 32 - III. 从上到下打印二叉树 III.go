package offer

// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
// @lc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

// queue
// 根据当前层的奇偶性输出vals，奇数层翻转vals即可
func levelOrderIII(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	// 翻转数组
	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	level := 0             // 当前遍历到的层级
	q := []*TreeNode{root} // init & Push
	for len(q) > 0 {
		n, vals := len(q), []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:] // Pop
			vals = append(vals, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if level&1 == 1 { // 奇数层需要翻转
			reverse(vals)
		}
		res = append(res, vals)
		level++ // 遍历下一层
	}
	return res
}
