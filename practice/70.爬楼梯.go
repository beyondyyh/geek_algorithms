package practice

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

// 方法一：动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 方法二：滚动迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func climbStairs(n int) int {
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		p, q = q, p+q
	}
	return q
}

// @lc code=end
