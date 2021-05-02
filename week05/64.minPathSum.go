package week05

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例 1：
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 1 	3	1
// 1	5	1
// 4	2	1
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。

// 动态规划
// 状态定义：dp[i][j]表示从开始位置走到(i,j)的最小路径和
// DP方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + dp[i][j]
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 复用grid的数据
	// 初始化：f(0,0) = A[0][0], f(i,0) = sum(0,0 -> i,0)、 f(0,y) = sum(0,0 -> 0,y)
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	// 求解f(i,j)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}
