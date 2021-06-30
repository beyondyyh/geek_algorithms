package practice

// 快速排序。不稳定排序
// 思路：分治思想的典型应用，首先选择一个标杆（轴点），通过一趟遍历将数据分割成左右两部分，将小于它的放到它的左侧，大于它的放到它的右侧。遍历结束之后标杆所在的位置就是排序之后它应该在的位置。
// 递归地将左右两部分分别调用快排。
func quickSort(nums []int, start, end int) []int {
	// terminator，长度小于等于1说明天然有序
	if len(nums) <= 1 {
		return nums
	}

	// 分，选择一个轴点 pivot，通过一次遍历将小于它的放到它的左侧，大于它的放到他的右侧，最后返回标杆所在的下标。
	// - start 起始位置的下标
	// - end 结束位置的下标
	// returns 轴点所在的下标
	pratition := func(start, end int) int {
		pivot := start // 选取第一个元素作为轴点，当然也可以选择最后一个
		start = start + 1
		for i := start; i <= end; i++ {
			if nums[i] <= nums[pivot] {
				nums[i], nums[start] = nums[start], nums[i] // 比轴点小的交换到前面
				start++
			}
		}
		// 遍历结束之后，把轴点放到正确的位置上，for退出时start多走了一步
		nums[pivot], nums[start-1] = nums[start-1], nums[pivot]
		return start - 1
	}

	// 分治，找到轴点对应的下标，然后把轴点左右两边的数组分别递归调用上面的思想
	// start, end := 0, len(nums)-1
	if start < end {
		pivot := pratition(start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
	return nums
}

// 归并排序。稳定排序算法
// 思路：分治思想的典型应用，利用递归不断将数组一分为二，直到只剩下一个元素，一个元素的数组是天然有序的，然后合并2个有序数组的过程称为：2-路归并。
// 时间复杂度： O(n log(n))，稳定排序！
// 空间复杂度： O(n)，2-路归并的额外空间
func mergeSort(nums []int) []int {
	// 2-路归并
	merge := func(left, right []int) []int {
		m, n := len(left), len(right)
		nums := make([]int, 0, m+n)
		var i, j int
		// 依次将较小的数放入nums，同时指针后移一位
		for i < m && j < n {
			if left[i] <= right[j] {
				nums = append(nums, left[i])
				i++
			} else {
				nums = append(nums, right[j])
				j++
			}
		}
		// 处理其中一个数组没遍历完的情况
		if i < m {
			nums = append(nums, left[i:]...)
		}
		if j < n {
			nums = append(nums, right[j:]...)
		}
		return nums
	}

	// terminator，最后切分只剩下一个元素，可以认为是天然有序的，直接返回
	if len(nums) <= 1 {
		return nums
	}
	// 分，一分为二
	mid := len(nums) / 2
	left, right := mergeSort(nums[:mid]), mergeSort(nums[mid:])
	// 治，合二为一
	return merge(left, right)
}
