package week10

// 1143. 最长公共子序列
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
// 示例 1：
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// @lc: https://leetcode-cn.com/problems/longest-common-subsequence/

// 动态规划
// dp状态定义：dp[i][j] 表示text1[0:i-1]长度为i的字串 和 text2[0:j-1]长度为j字串 的最长公共子序列
// dp转移方程：
//	text1[i-1] == text2[j-1]: dp[i][j] = dp[i-1][j-1]+1 最后一个字符相同，公共子序列的长度可以有前一个状态转移过来并且+1
//	text1[i-1] != text2[j-1]: dp[i][j] = max(dp[i-1][j],dp[i][j-1]) 最后一个字符不相同，那么可以由text1去掉一个字符或text2去掉一个字符，二者较大值
// dp初始化：dp[0][0] = 0
//	当i=0时，dp[0][j] 表示的是text1取空字符串 跟 text2的最长公共子序列，结果为0
//  当j=0时，dp[i][0] 表示的是text2取空字符串 跟 text1的最长公共子序列，结果为0
// 最终结果：dp[m][n] dp[i][j]表示的是text1[0:i-1]和text2[0:j-1]的最长公共子序列，因此dp[m][n]表示text1和text2的最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
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
	// 状态递推
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
