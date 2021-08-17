package offer

// 剑指 Offer 10- II. 青蛙跳台阶问题
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
// 示例 1：
// 输入：n = 2
// 输出：2

// 示例 2：
// 输入：n = 7
// 输出：21

// 示例 3：
// 输入：n = 0
// 输出：1
// @lc: https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

// 动态规划
// dp状态定义：dp[i]表示跳上i级台阶总共的跳法
// dp转移方程：dp[i] = dp[i-1] + dp[i-2]
// 初始化：
// 		青蛙跳台阶问题： f(0)=1, f(1)=1, f(2)=2；
// 		斐波那契数列问题： f(0)=0, f(1)=1, f(2)=1；
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= (1e9 + 7)
	}
	return dp[n]
}

// 滚动数组
func numWays2(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}
