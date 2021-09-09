package offer

// 剑指 Offer 14- I. 剪绳子
// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
// 每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
// 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
// 示例 1：
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1

// 示例 2:
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
// @lc: https://leetcode-cn.com/problems/jian-sheng-zi-lcof/

// dp状态定义：dp[i]表示长度为i的绳子剪成m段之后的最大乘积
// dp转移方程：dp[i] = max(dp[i], max(j * (i-j), j * dp[i-j]))
//	 	括号里的dp[i]: 表示维持原状，不剪；
//	 	j * (j-i):    表示从j处剪一下，剪下来的部分是i-j，i-j这根不再剪了，的乘积；
//		j * dp[j-i]:  表示从j处剪一下，剪下来的部分是i-j，i-j这根继续剪，的乘积。
//		由于从 j=1 处开始剪切的话对乘积没有增益，所以从 j=2 处开始剪
// dp初始化：dp[2] = 1：长度为2的绳子剪2下乘积为1*1=1，dp[0],dp[1]都是0，题目要求 n>1 且剪切次数 m>1
// 最终结果：dp[n]
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
