package daily

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

// 方法一：动态规划，bottom-up
// dp状态定义：dp[i][j]表示从(i,j)位置走到底部的最小路径和
// dp转移方程：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]，向左下或右下走
// dp初始化：等于triangle
// 最终结果：dp[0][0]
// 时间复杂度：O(n*n)
// 空间复杂度：O(n)
func minimumTotal1(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	copy(dp, triangle) // 拷贝triangle的值，要不就是复用triangle，会改变他的值
	// 递推：i->m-2..0, j->0..n-1 最后一行达到最下的值是它本身，所以从m-2开始递推
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

// 方法二：动态规划，top-down
// dp状态定义：dp[i][j]表示从塔顶走到当前位置(i,j)的最小路径和
// dp转移方程：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
// dp初始化：等于triangle
// 最终结果：最后一行的最小值
// 时间复杂度：O(n*n)
// 空间复杂度：O(n)
func minimumTotal2(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	copy(dp, triangle)
	// 递推：i->1..m-1, j->0..n-1 第一行的值是它本身
	for i := 1; i < m-1; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 这里分为两种情况：
			// 1、上一层没有左边值
			// 2、上一层没有右边值
			if j == 0 { // case 1
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i]) { // case 2
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	// 结果
	res := dp[m-1][0]
	for i := 1; i < len(dp[m-1]); i++ {
		res = min(res, dp[m-1][i])
	}
	return res
}
