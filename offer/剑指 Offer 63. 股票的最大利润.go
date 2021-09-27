package offer

import "math"

// 剑指 Offer 63. 股票的最大利润
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
// 示例 1:
// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
// 示例 2:
// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
// @lc: https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

// 计算前 i 日的最低价格 min(prices[0:i]) 时间复杂度为 O(i)。而在遍历 prices 时，可以借助一个变量（记为成本 cost ）每日更新最低价格。
// 优化后的转移方程为： dp[i]=max(dp[i−1], prices[i]−min(cost, prices[i])
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit(prices []int) int {
	cost, res := math.MaxInt64, 0
	for _, price := range prices {
		cost = min(int(cost), price)
		res = max(res, price-cost)
	}
	return res
}

// 在遍历数组的过程中，维护一个最小值，最小值初试为prices[0]
//     如果prices[i]大于min，则去更新一下利润res
//     否则说明当前的prices[i]比min还小，则更新min
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	res, minv := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] <= minv {
			minv = prices[i]
		} else {
			res = max(res, prices[i])
		}
	}
	return res
}
