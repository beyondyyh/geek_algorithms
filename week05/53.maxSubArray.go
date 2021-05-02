package week05

import "sort"

// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 示例 1：
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// @lc: https://leetcode-cn.com/problems/maximum-subarray/

// 动态规划
// DP状态数组定义：dp[i]
// DP状态转移方程：dp[i] = max(0, dp[i-1]) + nums[i]
// 最大子序和 = 当前元素最大，或者 包含之前最大元素+自身后最大
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	copy(dp, nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(0, dp[i-1]) + nums[i]
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}
