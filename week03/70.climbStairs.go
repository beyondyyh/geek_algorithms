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
func climbStairs2(n int) int {
	p, q, ans := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = ans
		ans = p + q
	}
	return ans
}
