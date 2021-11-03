package offerII

import (
	"math"
)

// 剑指 Offer II 044. 二叉树每层的最大值
// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
// 示例1：
// 输入: root = [1,3,2,5,3,null,9]
// 输出: [1,3,9]
// 解释:
//           1
//          / \
//         3   2
//        / \   \
//       5   3   9
// 示例2：
// 输入: root = [1,2,3]
// 输出: [1,3]
// 解释:
//           1
//          / \
//         2   3
// 示例3：
// 输入: root = [1]
// 输出: [1]
// 示例4：
// 输入: root = [1,null,2]
// 输出: [1,2]
// 解释:
//            1
//             \
//              2

func largestValues(root *TreeNode) []int {
	q, res := []*TreeNode{}, []int{}
	if root == nil {
		return res
	}
	q = append(q, root)
	for len(q) > 0 {
		n, levelMax := len(q), math.MinInt64
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			levelMax = max(levelMax, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, levelMax)
	}
	return res
}
