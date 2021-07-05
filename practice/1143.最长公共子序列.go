package practice

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
// 方法一：动态规划，bottom-up，二维数组
// dp状态定义：dp[i][j] i:text1的下标0..n，j:text2的下标0..n，dp[i][j]表示text1下标0..i和text2下标0..j的最长公共子序列
// dp转移方程：
//      dp[i][j] = dp[i-1][j-1]+1 如果text1和text2最后一个字符相同，则继承前一个状态并加1
//      dp[i][j] = max(dp[i-1][j], dp[i][j-1]) 如果text1和text2最后一个字符不相同，则分别去掉最后一个与对方求最长子序列
// 最终结果：dp[m][n]
// 时间复杂度：O(m*n)
// 空间复杂的：O(m*n)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m*n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 初始化
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 递推
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
