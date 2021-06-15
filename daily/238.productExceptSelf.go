package daily

// 238. 除自身以外数组的乘积
// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 示例:
// 输入: [1,2,3,4]
// 输出: [24,12,8,6]
// @lc: https://leetcode-cn.com/problems/product-of-array-except-self/

// 方法一：乘积 = 当前数左边的乘积 * 当前数右边的乘积
func productExceptSelf(nums []int) []int {
	k, n := 1, len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = k
		k *= nums[i] // 此时数组存的是除去当前元素左边的元素乘积
	}
	k = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= k  // 此处k为该数右边的乘积
		k *= nums[i] // 此时数组等于该数左边的乘积 * 该数右边的
	}
	return res
}
