package practice

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

// 方法一：双指针两端加逼，此题最优解
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxArea(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	res, L, R := 0, 0, n-1
	for L < R {
		area := min(height[L], height[R]) * (R - L)
		res = max(res, area)
		if height[L] <= height[R] {
			L++
		} else {
			R--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
