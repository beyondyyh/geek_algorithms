package week07

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// @lc：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

// 动态规划，bottom-up，二维数组
// dp状态数组：dp[i][j]表示第i天拥有的收益，j有2中状态，j=0:第i天不持股，j=1:第i天持有股票
// dp转移方程：只能买卖一次。。。
//	当天不持股：可以由前一天不持股 或 持股今天卖出转移过来，dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
// 	当天持股：可以由前一天持股 或 不持股今天买入转移过来，dp[i][1] = max(dp[i-1][1], -prices[i])
// 最终结果：最后一天不持股的收益最大，dp[n-1][0]
// 时间复杂度：O(n)
// 空间复杂度：O(n) 虽然是2位数组但是第二维与数据规模无关，故是O(n)
func maxProfit1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// 定义dp状态数组，初始化
	dp := make([][2]int, n)
	dp[0] = [2]int{0, -prices[0]}
	// 由dp方程递推
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// 动态规划，一维数组，优化内存空间
// dp状态数组：dp[j]表示当天持股或不持股的收益，后一天覆盖前一天的数据
// dp转移方程：
// 	当天不持股：前一天不持股或持股今天卖出，dp[0] = max(dp[0], dp[1]+prices[i])
//	当天持股：前一天持股或不持股今天买入，dp[1] = max(dp[1], -prices[i])
// 最终结果：dp[0]
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := [2]int{0, -prices[0]}
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}
	return dp[0]
}
