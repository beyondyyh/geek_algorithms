package kit

import (
	"container/list"
	"math"
)

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL declare
var NULL = math.MinInt64

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

//------------------------------------------------------------//
// 非递归遍历：                                                 //
// 		1
// 	 2	    3
// 4   5  6	  7

// 前序：根->左->右，1245367
func PreorderIter(root *TreeNode) []int {
	// 可以用go标准库里的list(基于双链表)，或自己实现的栈(基于数组)
	stack := list.New()

	res := []int{}
	cur := root
	for cur != nil || stack.Len() > 0 {
		for ; cur != nil; cur = cur.Left {
			res = append(res, cur.Val) // visited
			stack.PushBack(cur)
		}
		// fmt.Printf("s:%+v\n", showStack(stack))
		// 栈顶元素，= stack.Peek()
		top := stack.Back()
		cur = top.Value.(*TreeNode)
		// 转向右子树
		cur = cur.Right
		// 出栈，= stack.Pop
		stack.Remove(top)
	}

	return res
}

// preorder 不用系统库，自己维护一个栈
func PreorderIter2(root *TreeNode) []int {
	res := []int{}
	cur, stack := root, []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for ; cur != nil; cur = cur.Left { // 左子树，一插到底
			res = append(res, cur.Val)
			stack = append(stack, cur) // Push
		}
		// 转向右子树
		cur = stack[len(stack)-1].Right // Peek栈顶节点
		stack = stack[:len(stack)-1]    // Pop
	}
	return res
}

// 中序：左->根->右，4251637
func InorderIter(root *TreeNode) []int {
	stack := list.New()

	res := []int{}
	cur := root
	for cur != nil || stack.Len() != 0 {
		// 左子树，一插到底
		for ; cur != nil; cur = cur.Left {
			stack.PushBack(cur)
		} // 124

		// 栈顶元素
		top := stack.Back()
		cur = top.Value.(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
		stack.Remove(top)
	}
	return res
}

// inorder 左->根->右
func InorderIter2(root *TreeNode) []int {
	res := []int{}
	cur, stack := root, []*TreeNode{}
	for cur != nil || len(stack) > 0 {
		for ; cur != nil; cur = cur.Left { // 左子树，一插到底
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]  // 栈顶节点
		res = append(res, cur.Val) // 根
		cur = cur.Right
		stack = stack[:len(stack)-1] // Pop
	}
	return res
}

// 后序：左->右->根，4526731
func PostorderIter(root *TreeNode) []int {
	stack := list.New()

	res := []int{}
	cur := root
	var lastVisit *TreeNode = nil
	for cur != nil || stack.Len() != 0 {
		// 左子树，一插到底
		for ; cur != nil; cur = cur.Left {
			stack.PushBack(cur)
		}
	}

	cur = stack.Back().Value.(*TreeNode)
	if cur.Right == nil || cur.Right == lastVisit {
		res = append(res, cur.Val)
		stack.Remove(stack.Back()) // = s.Pop()
		// 记录上一个访问的节点，用于判断“访问根节点之前，右子树是否已访问过”
		lastVisit = cur
		// 表示不需要转向，继续弹栈
		cur = nil
	} else {
		cur = cur.Right
	}
	return res
}

// showStack, 迭代list
func showStack(l *list.List) []int {
	eles := make([]int, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		eles = append(eles, e.Value.(*TreeNode).Val)
	}
	return eles
}
