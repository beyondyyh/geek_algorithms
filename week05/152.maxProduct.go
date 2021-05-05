package week05

import "math"

// 152. 乘积最大子数组
// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// 示例 1:
// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// @lc: https://leetcode-cn.com/problems/maximum-product-subarray/

// 动态规划，bottom-up
// DP状态定义：imax:当前阶段最大值，imin:当前阶段最小值，由于存在负数且负负得正，因此还需要维护当前阶段最小值imin
// DP方程：imax = max(imax*nums[i], nums[i]), imin = min(imin*nums[i], nums[i])
// 关键点：当前为负数时，会使最大变成最小，最小变成最大，因此出现负数时将imax和imin进行交换后再进行下一步计算
// 空间复杂度：O(n)
// 时间复杂度：O(1)
func maxProduct(nums []int) int {
	imax, imin, res := 1, 1, math.MinInt32
	for _, num := range nums {
		if num < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax*num, num)
		imin = min(imin*num, num)
		// 更新res
		res = max(res, imax)
	}
	if res == math.MinInt32 {
		return 0
	}
	return res
}
