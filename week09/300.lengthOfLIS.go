package week09

// 300. 最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 示例 1：
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// @lc: https://leetcode-cn.com/problems/longest-increasing-subsequence/

// 方法一：动态规划，二维数组
// 1. dp状态定义：dp[i]表示以num[i]结尾的「上升子序列」的长度，nums[i]必须最后一个元素
// 2. dp转移方程：dp[i] = max{dp[i], 1+dp[j] for j < i if nums[j] < nums[i]}
//	解释：i->0..n-1, j->0..i, 只要当前数字 nums[i] 大于它前面的某个数 nums[j] ，那么 nums[i] 就可以接在这个数的后面形成一个更长的上升子序列
// 3. dp初始化：dp[0..n] = 1 1个数字的递增子序列长度必然是1
// 4. 最终结果：res = max(dp[0..n])，迭代的过程中可以求得最大值
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	// dp状态定义，并初始化
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	// dp方程递推
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
				res = max(res, dp[i])
			}
		}
	}
	return res
}
