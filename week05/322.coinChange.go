package week05

import "math"

// 322. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 你可以认为每种硬币的数量是无限的。
// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// @lc: https://leetcode-cn.com/problems/coin-change/

// 方法一：暴力递归、回溯
// 时间复杂度：O(S^n) 指数级, S是amount，n是coins长度，解释：对于每个剩余的amount都需要遍历n次
func coinChange1(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var res int = math.MaxInt32
	// 递归函数
	// - amount 剩下的金额
	// - count 	硬币个数
	var dfs func(int, int)
	dfs = func(amount, count int) {
		// 1. terminaotr
		if amount < 0 {
			return
		}
		// 2. current logic, amount=0 说明凑够了，取当前count和历史count较小者
		if amount == 0 {
			res = min(res, count)
		}
		// 3. drill down
		for _, coin := range coins {
			dfs(amount-coin, count+1)
		}
	}
	dfs(amount, 0)
	// 如果没有任何一种硬币组合能组成总金额，返回 -1
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// 方法二：动态规划，bottom-up
// 子问题：n-k[1,2,5]，n表示凑成n面值的最小硬币数，k表示面值1,2,5，子问题退化为减去k面值的硬币的个数f(n-k) + 1
// DP状态数组：f(n)表示凑成n面值的最小硬币数
// DP方程：f(n) = min(f(n-k) k为遍历coins{for k in coins} + 1)， n-k，k为遍历面值数组，
// 时间复杂度：O(Sn) S是金额amount，n是面值数coins的长度，解释：一共需要计算O(S)个状态，对于每个状态，需要枚举n个面值来转移状态，所以一共需要O(Sn)
// 空间复杂度：O(n) 解释：数组 dp 需要开长度为总金额 S 的空间
func coinChange2(coins []int, amount int) int {
	// 最大个数是全部为面值1凑成的结果也就是等于amount，所以amount+1肯定达不到
	maxNum := amount + 1
	// 定义dp状态数组
	dp := make([]int, amount+1)
	// amount为0的有0种方式
	dp[0] = 0
	// fill dp with maxNum
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 递推，遍历剩下的amount
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// i表示还需要凑的面值数，如果coin小于等于i才可以拿coin来凑，不然不就超额了么
			if coin <= i {
				// coin表示要用当前这个面值的硬币，dp[i-coin]表示减去coin剩下的amount要用多少个硬币，+1表示用了当前面coin
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 如果没有任何一种硬币组合能组成总金额，返回 -1
	if dp[amount] == maxNum {
		return -1
	}
	return dp[amount]
}
