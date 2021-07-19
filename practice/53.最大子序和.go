package practice

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
// 动态规划
// dp状态定义：dp[i]表示以i下标结尾的最大子序和
// dp转移方程：dp[i] = max(dp[i-1]+nums[i], nums[i])，nums[i]单独作为一段 或 加到前面一段里，二者较大值
// dp初始化：dp[0]=nums[0]
// dp结果：迭代的过程中求得
func maxSubArray(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

// @lc code=end
