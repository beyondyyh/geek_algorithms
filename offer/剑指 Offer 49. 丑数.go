package offer

// 剑指 Offer 49. 丑数
// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
// 示例:
// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
// @lc: https://leetcode-cn.com/problems/chou-shu-lcof/

// 动态规划
// dp状态定义：dp[i]代表第i+1个丑数，初始状态：dp[0] = 1 为第一个丑数为
// dp转移方程：dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
func nthUglyNumber(n int) int {
	var a, b, c int // default is 0
	dp := make([]int, n)
	dp[0] = 1 // 第一个丑数
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}
