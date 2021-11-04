package offerII

// 给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
// 示例 1：
// 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/NYBBNL

// 方法一：
// 1. 中序遍历将节点放到数组中
// 2. 修改数组中每个节点的左右指针指向；把左指针设置为nil，右指针指向数组的下一个节点
// 时间复杂度：O(n)，每个节点访问了2次，inorder一次，遍历数组一次，常数级
// 空间复杂度：O(n), 递归调用栈O(n)，数组O(n)，常数级
func increasingBST(root *TreeNode) *TreeNode {
	// inorder standard code
	var inorder func(*TreeNode) []*TreeNode
	inorder = func(root *TreeNode) []*TreeNode {
		res := []*TreeNode{}
		if root == nil {
			return res
		}
		res = append(res, inorder(root.Left)...)
		res = append(res, root)
		res = append(res, inorder(root.Right)...)
		return res
	}

	nodes := inorder(root)
	dummy := &TreeNode{Val: -1}
	cur := dummy
	for _, node := range nodes {
		node.Left, node.Right = nil, nil
		cur.Right = node
		cur = cur.Right
	}
	return dummy.Right
}

// 方法二：
// 只需要一个变量 pre 保存在中序遍历时，上一次被访问的节点。那么我们每次遍历的时候：
// 1. 把当前节点 root.Left设置为nil；
// 2. 把 pre.Right 指向当前节点root；
// 3. 把当前节点 root 赋值给 pre，也就相当于是pre指针后移了一位
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{Val: -1}
	pre := dummy

	// inorder code
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		root.Left = nil
		pre.Right = root
		pre = root
		inorder(root.Right)
	}

	inorder(root)
	return dummy.Right
}
