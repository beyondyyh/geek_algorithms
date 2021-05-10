package week06

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例 1：
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// @lc: https://leetcode-cn.com/problems/minimum-path-sum/
// 1	3	1
// 1	5	1
// 4	2	1

// 方法一：动态规划 top-down，二维
// dp状态数组：dp[i][j] 表示从左上角走到位置(i,j)的最小路径和
// dp转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1])+grid[i][j]
// 最终结果：dp[m-1][n-1]
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	// 1. 初始化dp状态数组，复用grid
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for j := 1; j < m; j++ {
		grid[j][0] += grid[j-1][0]
	}
	// 2. 根据dp方程递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

// 方法二：bottom-up，二维
// dp状态数组：dp[i][j] 表示从位置(i,j)走到右下角的最小路径和
// dp转移方程：dp[i][j] = min(dp[i+1][j], dp[i][j+1])+grid[i][j]
// 最终结果：dp[0][0]
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	// 1. 初始化dp状态数组，复用grid
	// 按列数，初始化最后一行，从倒数第二个数开始就行，因为最后一个位置只有一条路可走，而且值就是他自己
	for i := n - 2; i >= 0; i-- {
		grid[m-1][i] += grid[m-1][i+1]
	}
	// 按行数，初始化最后一列，从倒数第二个数开始就行，因为最后一个位置只有一条路可走，而且值就是他自己
	for j := m - 2; j >= 0; j-- {
		grid[j][n-1] += grid[j+1][n-1]
	}
	// 2. 根据dp方程递推，直接从倒数第二行、倒数第二列开始递推即可，因为最后一行或最后一列到达右下角的路径只有一条路，而且值就是上面初始化的值
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			grid[i][j] = min(grid[i+1][j], grid[i][j+1]) + grid[i][j]
		}
	}
	return grid[0][0]
}

// 方法三：top-down 优化空间
// 一维数组
func minPathSum3(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	return 0
}
