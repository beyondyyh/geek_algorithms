package week06

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
// 示例 1：
// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// @lc：https://leetcode-cn.com/problems/house-robber-ii

// 由于第一个房子和最后一个房子是相邻的，只能偷一个，所以问题可退化为2个198的单排打家劫舍的问题：
//  1.第1个房子不偷，nums[1:]的最大值p1
//  2.最后一个房子不偷，nums[0:n-1]的最大值p2
// 结果：max(p1,p2)
// 时间复杂度：O(n)	两次遍历nums需要线性时间；
// 空间复杂度：O(1) cur和 pre 使用常数大小的额外空间
func robII1(nums []int) int {
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
	// 正常的打家劫舍，类似198的题解
	myRob := func(nums []int) int {
		pre, cur := 0, 0
		for _, i := range nums {
			pre, cur = cur, max(cur, pre+i)
		}
		return cur
	}
	// 第1个房子不偷，nums[1:]的最大值p1
	p1, p2 := myRob(nums[1:]), myRob(nums[:n-1])
	return max(p1, p2)
}

// 方法二：动态规划，一维数组
// DP状态数组：dp[i]表示从0..i所有情况下能偷到的最大值，第i个房子可偷可不偷
// DP转移方程：dp[i] = max(dp[i-1]+0, dp[i-2]+nums[i])，
func robII2(nums []int) int {
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
	myRob := func(nums []int) int {
		n := len(nums)
		dp := make([]int, n)
		dp[0], dp[1] = nums[0], max(nums[0], nums[1])
		for i := 2; i < n; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+i)
		}
		return dp[n-1]
	}
	return max(myRob(nums[1:]), myRob(nums[:n-1]))
}
