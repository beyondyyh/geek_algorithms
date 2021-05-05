# 动态规划专项练习

## 关键点

1. 找到子问题subproblem，只考虑局部；
2. 定义DP状态数组，一维不好理解的可以升维，比如经典的 `打家劫舍、买卖股票` 等问题，可以用二维[0,1]标记当前房子偷或不偷（持股或卖出）的2种状态；
3. 定义DP转移方程，根据状态定义转移方程，如果同一时间有多个状态则需要定义多个方程；
4. 初始化DP转移方程的第一个值，也可能是二维矩阵的第一行和第一列；
5. 根据DP方程递推，一般 `自底向上bottom-up` 的思路比较好理解；`自顶向下top-down` 相当于记忆化搜索，需要处理各种边界值的情况，比如矩阵走到边了要规避数组越界等，会多出很多判断分支。

> 第2步非常关键，根据不同的状态定义写出的DP转移方程截然不同。

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)|[198.rob.go](198.rob.go)|M|DP、升维、滚动数组|
|[213. 打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/)|[213.robII.go](213.robII.go)|M|DP、滚动数组|
|[121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[121.maxProfit.go](121.maxProfit.go)|S|递归、DP、升维、滚动数组|
|[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[122.maxProfitII.go](122.maxProfitII.go)|S|递归、DP、升维、贪心算法|
|[123. 买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)|[123.maxProfitIII.go](123.maxProfitIII.go)|H|DP|
|[309. 最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)|[309.maxProfitWithCooldown.go](123.maxProfitIII.go)|M|DP|
|[188. 买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)|[188.maxProfitIV.go](188.maxProfitIV.go)|H|DP|
|[714. 买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)|[714.maxProfitWithFee.go](714.maxProfitWithFee.go)|M|DP|
