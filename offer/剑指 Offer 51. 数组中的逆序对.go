package offer

// 剑指 Offer 51. 数组中的逆序对
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
// 示例 1:
// 输入: [7,5,6,4]
// 输出: 5
// @lc: https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 归并排序的思想，在merge的过程中计算逆序对的个数
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

// 在merge的过程中计算逆序对的个数
// 时间复杂度：O(nlogn) 没有增加归并排序的时间复杂度
func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	// 一分为二
	mid := (left + right) >> 1
	// 左边数组逆序对的个数 + 右边数组逆序对的个数
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	// 中间数组，存放 [left, right] 区间合并后的有序数组
	temp := make([]int, 0, right-left+1)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			count += j - (mid + 1)
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	// 处理 [left, mid] 未遍历完的情况，
	for ; i <= mid; i++ {
		temp = append(temp, nums[i])
		count += right - (mid + 1) + 1
	}
	// 处理 [mid+1, right] 未遍历完的情况
	for ; j <= right; j++ {
		temp = append(temp, nums[j])
	}
	// 将 有序数组temp 合并回nums[left, right]区间
	for k := 0; k < len(temp); k++ {
		nums[left+k] = temp[k]
	}
	return count
}
