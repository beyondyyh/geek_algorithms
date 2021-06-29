package practice

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
// 动态规划
// dp状态定义：dp[i][j]表示 word1[0..i) 和 word2[0..j) 的编辑距离
// dp转移方程：
// 	word1[i-1] == word2[j-1]: dp[i][j] = dp[i-1][j-1]
//	word1[i-1] != word2[j-1]: dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])+1 word1去掉最后一个字符 和 word2去掉最后一个字符 的编辑距离
// dp初始化：dp[i][0] = i, dp[0][j] = j
// 最终结果：dp[m][n]
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 特殊case处理，有一个为空串
	if m*n == 0 {
		return m + n
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// dp定义
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// dp初始化
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	// 递推
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left, down, leftDown := dp[i][j-1], dp[i-1][j], dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = leftDown
			} else {
				dp[i][j] = min(min(left, down), leftDown) + 1
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
