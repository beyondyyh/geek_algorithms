package week04

// 154. 寻找旋转排序数组中的最小值 II
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
// 若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
// 给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

// 二分查找法
// 时间复杂度：平均时间复杂度为O(log(n))，最坏情况为O(n)，全是重复的元素
// 空间复杂度：O(1)
func findMinWithDup(nums []int) int {
	left, right := 0, len(nums)-1 // 左闭右闭区间，如果用右开区间则不方便判断右值
	for left < right {            // 循环不变式，如果left == right，则循环结束
		mid := left + (right-left)/2 // 向下取整，mid更靠近left，因为要找最小值，这是旋转前升序数组，所以最小值肯定是某个区间内偏向左找
		// fmt.Printf("left:%d->%d right:%d->%d mid:%d->%d\n", left, nums[left], right, nums[right], mid, nums[mid])
		if nums[mid] > nums[right] { // 中值 > 右值，最小值在右半边，收缩左边界
			left = mid + 1 // 因为中值 > 右值，中值肯定不是最小值，左边界可以跨过mid
		} else if nums[mid] < nums[right] { // 明确中值 < 右值，最小值在左半边，收缩右边界
			right = mid // 因为中值 < 右值，中值也可能是最小值，右边界只能取到mid处
		} else {
			right-- // 遇到重复元素，挪动右边界，向前收缩一位即可，
		}
	}
	return nums[left] // 循环结束，left == right，最小值输出nums[left]或nums[right]均可
}
