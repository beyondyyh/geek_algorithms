package offer

// 剑指 Offer 57. 和为s的两个数字
// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[2,7] 或者 [7,2]

// 示例 2：
// 输入：nums = [10,26,30,31,47,60], target = 40
// 输出：[10,30] 或者 [30,10]
// @lc: https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

// 两数之和可以用hashmap，加1次遍历，时间和空间复杂度均为：O(n)
// 此题给出的条件是输入一个递增排序的数组，所以可以用双指针两端加逼将空间复杂度将为O(1)
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if target == sum {
			return []int{nums[left], nums[right]}
		} else if target > sum {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
