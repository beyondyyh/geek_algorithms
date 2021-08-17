package offer

// 剑指 Offer 42. 连续子数组的最大和
// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
// 要求时间复杂度为O(n)。
// 示例1:
// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// @lc: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

// dp状态定义：dp[i]表示[0..i]的最大子序和
// dp转移方程：dp[i] = max(dp[i-1]+nums[i], nums[i])，nums[i]单独作为一段 或 加到前面一段里，二者较大值
// dp初始化：dp[0] = nums[0]
// 最终结果：max(dp)，迭代的过程中可以求得
func maxSubArray(nums []int) int {
	n, res := len(nums), nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}
