package week04

// 199. 二叉树的右视图
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 示例:
// 输入: [1,2,3,null,5,null,4]
// 输出: [1, 3, 4]
// 解释:
//    1            <---
//  /   \
// 2     3         <---
//  \     \
//   5     4       <---

// BFS，层级遍历，取每层最后一个值
func rightSideView(root *TreeNode) []int {
	q, ans := &Queue{}, []int{}
	if root != nil {
		q.Push(root)
	}

	for !q.IsEmpty() {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Pop().(*TreeNode)
			if i == n-1 { // 将当前level最后一个节点的值放入结果集
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return ans
}
