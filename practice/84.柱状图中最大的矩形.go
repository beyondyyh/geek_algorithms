package practice

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
// 方法一：暴力枚举
// 思路：遍历i->[0..n), j->[i+1..n-1], width = j-i+1, minHeight = min[i..j], area=width*minHeight 遍历的过程中迭代res
// 时间复杂度：O(n^3) 接近O(n^3)  2层遍历O(n^2) + (i->j)找最小值 O(k) k最大为n，故为O(n^3)
// 空间复杂度：O(1)
func largestRectangleArea1(heights []int) int {
	res, n := len(heights), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {
			width := j - i + 1
			minHeight := getMinHeight(i, j, heights)
			res = max(res, width*minHeight)
		}
	}
	return res
}

// getMinHeight [i..j]找到区间最小高度
func getMinHeight(i, j int, heights []int) int {
	minHeight := heights[j]
	for ; i < j; i++ {
		if heights[i] < minHeight {
			minHeight = heights[i]
		}
	}
	return minHeight
}

// 方法二：只需一次遍历，
// 思路：i->[0..n) 以当前柱子为中心向左右进行扩散，向左找到第一个比它矮的柱子即为left bound，向右找到第一个比它矮的柱子作为right bound
// width=right-left+1， height=heights[i]，不断更新maxarea
// 时间复杂度：O(n^2) 一层循环O(n)，循环内左右扩散最大为n
// 空间复杂度：O(1)
func largestRectangleArea2(heights []int) int {
	res, n := 0, len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	for i := 0; i < n; i++ {
		// 当前柱子的高度
		height := heights[i]
		// 左右bound下标
		left, right := i, i
		// 向左扩散 找到第一个比它矮的柱子 作为left bound，只要heights[left-1]比当前柱子高，则一直扩散直到最左边
		for left-1 >= 0 && heights[left-1] >= height {
			left--
		}
		// 同理可得 向右扩散找到 right bound
		for right+1 <= n-1 && heights[right+1] >= height {
			right++
		}
		// 以当前柱子为中心的宽度
		width := right - left + 1
		res = max(res, width*height)
	}
	return res
}

// 方法三：单调栈
// 思路：维护一个单调递增的栈，遍历柱子，
//	1.如果遍历到的柱子高度比栈顶柱子大：将柱子的下标入栈，便于计算宽度，
//	2.如果遍历到的柱子高度比栈顶柱子小：说明栈顶柱子的右边界可以确定了就是i，由于栈单调递增，它的左边界就是栈顶的下一个元素，此时可以处理栈顶柱子了。
//		那么：top=stack.Pop(), height=heights[top], width=i-stack.Top()+1，更新maxarea
//	最后：如果栈不为空，则依次执行1，2，将栈搞空
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func largestRectangleArea(heights []int) int {
	n, maxarea := len(heights), 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}

	// 维护一个单调递增的栈，方便操作可以将下标i 和 val一起放入栈，其实没必要只放索引下标即可
	stack := &Stack{}
	stack.Push([2]int{-1, -1})
	for i := 0; i < n; i++ {
		// 当前柱子高度比栈顶小，可以处理栈顶的柱子了
		for heights[i] <= stack.Peek().([2]int)[1] {
			// 弹出栈顶柱子，开始干他
			curHeight := stack.Pop().([2]int)[1]
			curWidth := i - stack.Peek().([2]int)[0] - 1
			maxarea = max(maxarea, curWidth*curHeight)
		}
		// 当前柱子比栈顶大，直接入栈
		stack.Push([2]int{i, heights[i]})
	}

	// 把栈搞空
	for stack.Len() != 1 { // 栈底的-1不用处理
		curHeight := stack.Pop().([2]int)[1]
		curWidth := n - stack.Peek().([2]int)[0] - 1
		maxarea = max(maxarea, curHeight*curWidth)
	}
	return maxarea
}

// @lc code=end
