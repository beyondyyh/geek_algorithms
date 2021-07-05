package practice

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
// 方法一：类似300.最长递增子序列 用动态规划
// 不同点在用：本题要求连续所以dp[i]只与dp[i-1]有关
// dp[i+1] = dp[i]+1 if nums[i+1] > nums[i]
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findLengthOfLCIS1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		res = max(res, dp[i+1])
	}
	return res
}

// 贪心算法
// 遇到nums[i + 1] > nums[i]的情况，count就++，否则count为1，记录count的最大值就可以了
func findLengthOfLCIS(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(nums)
	res, cnt := 1, 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] { // 连续记录
			cnt++
		} else { // 不连续，从头开始计数
			cnt = 1
		}
		res = max(res, cnt)
	}
	return res
}

// @lc code=end
