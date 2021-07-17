package practice

// 快速排序。不稳定排序 in-place
// 思路：分治思想的典型应用，首先选择一个标杆（轴点），通过一趟遍历将数据分割成左右两部分，将小于它的放到它的左侧，大于它的放到它的右侧。遍历结束之后标杆所在的位置就是排序之后它应该在的位置。
// 递归地将左右两部分分别调用快排。
// 时间复杂度：O(N log(n))，最坏为O(n^2)，不稳定排序
func quickSort(nums []int, left, right int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间 [left, right] 执行 partition 操作，返回 nums[left] 排序以后应该在的位置
	// 在遍历过程中保持循环不变量的语义
	// 1、[left + 1, j] < nums[left]
	// 2、(j, i] >= nums[left]
	// 3、交换 nums[left]和nums[j]
	partition := func(nums []int, left, right int) int {
		// 选取第一个元素作为轴点
		pivot := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < pivot {
				// 小于 pivot 的元素都被交换到前面
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		// 在之前遍历的过程中，满足 [left + 1, j] < pivot，并且 (j, i] >= pivot
		nums[j], nums[left] = nums[left], nums[j]
		// 交换以后 [left, j - 1] < pivot, nums[j] = pivot, [j + 1, right] >= pivot
		return j
	}

	// recursion
	if left < right {
		pivot := partition(nums, left, right)
		// fmt.Printf("pivot:%d->%d nums:%v\n", pivot, nums[pivot], nums)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
	return nums
}

// 归并排序。稳定排序算法
// 思路：分治思想的典型应用，利用递归不断将数组一分为二，直到只剩下一个元素，一个元素的数组是天然有序的，然后合并2个有序数组的过程称为：2-路归并。
// 时间复杂度： O(n log(n))，稳定排序！
// 空间复杂度： O(n)，2-路归并的额外空间
func mergeSort(nums []int) []int {
	// terminator，最后切分只剩下一个元素，可以认为是天然有序的，直接返回
	if len(nums) <= 1 {
		return nums
	}

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

	// 分，一分为二
	mid := len(nums) / 2
	left, right := mergeSort(nums[:mid]), mergeSort(nums[mid:])
	// 治，合二为一
	return merge(left, right)
}
