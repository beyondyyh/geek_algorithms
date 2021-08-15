package practice

/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
// 归并排序的思想，在merge的过程中统计翻转对的个数
func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSortHelper(nums, 0, len(nums)-1)
}

// 归并排序。稳定排序算法
// 思路：分治思想的典型应用，利用递归不断将数组一分为二，直到只剩下一个元素，一个元素的数组是天然有序的，然后合并2个有序数组的过程称为：2-路归并。
// 时间复杂度：O(nlog(n))，稳定排序
// 空间复杂度：O(n)，2-路归并的额外空间
func mergeSortHelper(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	// 一分为二
	mid := (left + right) >> 1
	// 左边数组逆序对的个数 + 右边数组逆序对的个数
	count := mergeSortHelper(nums, left, mid) + mergeSortHelper(nums, mid+1, right)

	// 中间过程数组，保存 [left, mid] 和 [mid+1, right] 区间合并后的有序数组
	cache := make([]int, right-left+1)
	// t:左边数组包含有效逆序对的下标，i:左边数组下标，j:右边数组下标，c:cache数组下标
	t, i, c := left, left, 0
	// 遍历右边数组，统计左右两边数组的逆序对个数
	for j := mid + 1; j <= right; j, c = j+1, c+1 {
		// 不是逆序对，count不需要增加
		for t <= mid && nums[t] <= 2*nums[j] {
			t++
		}
		for ; i <= mid && nums[i] < nums[j]; i, c = i+1, c+1 {
			cache[c] = nums[i]
		}
		cache[c] = nums[j]
		count += mid - t + 1
	}
	// 处理左边数组没有遍历完的情况
	for ; i <= mid; i, c = i+1, c+1 {
		cache[c] = nums[i]
	}
	// 将 排好序的cache 拷贝到 nums[left, right] 区间
	for k := 0; k < len(cache); k++ {
		nums[left+k] = cache[k]
	}
	return count
}

// @lc code=end
