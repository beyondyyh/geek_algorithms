package offerII

// 剑指 Offer II 054. 所有大于等于节点的值之和
// 给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 示例 1：
// 输入：root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// 输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
// 示例 2：
// 输入：root = [0,null,1]
// 输出：[1,null,1]
// 示例 3：
// 输入：root = [1,0,2]
// 输出：[3,3,2]
// 示例 4：
// 输入：root = [3,2,4,1]
// 输出：[7,9,4,10]

// 思路：逆向的中序遍历（右->根->左）
// 1. 访问每个节点时，时刻维护变量 sum，保存「比当前节点值大的所有节点值的和」。
// 2. 二叉搜索树的中序遍历，访问的节点值是递增的。
// 3. 如果先访问右子树，反向的中序遍历，访问的节点值是递减的，之前访问的节点值都比当前的大，每次累加给 sum 即可。

// 方法一：inorder递归
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		// terminator
		if root == nil {
			return
		}
		inorder(root.Right) // 先访问右子树
		sum += root.Val     // 节点值累加给sum
		root.Val = sum      // 累加的结果，赋给root.Val
		inorder(root.Left)  // 后访问左子树
	}
	inorder(root) // 递归入口
	return root
}

// 方法二：inorder迭代
func convertBST1(root *TreeNode) *TreeNode {
	sum := 0
	stack, cur := []*TreeNode{}, root
	for cur != nil { // 右子节点先不断压栈
		stack = append(stack, cur)
		cur = cur.Right
	}

	for len(stack) > 0 { // 一直到清空递归栈
		cur = stack[len(stack)-1]    // 栈顶元素
		stack = stack[:len(stack)-1] // =stack.Pop
		sum += cur.Val               // do something
		cur.Val = sum                // do something
		cur = cur.Left               // 找左子节点
		for cur != nil {             // 存在，让左子节点压栈
			stack = append(stack, cur) //
			cur = cur.Right            // 让当前左子节点的右子节点不断压栈
		}
	}
	return root
}
