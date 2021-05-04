package week05

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
// 示例 1：
// 输入：m = 3, n = 7
// 输出：28
// @leetcode: https://leetcode-cn.com/problems/unique-paths

// 方法一：自顶向下 top-down，DP搜索
// 状态定义：dp[i][j]表示走到格子(i,j)的方法数
// DP方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 当前位置(i,j)只能是从上边或左边走过的，dp[i][j]表示：从“Start”位置走到(i,j)位置的不同路径数，dp[m-1][n-1]是最终结果
// 时间复杂度：O(m * n)
func uniquePaths1(m int, n int) int {
	// 1. 初始化m x n的网格，第一列和第一行初始化为1，因为从“Start”只有一种方式能到达此处
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 2. 自顶向底，状态定义：dp[i][j]表示从“Start”走到位置i,j的路径数，根据DP方程 dp[i][j] = dp[i-1][j] + dp[i][j-1] 进行递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	// 3. 返回结果
	return dp[m-1][n-1]
}

// 方法二：自底向上 bottom-up，DP搜索
// 状态定义：dp[i][j]表示从格子(i,j)走到“Finished”的方法数
// DP方程：dp[i][j] = dp[i+1][j] + dp[i][j+1]
// 当前位置(i,j)只能往右走或往下走，dp[i][j]表示：从当前位置(i,j)走到“Finished”位置的不同路径数，dp[0][0]是最终结果
func uniquePaths2(m int, n int) int {
	// 1. 初始化m x n的网格，最后一列和最后一行都初始化为1，因为只有一种方法能到达“Finished”处
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][n-1] = 1
	}
	for j := 0; j < n; j++ {
		dp[m-1][j] = 1
	}

	// 2. 自底向顶 根据DP方程 dp[i][j] = dp[i+1][j] + dp[i][j+1] 进行递推
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	// 3. 结果
	return dp[0][0]
}
