package offer

// 剑指 Offer 47. 礼物的最大价值
// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
// 示例 1:
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
// @lc: https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

// 类似于 [62.不同路径](https://leetcode-cn.com/problems/unique-paths/)
// dp状态定义：dp[i][j]表示从“start”位置走到(i,j)位置的最大权重
// dp转移方程：dp[i][j] = max(dp[i-1][j], dp[i][j-1])+grid[i][j] 只能从上面dp[i-1][j] 或 左边dp[i][j-1] 走过来
// dp初始化：
//      dp[0][0] = grid[0][0],
//      dp[0][j] = dp[0][j-1] + grid[0][j],
//      dp[i][0] = dp[i-1][0] + grid[i][0]
// 最终结果：res = dp[m-1][n-1]
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	// 初始化 m*n 的矩阵，此时矩阵内每个格子的值都为0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	// 初始化第一行
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 初始化第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
