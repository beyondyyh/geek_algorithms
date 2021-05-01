package week04

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
// 示例：
// 二叉树：[3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层序遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
// @leetcode: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

// 方法一：迭代法，利用队列
// 思路：在每一层遍历开始前，先记录队列中的结点数量n（也就是这一层的结点数量），然后一口气处理完这一层的n个结点
// 然后再将该层节点的孩子节点依次入队
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func levelOrder1(root *TreeNode) [][]int {
	q := &Queue{}
	if root != nil {
		q.Push(root)
	}
	res := [][]int{}
	for !q.IsEmpty() {
		n, level := q.Len(), []int{}
		for i := 0; i < n; i++ { // 注意：这里一定要使用固定大小size，不要使用q.Len()，因为q.Len()是不断变化的
			node := q.Pop().(*TreeNode) // 依次将当前level的节点处队列，同时将当前leve层的节点的 左右子节点再入队列
			level = append(level, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var bfs func(*TreeNode, int)
	bfs = func(node *TreeNode, level int) {
		// level与当前保存的结果集里索引相等，初始当前level为一个空slice
		if len(res) == level {
			res = append(res, []int{})
		}
		// 当前level层的节点值入结果集里对应索引的子数组
		res[level] = append(res[level], node.Val)

		// fmt.Printf("level:%d res:%+v\n", level, res)
		if node.Left != nil {
			bfs(node.Left, level+1)
		}
		if node.Right != nil {
			bfs(node.Right, level+1)
		}
	}
	bfs(root, 0)
	return res
}
