package week06

// 122. 买卖股票的最佳时机 II
// 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 示例 1:
// 输入: prices = [7,1,5,3,6,4]
// 输出: 7
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//      随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
// @lc: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

// 方法一：暴力检索、递归
// 时间复杂度：O(2^n)
// 空间复杂度：O(1)
func maxProfitII1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	var res int
	// dfs
	// - index 	当前第几天，从0开始
	// - status	当前天的状态 0:不持股，1:持股
	// - profit 当前天所持有的收益
	var dfs func(int, int, int)
	dfs = func(index, status, profit int) {
		// fmt.Printf("index:%d status:%d profit:%d res:%d\n", index, status, profit, res)
		// terminator
		if index == n {
			res = max(res, profit)
			return
		}

		// process current logic & drill down
		// a. 当天啥玩意儿也不操作，继承前一天的状态
		dfs(index+1, status, profit)
		if status == 0 {
			// b. 当天不持股，可以考虑将status转向1，也就是第二天买入股票，收益要减去prices[i]
			dfs(index+1, 1, profit-prices[index])
		} else {
			// c. 当天持股，可以考虑将status转向0，也就是第二天卖出股票，获得收益
			dfs(index+1, 0, profit+prices[index])
		}
	}
	// 第1天 index=0, 不持有股票，收益为0
	dfs(0, 0, 0)
	return res
}

// 方法二：动态规划、升维
// DP状态数组：dp[i][j] i:第i天，j->0,1 0:不持股，1:持股
// DP转移方程：
//	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])，第i天不持股，若前一天不持股则直接转移 或 前一天持股则第今天卖出获得收益
//	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])，第i天持股，若前一天持股则直接转移 或 前一天不持股则今天买入
// 结果：dp[n-1][0]，最后一天卖出股票的收益肯定比持股的收益大
// 时间复杂度：O(n) 遍历一次
// 空间复杂度：O(n) 虽然是二维数组，但是第二维是常数，与问题规模无关
func maxProfitII2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// 初始化dp状态数组
	dp := make([][2]int, n)
	// 第一天 不持股:收益0，持股:负收益
	dp[0] = [2]int{0, -prices[0]}
	// 根据dp方程递推
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 方法三：动态规划，空间优化
// 从方法二观察得出，第i天的状态只与前一天有关，因此可以用滚动数组思想优化存储空间
// 时间复杂度：O(n) 遍历一次
// 空间复杂度：O(1) 2个变量是常数，与问题规模无关
func maxProfitII3(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// cash：表示卖出股票持有的收益也就是不持有股票，hold：表示持有股票
	// 第1天的状态
	cash := 0
	hold := -prices[0]
	preCash, preHold := cash, hold
	for i := 1; i < n; i++ {
		cash = max(preCash, preHold+prices[i]) // 今天不持有股票：前一天不持有股票 或 前一天持有股票今天卖出
		hold = max(preHold, preCash-prices[i]) // 今天持有股票：  前一天持股 或 前一天不持股今天买入
		// 滚动迭代
		preCash, preHold = cash, hold
	}
	// 最后一天一定是卖出股票比持有的收益大
	return cash
}

// 方法四：贪心算法
// 关键点：由于不限制交易次数，只要今天股价比昨天高（今日-昨日>0），就交易，收益叠加
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfitII4(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	res := 0
	for i := 1; i < n; i++ {
		res += max(prices[i]-prices[i-1], 0)
	}
	return res
}
