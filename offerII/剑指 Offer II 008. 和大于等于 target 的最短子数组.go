package offerII

import "math"

// 剑指 Offer II 008. 和大于等于 target 的最短子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
// 示例 1：
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

// 数组题目多用双指针，一般双指针有3种常见：
// 1. 首尾指针：范围查找，比如二分查找等；
// 2. 滑动窗口：指针处在数组同一方向，根据条件移动左右指针，用于获取范围和等
// 3. 快慢指针：多用于链表计算时，判断是否有环等

// 滑动窗口解题模板：
// 初始化左边界 left = 0
// 初始化返回值 ret = 最小值 or 最大值
// for 右边界 in 可迭代对象:
// 	更新窗口内部信息
// 	while 根据题意进行调整：
// 		比较并更新ret(收缩场景时)
// 		扩张或收缩窗口大小
// 	比较并更新ret(扩张场景时)
// 返回 ret

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	res := math.MaxInt32
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}
