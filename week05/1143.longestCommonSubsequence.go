package week05

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
// 示例 1：
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// @leetcode: https://leetcode-cn.com/problems/longest-common-subsequence

// 状态定义：dp[i][j] 最大子序和
// DP方程：if最后一个字符相同：dp[i][j] = dp[i-1][j-1]+1; else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 1. 初始化(m+1)*(n+1)的二维数组
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 2. 根据DP方程递推
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] { // 最后一个元素相同，则问题退化为都去掉最后一个元素的最长公共子序列+1
				dp[i][j] = dp[i-1][j-1] + 1
			} else { // 否则就是text1去掉最后一个元素和text2，以及text2去掉最后一个元素和text1中二者较大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
