package practice

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
// 动态规划，bottom-up
// dp状态定义：dp[i][j]表示从从“start”走到(i,j)位置的不通路径
// dp转移方程：dp[i][j] = dp[i][j-1]+dp[i-1][j] 当前位置只能从上边或左边走过来
// dp初始化：dp[0..m-1][0] = 1, dp[0][0..n-1] = 1 第一列和第一行只有一种走法
// 最终结果：dp[m-1][n-1]
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
