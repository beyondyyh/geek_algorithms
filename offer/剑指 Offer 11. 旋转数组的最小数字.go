package offer

// 剑指 Offer 11. 旋转数组的最小数字
// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
// 示例 1：
// 输入：[3,4,5,1,2]
// 输出：1
// 示例 2：
// 输入：[2,2,2,0,1]
// 输出：0
// @lc: https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

// 二分查找法，寻找旋转点，左有序数组和右有序数组，那么右有序数组的第一个元素就是最小值了
// 时间复杂度：O(log(n))
// 空间复杂度：O(1)
func minArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		// 右半部分有序，则旋转点一定在左边，故收缩右边界
		if nums[mid] < nums[right] {
			// 由于nums[mid]比较小不确定是不是最小的，所以不能丢弃mid，因此右边界收缩为：mid，包含mid
			right = mid
		} else if nums[mid] > nums[right] {
			// 说明旋转点一定在右边，则左半部分有序，故收缩左边界
			// 由于nums[mid]比较大一定不是最小的，所以可以丢弃mid，因此左边界收缩为：mid+1，不包含mid
			left = mid + 1
		} else {
			// nums[mid] = nums[right]，此时可以暴力收缩右边界，而且不会有漏掉的情况，right--
			right--
		}
	}
	// 退出循环时，left=right
	return nums[left]
}
