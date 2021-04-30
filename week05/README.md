# 动态规划

## 动态规划、递归、分治的区别

- **关键点：** 动态规划 和 递归、或 分治没有根本上的区别，关键看有无最优子结构
- **共性：** 找到最近重复子问题
- **差异性：** 最优子结构、中途可以淘汰次优解

## 动态规划的关键点

- 最优子结构 dp[n] = best_of(dp[n-1], dp[n-2])，第n步的值是前面几个值的最佳值（可能是累加、最大、最小等）
- 存储中间状态 dp[i]
- 递推公式（美其名曰：状态转移方程或DP方程）
  - Fibnacci：dp[n] = dp[n-1] + dp[n-2]
  - 二维路径：dp[i][j] = dp[i-1][j] + dp[i][j-1]（且判断dp[i][j]是否为合法）

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)|[70.climbStairs.go](70.climbStairs.go)|S|DP|
|[不同路径](https://leetcode-cn.com/problems/unique-paths/)|[62.uniquePaths.go](62.uniquePaths.go)|M|DP|
|[不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)|[63.uniquePathsWithObstacles.go](63.uniquePathsWithObstacles.go)|M|DP|
|[1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)|[1143.longestCommonSubsequence.go](1143.longestCommonSubsequence.go)|M|DP|
|[剑指 Offer 49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)|[49.nthUglyNumber.go](49.nthUglyNumber.go)|M|DP|
