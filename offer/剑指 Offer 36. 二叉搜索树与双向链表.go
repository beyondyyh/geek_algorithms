package offer

// 剑指 Offer 36. 二叉搜索树与双向链表
// 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
// 为了让您更好地理解问题，以下面的二叉搜索树为例：
// 我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
// 下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
// @lc: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/

// 思路：
// 1. BST的中序遍历为升序数组，所以一定是中序遍历，
// 2. 设中序遍历中访问当前节点 cur；并在访问每个节点时构建 cur 和前驱节点 pre 的指向；
// 3. 中序遍历完成后，最后构建头节点和尾节点的引用指向即可
// 构建链表的过程：
// a. 当 pre 为空时： 代表正在访问链表头节点，记为 head；
// b. 当 pre 不为空时： 修改双向节点引用，即 pre.Right = cur， cur.Left = pre；也即双链表的操作（pre.Next = cur, cur.Prev = pre）
// c. 保存 cur： 更新 pre = cur，即节点 cur 是后继节点的 pre，或者理解为 pre 指针后移一位。

// 时间复杂度 O(n)，中序遍历需要访问所有节点
// 空间复杂度 O(n)，最差情况下，即树退化为链表时，递归深度达到n，需要O(n)的栈空间
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// pre用于记录双向链表中位于cur左侧的节点，即上一次迭代中的cur
	var head, pre *TreeNode = nil, nil

	// 中序遍历
	var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(root.Left)

		if pre == nil {
			// 当pre==nil时，cur左侧没有节点，即此时cur为双向链表中的头节点
			head = cur
		} else {
			// 反之，pre!=nil 时，cur左侧存在节点pre，需要进行pre.right=cur的操作。
			pre.Right = cur
		}
		// 双向指向， pre.Next=cur, cur.Prev=pre
		cur.Left = pre

		// pre指针后移一位
		pre = cur
		// 处理右子树
		dfs(root.Right)
	}

	// main process
	dfs(root)
	// 尾部和头部节点的相互指向，形成环
	pre.Right = head
	head.Left = pre

	return head
}
