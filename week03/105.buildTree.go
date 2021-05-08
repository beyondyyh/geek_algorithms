package week03

// 根据一棵树的前序遍历与中序遍历构造二叉树。
// 注意: 你可以假设树中没有重复的元素。
// 例如，给出
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：
//     3
//    / \
//   9  20
//     /  \
//    15   7
// @leetcode: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

// 方法一：泛型递归
// 思路：
// 1、preorder 的第一个元素一定是树根，以该值在 inorder 中的索引位置(i)为界：
// 		A）、左半部分是左子树，即：inorder[0:i]，长度为：leftNum = len(inorder[:i])；
//		B）、右半部分是右子树，即：inorder[i+1:]；
// 2、对于同一棵树的根节点，他的左半部分和右半部分不管是前中后序遍历，集合都是不变的，当然长度也是固定的！所以对于inorder来说：
//		A）、左子树部分为：preorder[1:leftNum+1]；
//		B）、右子树部分为：preorder[leftNum+1:]；
// 同理可得，每一步都是如此，所以是标准的递归思想。
// 时间复杂度：O(n^2)，解释：构建树本身需要一次循环，在循环内部又遍历一次inorder查找root节点
// 空间复杂度：O(1)
func buildTree(preorder, inorder []int) *TreeNode {
	// terminator
	if len(preorder) == 0 {
		return nil
	}

	// preorder：根->左->右，inorder：左->根->右
	// 前序遍历的第一个元素一定是树根，找到树根之后在inorder中找到它的位置i，那么整体左子树的长度为（i+1)，那么：
	// 在inorder中左子树为：inorder[0:i]，右子树为：inorder[i+1:]   golang的数组下标是左闭右开区间。
	// 在preorder中左子树为：preorder[1:i+1]，右子树为：preorder[i+1:]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 构建树根
	root := &TreeNode{Val: preorder[0]}
	// 递归构建左子树部分
	root.Left = buildTree(preorder[1:i+1], inorder[0:i])
	// 递归构建右子树部分
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// // 方法二：
// // 递归每次都需要从inoder中查找root节点，这个线性遍历的时间可以优化，将preorder的 val->索引 存入一个map
// // 时间复杂度：O(n)， 一次遍历
// // 空间复杂度：O(n)， 引入了额外空间，hashmap
// func buildTree1(preorder, inorder []int) *TreeNode {
// 	if len(preorder) == 0 {
// 		return nil
// 	}
// 	m := make(map[int]int, len(inorder))
// 	for i, val := range inorder {
// 		m[val] = i
// 	}
// 	fmt.Println(m)

// 	// builder
// 	// - preorder_left 	preorder中左子树的起始位置
// 	// - preorder_right	preorder中右子树的结束位置
// 	// - inorder_left	inorder中左子树的起始位置
// 	// - inorder_right	inorder中右子树的起始位置
// 	var builder func(int, int, int, int) *TreeNode
// 	builder = func(preorder_left, preorder_right, inorder_left, inorder_right int) *TreeNode {
// 		// terminator
// 		if preorder_left > preorder_right {
// 			return nil
// 		}

// 		// 先把根节点建立出来，preorder中的第一个节点就是根节点
// 		root := &TreeNode{Val: preorder[preorder_left]}
// 		// 然后在inorder中定位根节点，返回的是数组下标
// 		inorder_root := m[preorder[preorder_left]]
// 		// 得到左子树中的节点个数
// 		left_subtree_size := inorder_root - inorder_left
// 		// 先序遍历中「从 左边界+1 开始的 size_left_subtree」个元素就对应了中序遍历中「从 左边界 开始到 根节点定位-1」的元素
// 		root.Left = builder(preorder_left+1, preorder_left+left_subtree_size, inorder_left, inorder_root-1)
// 		// 先序遍历中「从 左边界+1+左子树节点数目 开始到 右边界」的元素就对应了中序遍历中「从 根节点定位+1 到 右边界」的元素
// 		root.Right = builder(preorder_left+left_subtree_size+1, preorder_right, inorder_root+1, inorder_right)
// 		return root
// 	}

// 	n := len(preorder)
// 	return builder(0, n-1, 0, n-1)
// }
