package week04

import "sort"

// 有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
// 示例 1:
// 输入: k = 5
// 输出: 9
// @leetcode: https://leetcode-cn.com/problems/get-kth-magic-number-lcci

// 方法一：动态规划
// 时间复杂度：O(k)
// 空间复杂度：O(k)
// dp[1] = 1, dp[k] = min(dp[a]*3, dp[b]*3, dp[c]*7)
func getKthMagicNumber1(k int) int {
	var min = func(nums []int) int {
		sort.Ints(nums)
		return nums[0]
	}

	var i3, i5, i7 int
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = min([]int{dp[i3] * 3, dp[i5] * 5, dp[i7] * 7})
		if dp[i] == dp[i3]*3 {
			i3++
		}
		if dp[i] == dp[i5]*5 {
			i5++
		}
		if dp[i] == dp[i7]*7 {
			i7++
		}
	}
	return dp[k-1]
}

// 方法二：小顶堆
func getKthMagicNumber2(k int) int {
	return 0
}
