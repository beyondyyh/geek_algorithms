package week06

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// @lc：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

// 方法一：暴力枚举，2层遍历：i:0->n-2, j:i+1->n-1
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 有可能不发生交易，结果集的初始值设为 0
	res := 0
	// 枚举所有可能发生一次交易的股价差
	for i := 0; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

// 方法二：动态规划
// DP状态数组：当天i是否持股是很重要的一个因素，因此可以把是否持股设计到状态数组里
// 		dp[i][j] i:下标为i这天手上持有的收益，j=0:表示当前不持股，j=1:表示当前持股
// DP转移方程：
// 		当天不持股：如果前一天不持股则什么也不做，如果前一天持股则当天卖出：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// 		当天持股：如果前一天持股则什么也不做，如果前一天不持股则当天买入，即金额为当天股价的相反数： dp[i][1] = max(dp[i-1][1], -prices[i])
// 最终结果：dp[n-1][0] 卖出状态
// 时间复杂度：O(n)
// 空间复杂的：O(n)
func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 初始化dp状态数组
	dp := make([][2]int, n)
	// 不持股显然为0，持股则需要减去第1天（下标为 0）的股价
	dp[0] = [2]int{0, -prices[0]}
	// 从第2天开始，dp递推
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// 方法三：动态规划，滚动数组优化空间复杂度为O(1)
// DP状态数组：dp[i][j]表示当天持股或不持股的最大收益
// DP转移方程：
//	当天不持股：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
//  当天持股：  dp[i][i] = max(dp[i-1][1], -prices[i])
// 观察得出第i天的状态只与前一天有关
// 因此用一个长度为2的数组，表示当天持股或不持股的最大收益即可，滚动迭代
// 时间复杂度：O(n)
// 空间复杂的：O(1)
func maxProfit3(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 定义dp数组
	dp := [2]int{0, -prices[0]}
	// 递推
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}
	return dp[0]
}
