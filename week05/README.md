# 动态规划

## 动态规划、递归、分治的区别

- **关键点：** 动态规划 和 递归、或 分治没有根本上的区别，关键看有无最优子结构
- **共性：** 找到最近重复子问题
- **差异性：** 最优子结构、中途可以淘汰次优解（必须淘汰次优解）时间复杂度从指数级（泛型递归）降到多项式级（O(n)或O(n^2)）

## 动态规划的关键点

- 最优子结构 dp[n] = best_of(dp[n-1], dp[n-2])，第n步的值是前面几个值的最佳值（可能是累加、最大、最小等）
- 存储中间状态 dp[i]
- 递推公式（美其名曰：状态转移方程或DP方程）
  - Fibnacci：dp[n] = dp[n-1] + dp[n-2]
  - 二维路径
    - 自底向上Bottom-up，状态定义：dp[i][j]表示从位置i,j走到“End”的路径数；**DP方程：** dp[i][j] = dp[i+1][j] + dp[i][j+1]（且判断dp[i][j]是否为合法）
    - 自顶向下Top-bottom，状态定义：dp[i][j]表示从“Start”走到位置i,j的路径数；**DP方程：** dp[i][j] = dp[i-1][j] + dp[i][j-1]

**动态规划小结：**

- 打破自己的思维惯性，形成机器思维，计算机只会傻傻的if-else, for-loop, recursion，`自顶向下搜索其实是：递归+记忆化搜索，结果在最后一位；自底向上搜索是：递推公式计算`，一般更好理解。
- 理解复杂逻辑的关键，首先定义好DP状态，和DP方程
- 配合大量刷题，总结经验

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)|[70.climbStairs.go](70.climbStairs.go)|S|DP|
|[62. 不同路径](https://leetcode-cn.com/problems/unique-paths/)|[62.uniquePaths.go](62.uniquePaths.go)|M|DP|
|[63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)|[63.uniquePathsWithObstacles.go](63.uniquePathsWithObstacles.go)|M|DP|
|[64. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)|[64.minPathSum.go](64.minPathSum.go)|M|DP|
|[1143. 最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)|[1143.longestCommonSubsequence.go](1143.longestCommonSubsequence.go)|M|DP|
|[剑指 Offer 49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)|[49.nthUglyNumber.go](49.nthUglyNumber.go)|M|DP|
|[120. 三角形最小路径和](https://leetcode-cn.com/problems/triangle/)|[120.minimumTotal.go](120.minimumTotal.go)|M|DP|
|[53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)|[53.maxSubArray.go](53.maxSubArray.go)|S|DP|
