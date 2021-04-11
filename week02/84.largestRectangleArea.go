package week02

import "fmt"

// 84. 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 示例:
// 输入: [2,1,5,6,2,3]
// 输出: 10
// @leetcode: https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

// 方法一：暴力枚举
// 外层遍历: 左边界i -> 0, n-2，内层遍历：右边界j -> i+1, n-1，（i->j)找到最小高度，算出area，同时更新max-area
// 时间复杂度：O(n^3) 接近O(n^3)  2层遍历O(n^2) + (i->j)找最小值 O(k) k最大为n，故为O(n^3)
func largestRectangleArea1(heights []int) int {
	n, maxarea := len(heights), 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	for i := 0; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			width := j - i + 1
			minHeight := getMinHeight(i, j, heights) // (i->j)找到最小高度
			area := width * minHeight                // 算出area
			// fmt.Printf("i:%d j:%d width:%d minHeight:%d area:%d maxArea:%d\n", i, j, width, minHeight, area, maxarea)
			maxarea = maxInt(maxarea, area)
		}
	}
	return maxarea
}

func getMinHeight(i, j int, heights []int) int {
	minHeight := heights[j]
	for ; i < j; i++ {
		if heights[i] < minHeight {
			minHeight = heights[i]
		}
	}
	return minHeight
}

// 方法二：暴力加速
// 枚举棒子高度，一层遍历，然后分别找到该棒子的左右两边第一个比它矮的棒子，即为它的左右边界
// for i -> 0, n-1:
// 		找到left bound, right bound; area = height[i]*(right-left+1), 更新max-area
// 时间复杂度： O(n^2)
func largestRectangleArea2(heights []int) int {
	n, maxarea := len(heights), 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	for i := 0; i <= n-1; i++ {
		// 当前棒子高度
		height := heights[i]
		// 分别从当前位置开始向左右去找左右边界
		left, right := i, i
		// 往左直到找到第一个高度比height小的，此处的left是有效的左边界索引，不是第一个比height小的索引
		for left-1 >= 0 && heights[left-1] >= height {
			left--
		}
		// 往右直到找到第一个高度比height小的
		for right+1 <= n-1 && heights[right+1] >= height {
			right++
		}
		// 计算area
		area := (right - left + 1) * height
		// update max-area
		maxarea = maxInt(maxarea, area)
		// fmt.Printf("i:%d height:%d left:%d right:%d area:%d maxarea:%d\n", i, height, left, right, area, maxarea)
		// ex: i:0 height:2 left:0 right:0 area:2 maxarea:2
		// i:1 height:1 left:0 right:5 area:6 maxarea:6
	}
	return maxarea
}

// -----实现基本栈操作----
// 0:索引， 1:height
type Istack [][2]int

func (s *Istack) Len() int {
	return len(*s)
}

func (s *Istack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Istack) Push(item [2]int) {
	*s = append(*s, item)
}

// 弹出栈顶元素
func (s *Istack) Pop() [2]int {
	item := (*s)[s.Len()-1]
	*s = (*s)[0 : s.Len()-1]
	return item
}

func (s *Istack) Peek() [2]int {
	return (*s)[s.Len()-1]
}

// 方法三：单调栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func largestRectangleArea3(heights []int) int {
	n, maxarea := len(heights), 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}

	// 维护一个从底到顶递增的单调栈
	stack := &Istack{}
	stack.Push([2]int{-1, -1})
	for i := 0; i < n; i++ {
		// 如果当前棒子比栈顶大，则入栈
		// 否则说明当前棒子的右边界找到了，左边界怎么找呢？由于栈是有序的，所以依次比较栈顶元素直到栈顶比它小此时找到左边界，此时可以处理当前棒子了，
		for heights[i] <= stack.Peek()[1] {
			curHeight := stack.Pop()[1]
			curWidth := i - stack.Peek()[0] - 1
			maxarea = maxInt(maxarea, curWidth*curHeight)
		}
		stack.Push([2]int{i, heights[i]})
	}

	// 遍历结束 i=n，把栈清空
	for stack.Len() != 1 {
		curHeight := stack.Pop()[1]
		curWidth := n - stack.Peek()[0] - 1
		fmt.Printf("stack:%+v top:%+v height:%d width:%d-%d-(-1) area:%d\n", *stack, stack.Peek(), curHeight, n, stack.Peek()[0], curHeight*curWidth)
		maxarea = maxInt(maxarea, curHeight*curWidth)
	}
	return maxarea
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
