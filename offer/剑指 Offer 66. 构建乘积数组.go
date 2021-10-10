package offer

// 剑指 Offer 66. 构建乘积数组
// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

// 示例:
// 输入: [1,2,3,4,5]
// 输出: [120,60,40,30,24]
// @lc: https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/

func constructArr(nums []int) []int {
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
