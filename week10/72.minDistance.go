package week10

// 72. 编辑距离
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// @lc: https://leetcode-cn.com/problems/edit-distance/

// 动态规划，二维数组bottom-up
// dp状态定义：dp[i][j]表示 word1[0:i] 转换成 word2[0:j] 的最小步数
// dp转移方程：
//	当word1[i-1] == word2[j-1],	dp[i][j] = dp[i-1][j-1]
//	当word1[i-1] != word2[j-1], dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])+1
// dp初始化：第一列 dp[i][0] = i, 第一行 dp[0][j] = j
// 最终结果：dp[m][n]
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func minDistance(word1 string, word2 string) int {
	// word1作为列，word2作为行
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // m+1行
	// 初始化行赋为0，且初始化第一列
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	// 初始化第一行
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	// dp方程递推
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 可以理解为：
				// dp[i-1][j-1]的基础上replace
				// dp[i-1][j] word1的基础上insert
				// dp[i][j-1] word2的基础上delete
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
