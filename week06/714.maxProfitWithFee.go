package week06

// 714. 买卖股票的最佳时机含手续费
// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
// 示例 1:
// 输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
// 输出: 8
// 解释: 能够达到的最大利润:
// 在此处买入 prices[0] = 1
// 在此处卖出 prices[3] = 8
// 在此处买入 prices[4] = 4
// 在此处卖出 prices[5] = 9
// 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
// @lc: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// 方法一：动态规划，二维数组
// dp状态定义：dp[i][j]
//	j=0: 第i天不持股获得的收益
//	j=1: 第i天不持股获得的收益，注意不是第i天卖出
// dp转移方程：
//	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) 第i天不持股可以有2种状态转移过来，1:昨天不持股，2:昨天持股今天卖出。二者取较大
//	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) 第i天持股可以有2种状态转移过来，1:昨天持股，2:昨天不持股，今天买入。二者取较大
// 最终结果：dp[n-1][0]
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func maxProfitWithFee1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 定义dp状态数组
	dp := make([][2]int, n)
	// 初始化，第一天不持股：收益0，持股：收益为负
	dp[0] = [2]int{0, -prices[0]}
	// 根据dp方程递推
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 方法二：动态规划，一位数组，空间优化
// dp状态定义：dp[j] j->0,1
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfitWithFee2(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 定义dp数组，并初始化
	dp := [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		tmp := dp[0] // 不持有
		dp[0] = max(dp[0], dp[1]+prices[i]-fee)
		dp[1] = max(dp[1], tmp-prices[i])
	}
	return dp[0]
}
