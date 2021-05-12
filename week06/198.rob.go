package week06

// 198. 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 示例 1：
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。

// 方法一：动态规划，升维思想
// 子问题：当前第i个房子偷或不偷
// DP状态数组：dp[i][0,1] 0:第i个房子不偷，1:第i个房子偷
// DP转移方程：
//	第i个房子不偷，那么第i-1个房子可偷可不偷，返回二者较大值：dp[i][0] = max(dp[i-1][0], dp[i][1])
//	第i个房子偷，那么第i-i个房子就不能偷了，返回第i-1个房子不偷的金额加上第i个房子的金额：dp[i][1] = dp[i-1][0] + nums[i]
// 时间复杂度：O(n)
// 空间复杂度：O(2n)
func rob1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(nums)
	// 定义dp状态数组
	dp := make([][2]int, n)
	dp[0] = [2]int{0, nums[0]}
	// 根据dp方程递推
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

// 方法二：动态规划，一维数组
// DP状态数组：dp[i]表示从0..i所有情况下能偷到的最大值，第i个房子可偷可不偷
// DP转移方程：dp[i] = max(dp[i-1]+0, dp[i-2]+nums[i])，
//	dp[i-1]+0：表示第i个房子不偷，那么可以直接从第i-1个房子的最大值赋过来
// 	dp[i-2]+nums[i]：表示第i个房子偷，那么就不能从i-1的最大值赋过来了，而是从第i-2个房子的最大值赋过来，同时加上第i个房子的金额，
//	以上2种求最大值，就是第0..i个房子的最大值
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 定义dp状态数组
	dp := make([]int, n)
	// dp[i]存的是0..i的最大值，dp[0]：第一个房子最大值，那肯定是nums[0]；dp[1]：第1和2房子是相邻的只能偷一个，所以偷二者较大值
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// 方法三：动态规划，优化版，滚动数组
// 观察发现：dp[i] 只与 dp[i-1] 和 dp[i-2] 有关系，因此可以设两个变量 cur和 pre 交替记录，将空间复杂度降到 O(1)
// DP方程：dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func rob3(nums []int) int {
	pre, cur := 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, i := range nums {
		// DP方程：dp[i] = max(dp[i-1]+0, dp[i-2]+nums[i])
		pre, cur = cur, max(cur, pre+i)
	}
	return cur
}
