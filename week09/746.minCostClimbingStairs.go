package week09

// 746. 使用最小花费爬楼梯
// 数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
// 每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
// 请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。
// 示例 1：
// 输入：cost = [10, 15, 20]
// 输出：15
// 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
// @lc: https://leetcode-cn.com/problems/min-cost-climbing-stairs/

// 方法一：动态规划
// 1. dp状态定义：dp[i]表示到索引为i的位置，再选择向上爬一共需要的体力开销
// 2. dp转移方程：dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
// 3. dp初始化：dp[0] = dp[1] = 0
// 4. 最终结果：dp[n]
// 注意：楼层顶部在数组之外。设数组长度为 n，那么楼顶就在下标为 n，dp 数组开 n+1 个空间。
// 空间复杂度：O(n)
// 时间复杂度：O(n)
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// 方法二：滚动迭代
// 空间复杂度：O(n)
// 时间复杂度：O(1)
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}
