package offerII

// 剑指 Offer II 055. 二叉搜索树迭代器
// 实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
// BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
// boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
// int next()将指针向右移动，然后返回指针处的数字。
// 注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
// 可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

// 输入
// inputs = ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// 输出
// [null, 3, 7, true, 9, true, 15, true, 20, false]

// 解释
// BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
// bSTIterator.next();    // 返回 3
// bSTIterator.next();    // 返回 7
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 9
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 15
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 20
// bSTIterator.hasNext(); // 返回 False
// 链接：https://leetcode-cn.com/problems/kTOapQ

// 二叉搜索树BST满足以下3个特点：
// 1. nil是bst
// 2. 左子树小于根节点，右子树大于根节点，同理可的所有子树都满足此特性
// 3. 鉴于第2条特点，满足inorder是严格递增的数组
type BSTIterator struct {
	index int
	nums  []int
}

func NewBSTIterator(root *TreeNode) BSTIterator {
	return BSTIterator{
		index: -1,
		nums:  inorder(root),
	}
}

func (bst *BSTIterator) Next() int {
	bst.index += 1
	return bst.nums[bst.index]
}

func (bst *BSTIterator) HasNext() bool {
	return bst.index+1 < len(bst.nums)
}

// 左->根->右
func inorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)
	return res
}
