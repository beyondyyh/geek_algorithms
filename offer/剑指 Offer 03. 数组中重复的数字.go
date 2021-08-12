package offer

// 剑指 Offer 03. 数组中重复的数字
// 找出数组中重复的数字。
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// 示例 1：
// 输入：
// [2, 3, 1, 0, 2, 5, 3]
// 输出：2 或 3
// @lc: https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

// 方法一：hashmap
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = struct{}{}
	}
	return -1
}

// 原地交换
// 1. 遍历数组 numsnums ，设索引初始值为 i = 0 :
//     若 nums[i] = i：说明此数字已在对应索引位置，无需交换，因此跳过；
//     若 nums[nums[i]] = nums[i]：代表索引 nums[i] 处和索引 i 处的元素值都为 nums[i] ，即找到一组重复值，返回此值 nums[i]；
//     否则：交换索引为 i 和 nums[i] 的元素值，将此数字交换至对应索引位置。
// 2. 若遍历完毕尚未返回，则返回 -1
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func findRepeatNumber2(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
