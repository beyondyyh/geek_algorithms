package week05

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 示例 1：
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶
// @leetcode: https://leetcode-cn.com/problems/climbing-stairs

// 方法一：傻递归
// 时间复杂度：O(2^b) 指数级，慢的要死
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 方法二：递归加记忆化搜索，把中间计算过f(n)cache起来
// 时间复杂度：O(n)
func climbStairs2(n int) int {
	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}

// 方法二：动态规划
// DP方程：dp[n] = dp[n-1] + dp[n-2]
// 时间复杂度：O(n) 线性级
func climbStairs3(n int) int {
	// 多初始化了个dp[0]，所以size为n+1
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
