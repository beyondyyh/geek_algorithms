package practice

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
// 动态规划，bottom-up
// dp状态定义：dp[i][j]表示从“start”走到(i,j)的不同路径
// dp转移方程：如果(i,j)不是障碍物则：dp[i][j] = dp[i][j-1]+dp[i-1][j]，如果(i,j)是障碍物则dp[i][j]=0
// dp初始化：初始化第一行和第一列，如果遇到障碍物，后续就全是0了，此路不通
// 最终结果：dp[m-1][n-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	// 先初始化 m*n 的二维数组，所有元素都为0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 按行号初始化第一列
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	// 按列号初始化第一行
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	// 根据dp方程递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
