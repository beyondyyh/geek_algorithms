package practice

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
// 二分查找法
func findMin(nums []int) int {
	left, right := 0, len(nums)-1 // 左闭右闭区间，如果用右开区间则不方便判断右值
	for left < right {            // 循环不变式，如果left == right，则循环结束
		mid := left + (right-left)/2 // 地板除，mid更靠近left
		if nums[mid] > nums[right] { // 中值 > 右值，最小值在右半边，收缩左边界
			left = mid + 1 // 因为中值 > 右值，中值肯定不是最小值，左边界可以跨过mid
		} else if nums[mid] < nums[right] { // 明确中值 < 右值，最小值在左半边，收缩右边界
			right = mid // 因为中值 < 右值，中值也可能是最小值，右边界只能取到mid处
		}
	}
	return nums[left] // 循环结束，left == right，最小值输出nums[left]或nums[right]均可
}

// @lc code=end
