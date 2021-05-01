package week04

// 515. 在每个树行中找最大值
// 您需要在二叉树的每一行中找到最大的值。
// 示例：
// 输入:
//           1
//          / \
//         3   2
//        / \   \
//       5   3   9
// 输出: [1, 3, 9]

// 方法一：BFS，每一层求出最大值放入ans，很好理解
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func largestValues(root *TreeNode) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := []int{}
	q := &Queue{}
	if root != nil {
		q.Push(root)
	}
	for !q.IsEmpty() {
		n, levelMax := q.Len(), -1<<63 // 注意：这里一定要使用固定大小size，不要使用q.Len()，因为q.Len()是不断变化的
		for i := 0; i < n; i++ {
			node := q.Pop().(*TreeNode)
			levelMax = max(levelMax, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		ans = append(ans, levelMax)
	}
	return ans
}
