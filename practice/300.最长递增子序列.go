package practice

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
// 动态规划
// dp状态定义：dp[i]表示nums[0:i]的最长递增子序列
// dp转移方程：dp转移方程：dp[i] = max{dp[i], 1+dp[j] for j < i if nums[j] < nums[i]}
//	i->[0..n), j->[0..i) 只要nums[j] < nums[i]，nums[i] 就可以接在nums[j]的后面形成一个更长的上升子序列
// dp初始化：dp[0..n] = 1 1个数组的上升子序列长度必然是1
// 最终结果：res = max(dp[0..n])，以nums[i]结尾的上升子序列的最大值
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// @lc code=end
