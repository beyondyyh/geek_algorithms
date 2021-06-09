package daily

// 62. 不同路径
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
// 示例 1：
// 输入：m = 3, n = 7
// 输出：28
// @lc: https://leetcode-cn.com/problems/unique-paths/

// 方法一：动态规划，bottom-up
// dp状态定义：dp[i][j]表示从当前位置走到“Finish”的不同路径
// dp转移方程：dp[i][j] = dp[i+1][j]+dp[i][j+1] 只能从当前位置向下或向右走
// dp状态初始化：dp[m-1][0..n-1]=1 即最后一行走到“Finish”只有向右一种路径，dp[0..m-1][n-1]=1 即最后一列走到“Finish”只有向下一种路径
// 最终结果：dp[0][0]
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func uniquePaths1(m, n int) int {
	// dp定义
	dp := make([][]int, m)
	// 按行号 m 初始化最后一列
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][n-1] = 1
	}
	// 按列号 n 初始化最后一行
	for j := 0; j < n; j++ {
		dp[m-1][j] = 1
	}

	// 递推
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

// 方法二：动态规划，top-down
// dp状态定义：dp[i][j]表示从“Start”走到当前位置的路径不同路径
// dp转移方程：dp[i][j] = dp[i-1][j]+dp[i][j-1] 当前位置只能从上边或左边走过来
// dp状态初始化：dp[0][0..n]=1 即第一行只有向右一种走法，dp[0..m-1][0]=1 即第一列只有向下一种走法
// 最终结果：dp[m-1][n-1]
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func uniquePaths2(m, n int) int {
	dp := make([][]int, m)
	// 按行号m初始化第一列
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	// 按列号n初始化第一行
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
