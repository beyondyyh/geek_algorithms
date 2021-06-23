package week07

// 33. 搜索旋转排序数组
// 整数数组 nums 按升序排列，数组中的值互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你旋转后的数组 nums 和一个整数target，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 示例 1：
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4

// 二分查找法
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if nums[left] <= nums[mid] { // 前半段有序，此处用小于等于，为了最后只剩两个数的时候，跟上面逻辑匹配
			if nums[left] <= target && target < nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 后半段有序
			if nums[right] >= target && target > nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
