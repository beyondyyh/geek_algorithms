package week05

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
// 示例 1：
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
// 解释：
// 3x3 网格的正中间有一个障碍物。
// 从左上角到右下角一共有 2 条不同的路径：
// 1. 向右 -> 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右 -> 向右
// @leetcode: https://leetcode-cn.com/problems/unique-paths-ii

// 方法一：动态规划，一位数组
// 斐波那契：fib(n) = fib(n-1) + fib(n-2)
// 不需要二维数组，一维即可，由第一二行递推第三行，。。
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ { // 注意：j从1开始
			if obstacleGrid[i][j-1] == 1 { // 障碍物
				dp[j] = 0
			} else {
				dp[j] += dp[j-1] // 上一行的dp[j]直接被传下来了
			}
		}
	}
	return dp[n]
}

// 好理解吗？？？
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(dp)-1]
}

// 动态规划：bottom-up，二维数组更易理解
// dp状态定义：dp[i][j]表示走到格子(i,j)的方法数
// dp转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1] 表示到达位置(i,j)的不同路径数，只有2中走法，从上面下来dp[i-1][j] 或 从左边过去dp[i][j-1]
// 障碍物：到达它的路径数=0，即它不能给其他点贡献路径数
// dp初始化：狠重要
//		第 1 列的格子只有从其上边格子走过去这一种走法，因此初始化 dp[i][0] 值为 1，存在障碍物时为 0；
//		第 1 行的格子只有从其左边格子走过去这一种走法，因此初始化 dp[0][j] 值为 1，存在障碍物时为 0
// 最终结果：dp[m-1][n-1]
func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	// 1. 定义dp数组，先初始化dp二维数组，此时格子全是0
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 2.1 初始化第一列，只要遇到“障碍物”则第一列后续的格子就不必看了，都是0
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	// 2.2 初始化第一行，只要遇到“障碍物”则第一行后续的格子就不必看了，都是0
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	// 3. 根据dp方程递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// 4. result
	return dp[m-1][n-1]
}
