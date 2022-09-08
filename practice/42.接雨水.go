package practice

import "container/list"

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
// 方法一：单调栈，将元素下标入栈
// 时间复杂度：O(n) 每个元素最多入栈和出栈一次，所以为O(n)
// 空间复杂度：O(n) 栈的空间
// 思路：
// 1. 当前高度小于等于栈顶高度，入栈，指针后移。
// 2. 当前高度大于栈顶高度，出栈，计算出当前墙和栈顶的墙之间水的多少，然后计算当前的高度和新栈的高度的关系，重复第 2 步。
// 直到当前墙的高度不大于栈顶高度或者栈空，然后把当前墙入栈，指针后移。
func trap(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var res, curIndex int
	stack := list.New()
	for curIndex < len(height) {
		// 如果栈不为空且当前指向的高度大于栈顶元素，就一直遍历，弹栈
		for stack.Len() > 0 && height[curIndex] > height[stack.Back().Value.(int)] {
			h := height[stack.Back().Value.(int)] // 取出要出栈的栈顶元素
			stack.Remove(stack.Back())            // 弹栈
			if stack.Len() == 0 {                 // 栈空就退出
				break
			}

			top := stack.Back().Value.(int)                 // 当前栈顶元素，是索引啊
			distince := curIndex - top - 1                  // 两根柱子间的距离
			minHeight := min(height[top], height[curIndex]) // 两根柱子能形成的面积，高度取较小
			res = res + distince*(minHeight-h)
		}
		stack.PushBack(curIndex) // 当前指向的柱子索引入栈
		curIndex++               // 指针后移
	}
	return res
}

// @lc code=end
