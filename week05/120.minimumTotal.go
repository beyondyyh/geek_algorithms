package week05

// 120. 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
// 示例 1：
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//    2
//   3 4
//  6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
// @lc: https://leetcode-cn.com/problems/triangle/

// 方法一：递归
// 与(i, j) 点相邻的结点为 (i+1, j) 和 (i+1, j+1)
// f(i,j) 为 (i,j) 点到底边的最小路径和，递归公式：f(i,j) = min(f(i+1,j), f(i+1,j+1)) + triangle[i][j]
// 将任一点到底边的最小路径和，转化成了与该点相邻两点到底边的最小路径和中的较小值，再加上该点本身的值
// 时间复杂度：O(2^n)
func minimumTotal1(triangle [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// 1. terminator 走到最下一层（此时i=i+1），路径为0
		if i >= len(triangle) {
			return 0
		}
		// 2. process current logic & drill down
		return min(dfs(i+1, j), dfs(i+1, j+1)) + triangle[i][j]
	}
	return dfs(0, 0)
}

// 方法二：递归+记忆化搜索
func minimumTotal2(triangle [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 记忆化缓存
	memo := make([][]int, len(triangle))
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == len(triangle) {
			return 0
		}
		if memo[i] == nil {
			memo[i] = make([]int, len(triangle[i]))
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		memo[i][j] = min(dfs(i+1, j), dfs(i+1, j+1)) + triangle[i][j]
		return memo[i][j]
	}
	return dfs(0, 0)
}

// 方法三：动态规划，自底向上bottom-up，更易理解
// DP状态定义：dp[i][j]表示位置(i,j)走到最后一层的最小路径和
// DP方程：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + A[i][j]
// 最终结果：dp[0][0]
// 时间复杂度：O(n*n)
// 空间复杂度：O(n)
func minimumTotal3(triangle [][]int) int {
	l := len(triangle)
	// 1. 初始化dp数组，与triangle一样
	dp := make([][]int, l)
	copy(dp, triangle)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 2. 根据dp方程 dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + A[i][j] 递推
	// 从倒数第二行开始递推
	for i := l - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

// 方法四：动态规划，自顶向下top-down，不好写
// DP状态定义：dp[i][j]表示从顶走到位置(i,j)的最小路径和
// DP方程：dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + A[i][j]
// 最终结果：最后一行中的最小值
// 时间复杂度：O(n*n)
// 空间复杂度：O(n)
func minimumTotal4(triangle [][]int) int {
	l := len(triangle)
	// 1. 初始化dp数组，与triangle一样
	dp := make([][]int, l)
	copy(dp, triangle)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 自顶向下递推，DP方程：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + A[i][j]
	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 这里分为两种情况：
			// 1、上一层没有左边值
			// 2、上一层没有右边值
			if j == 0 { // 最左边
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 { // 最右边
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else { // 中间
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	ans := dp[l-1][0]
	for i := 1; i < len(dp[l-1]); i++ {
		ans = min(ans, dp[l-1][i])
	}
	return ans
}
