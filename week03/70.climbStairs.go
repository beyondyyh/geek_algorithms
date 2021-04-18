package week03

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 示例 1：
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// @leetcode: https://leetcode-cn.com/problems/climbing-stairs/

// f(x) = f(x-1) + f(x-2)
// 方法一：傻笨递归，类似斐波那契数列
// 时间复杂度：O(2^n) 结果就是超时。。。
func climbStairs1(n int) int {
	// 终结者
	if n <= 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 方法二：动态规划，或者叫递归+缓存
// 本问题其实常规解法可以分成多个子问题，爬第n阶楼梯的方法：
// 1. 从n-2阶一步爬2个台阶 或者 2. 从n-1阶爬1个台阶。那么爬n阶的走法就是这两种方法之和
// 所以问题简化为：dp[n] = dp[n-2] + dp[n-1]
// 也就是斐波那契数列
// 时间复杂度：O(n)
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
