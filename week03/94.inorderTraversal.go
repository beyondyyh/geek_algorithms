package week03

// 方法一：递归
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)

	return res
}

func inorderTraversal2(root *TreeNode) []int {
	var res []int
	inorderRecursion(root, &res)
	return res
}

// inorderRecursion 递归四部曲
func inorderRecursion(root *TreeNode, res *[]int) {
	// terminator 终结者，先找返回条件
	if root == nil {
		return
	}

	// drill down  向下进入下一层
	inorderRecursion(root.Left, res)

	// process current logic 处理当前逻辑
	*res = append(*res, root.Val)

	// drill down  向下进入下一层
	inorderRecursion(root.Right, res)

	// revert current status 销毁当前状态
	// nothing todo
}

type item [2]interface{}

// 方法二：颜色标记法
// 其核心思想如下：
// 使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色。
// 如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈。
// 如果遇到的节点为灰色，则将节点的值输出。
// 时间复杂度：O(n)
func inorderTraversal3(root *TreeNode) []int {
	var res []int
	white, gray := 0, 1
	stack := &Istack{item{white, root}}
	for !stack.IsEmpty() {
		// stack.print()
		it := stack.Pop().(item)
		color, node := it[0].(int), it[1].(*TreeNode)
		if node == nil {
			continue
		}
		if color == white {
			stack.Push(item{white, node.Right})
			stack.Push(item{gray, node})
			stack.Push(item{white, node.Left})
		} else {
			res = append(res, node.Val)
		}
	}
	return res
}
