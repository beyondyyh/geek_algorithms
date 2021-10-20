package offerII

// 剑指 Offer II 009. 乘积小于 K 的子数组
// 给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
// 示例 1:
// 输入: nums = [10,5,2,6], k = 100
// 输出: 8
// 解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
// 需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
// 示例 2:
// 输入: nums = [1,2,3], k = 0
// 输出: 0

// 窗口每次移动后，res都可以增加 right - left + 1 个子数组。
func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, right, res int
	total := 1
	for ; right < len(nums); right++ {
		total *= nums[right]
		for left <= right && total >= k {
			total /= nums[left]
			left++
		}
		if left <= right {
			res += right - left + 1
		}
	}
	return res
}
