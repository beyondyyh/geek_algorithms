package practice

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstPos := findFirstPosition(nums, target)
	if firstPos == -1 {
		return []int{-1, -1}
	}
	lastPos := findLastPosition(nums, target)
	return []int{firstPos, lastPos}
}

func findFirstPosition(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			// ① 不可以直接返回，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		} else if target > nums[mid] {
			// 应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		} else {
			// 此时 nums[mid] > target，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		}
	}
	// 退出循环时，left 和 right 的位置关系是 [right, left]，注意上面的 ①，此时 left 才是第 1 次元素出现的位置
	// 因此还需要特别做一次判断
	if left != n && target == nums[left] {
		return left
	}
	return -1
}

func findLastPosition(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			// 只有这里不一样：不可以直接返回，应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		} else if target > nums[mid] {
			// 应该继续向右边找，即 [mid + 1, right] 区间里找
			left = mid + 1
		} else {
			// 此时 nums[mid] > target，应该继续向左边找，即 [left, mid - 1] 区间里找
			right = mid - 1
		}
	}
	// 由于 findFirstPosition 方法可以返回是否找到，这里无需单独再做判断
	return right
}

// @lc code=end
