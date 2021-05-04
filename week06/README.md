# 动态规划专项练习

## 动态规划关键点

- 找到子问题subproblem
- 定义dp状态数组，一维不好理解可以升维，比如‘打家劫舍’，可以二维(0,1)标记当前第i个房子偷或不偷2种状态
- 定义dp转移方程
- 根据dp方程递推，一般自底向上bottom-up比较好理解

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/)|[198.rob.go](198.rob.go)|M|DP|
|[213. 打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/)|[213.robII.go](213.robII.go)|M|DP|
|[121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)|[121.maxProfit.go](121.maxProfit.go)|S|DP|
