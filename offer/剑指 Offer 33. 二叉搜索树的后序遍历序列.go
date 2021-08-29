package offer

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
// 参考以下这颗二叉搜索树：
//      5
//     / \
//    2   6
//   / \
//  1   3

// 示例 1：
// 输入: [1,6,3,2,5]
// 输出: false

// 示例 2：
// 输入: [1,3,2,6,5]
// 输出: true

// 二叉搜索树BST满足以下特性：
// 	1. 空树是BST
//	2. 左子树小于根节点，右子树大于根节点，所有子树都满足该特性
// 	3. 鉴于特性2，BST的中序遍历（左->根->右）是一个有序的升序数组

// 解题思路：postorder（左->右->根）
// 1. 划分左右子树：遍历postorder数组的 [i,j] 区间，找到 “第一个大于根节点” 的节点，记索引为 m，此时可划分出
//	左子树区间：[i, m-1]、右子树区间：[m, j-1]、根节点索引：j
// 2. 递归判断是否为二叉搜索树：
//	1）、左子树区间 [i, m-1] 的所有节点的值 < postorder[j]。而上面步骤 `1. 划分左右子树` 已经保证了左子树的正确性，因此只需要判断右子树区间即可；
// 	2）、右子树区间 [m, j] 的所有节点的值 > postorder[j]，通过遍历判断，当遇到 <= postorder[j] 的节点则退出返回false，可通过 p = j 判断是否为二叉搜索树。
// 3. p = j，判断此树是否正确，递归左子树和右子树判断 子树是否正确

// 时间复杂度：O(n^2) 每次调用 recur(i,j) 减去一个根节点，因此递归占用 O(n)，最差情况下（即当树退化为链表），每轮递归都需遍历树所有节点，占用O(n)。
// 空间复杂度：O(n)  最差情况下（即当树退化为链表），递归深度将达到 n
func verifyPostorder(postorder []int) bool {
	var recur func(int, int) bool
	recur = func(i, j int) bool {
		if i >= j {
			return true
		}
		p, rootVal := i, postorder[j]
		for postorder[p] < rootVal {
			p++
		}
		// 循环退出时 p 指向右子树第一个节点的索引（循环退出条件 postorder[p] >= rootVal）
		m := p
		for postorder[p] > rootVal {
			p++
		}
		// 循环退出时 p 指向根节点的索引（循环退出条件 postorder[p] <= rootVal）
		// 所以 p == j 即可判断此树是否正确，同理可得：递归判断左右子树
		return p == j && recur(i, m-1) && recur(m, j-1)
	}

	return recur(0, len(postorder)-1)
}
