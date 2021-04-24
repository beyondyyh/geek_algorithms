package kit

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL declare
var NULL = -1 << 63

// Ints2Tree convert []int to binary-search-tree
// 广度优先遍历BFS的 逆操作
func Ints2Tree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 注意i可能越界
		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Tree2BFS 二叉树的广度优先遍历
func Tree2BFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return res
}

//------------------------------------------------------------//
// 递归四部曲：                                                 //
// 1. terminator                 终结者                        //
// 2. process current logic      处理当前层逻辑                 //
// 3. dirll down                 向下进入下层逻辑               //
// 4. restvert current status    恢复当前层状态                 //
//------------------------------------------------------------//

// Preorder 前序遍历 根->左->右
func Preorder(root *TreeNode) []int {
	// var res []int
	// if root == nil {
	// 	return res
	// }

	// res = append(res, root.Val)
	// res = append(res, Preorder(root.Left)...)
	// res = append(res, Preorder(root.Right)...)
	// return res

	// 标准的递归写法如下：
	var res []int
	var recursion func(*TreeNode)
	recursion = func(root *TreeNode) {
		// 1. terminator 递归终止条件
		if root == nil {
			return
		}
		// 2. 当前层逻辑
		res = append(res, root.Val)
		// 3. drill down 递推下探
		recursion(root.Left)
		recursion(root.Right)
	}
	recursion(root)
	return res
}

// Inorder 中序遍历 左->根->右
func Inorder(root *TreeNode) []int {
	// var res []int
	// if root == nil {
	// 	return res
	// }

	// res = append(res, Inorder(root.Left)...)
	// res = append(res, root.Val)
	// res = append(res, Inorder(root.Right)...)
	// return res

	// 标准的递归写法如下：
	var res []int
	var recursion func(*TreeNode)
	recursion = func(root *TreeNode) {
		// 1. terminator 递归终止条件
		if root == nil {
			return
		}

		// 3. drill down 递推下探
		recursion(root.Left)

		// 2. 当前层逻辑
		res = append(res, root.Val)

		recursion(root.Right)
	}
	recursion(root)
	return res
}

// Postorder 后续遍历 左->右->根
func Postorder(root *TreeNode) []int {
	// var res []int
	// if root == nil {
	// 	return res
	// }

	// res = append(res, Postorder(root.Left)...)
	// res = append(res, Postorder(root.Right)...)
	// res = append(res, root.Val)
	// return res

	// 标准的递归写法如下：
	var res []int
	var recursion func(*TreeNode)
	recursion = func(root *TreeNode) {
		// 1. terminator 递归终止条件
		if root == nil {
			return
		}
		// 3. drill down 递推下探
		recursion(root.Left)
		recursion(root.Right)
		// 2. 当前层逻辑
		res = append(res, root.Val)
	}
	recursion(root)
	return res
}
