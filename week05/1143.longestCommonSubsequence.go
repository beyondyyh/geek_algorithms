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

// 方法一：动态规划bottom-up，二维数组
// DP状态定义：dp[i][j]表示text1下标[0..i-1]与text2下标[0..j-1]的最长公共子序列
// DP转移方程：
//		dp[i][j] = dp[i-1][j-1]+1	如果text1和text2最后一个字符相同，则同时去掉最后一个元素求得二者的最长公共子序列，从上一个状态转移过来
//		dp[i][j] = max(dp[i-1][j], dp[i][j-1]) 如果text1和text2最后一个字符不相同，则分别去掉最后一个字符，与对方求最长公共子序列
// 最终结果：dp[m][n]
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func longestCommonSubsequence1(text1 string, text2 string) int {
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
	// 2. 初始化dp数组，长度为0的字串公共序列肯定是0
	for i := 0; i < m+1; i++ {
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

// 方法二：动态规划，滚动数组优化空间
// 由方法一观察得出，dp[i][j]的状态只由dp[i-1][j-1]（左上对角）、dp[i-1][j]（正上）和dp[i][j-1]（正左）转移而来，因此可以用一位数组迭代
// 由于dp[i-1][j]（正上）和dp[i][j-1]（正左）可以直接使用，需要注意的是左上对角，在计算前一列的时候会被覆盖，因此需要在被覆盖之前保存下旧值
// 时间复杂度：O(m*n)
// 空间复杂度：O(n)
func longestCommonSubsequence2(text1, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	// 定义dp数组，并初始化
	dp := make([]int, n+1)
	for i := 1; i < m+1; i++ {
		upLeft := dp[0] // 每行开始的时候需要更新下upleft, 这里其实每次都是0
		for j := 1; j < n+1; j++ {
			tmp := dp[j] // 记录未被覆盖之前的dp[j], 它会在计算 j+1的时候作为upLeft用到
			if text1[i-1] == text2[j-1] {
				dp[j] = upLeft + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			upLeft = tmp // 更新upLeft
		}
	}
	return dp[n]
}
