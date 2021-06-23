/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
// 二分查找 的变种
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			return mid
		}
		// 左边部分有序，此处用小于等于，为了最后只剩两个数的时候，跟上面逻辑匹配
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边部分有序
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end