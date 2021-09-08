package kit

import (
	"container/heap"
	"sort"
)

// 快速排序，指定排序区间 in-place
// 时间复杂度：O(N log(n))，最坏为O(n^2)，不稳定排序
// 基本思想：归并思想，先选择一个轴点pivot，一次遍历将小于pivot的元素放到左边，大于pivot的放到右边，递归对左右两部分进行快排
func quickSort(arr []int, left, right int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 分区操作，返回轴点索引下标
	// 在数组arr的子区间 [left, right] 执行 partition 操作，返回 arr[left] 排序以后应该在的位置
	// 在遍历过程中保持循环不变量的语义
	// 1、[left + 1, j] < arr[left]
	// 2、(j, i] >= arr[left]
	// 3、交换 arr[left]和arr[j]
	partition := func(arr []int, left, right int) int {
		// 选取第一个元素作为轴点
		pivot := arr[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if arr[i] < pivot {
				// 小于 pivot 的元素都被交换到前面
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		// 在之前遍历的过程中，满足 [left + 1, j] < pivot，并且 (j, i] >= pivot
		arr[j], arr[left] = arr[left], arr[j]
		// 交换以后 [left, j - 1] < pivot, nums[j] = pivot, [j + 1, right] >= pivot
		return j
	}

	// recursion
	if left < right {
		pivot := partition(arr, left, right)
		// fmt.Printf("pivot:%d->%d arr:%v\n", pivot, arr[pivot], arr)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
	return arr
}

// 归并排序，整体排序
// 时间复杂度：O(N log(n))，稳定排序
// 基本思想：利用归并的思想，采用分治的策略
// 分治：分是将大问题分成一些小问题然后递归求解，而治则是将分的阶段得到的答案“合并”在一起
func mergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}

	// 合并2个有序数组
	var _merge func(left, right []int) []int
	_merge = func(left, right []int) []int {
		m, n := len(left), len(right)
		res := make([]int, 0, m+n)
		var i, j int
		for i < m && j < n {
			if left[i] <= right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		if i < m {
			res = append(res, left[i:]...)
		}
		if j < n {
			res = append(res, right[j:]...)
		}
		return res
	}

	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return _merge(left, right)
}

// 堆排序
// 时间复杂度：O(N log(n))，稳定排序
func heapSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	res := make([]int, 0, n)
	// 构建小顶堆
	h := &iheap{arr}
	heap.Init(h)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 利用golang标准库实现堆，只需要实现 heap接口的 `Push, Pop`方法即可
type iheap struct{ sort.IntSlice }

// // 默认升序，倒序的话需要重写 sort.Less 比较器
// func (h *iheap) Less(i, j int) bool {
// 	return h.IntSlice[i] > h.IntSlice[j]
// }

func (h *iheap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// Pop 返回数组最后一个元素，heap接口通过 heapifyDown 策略去调整堆，以维护二叉堆的特性
// heapifyDown：1.弹出堆顶元素 2.把数组最后一个元素放到堆顶（数组第0位置），然后自上而下heapifyDown调整以到达平衡
func (h *iheap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

// 归并排序，指定排序区间 in-place
// 时间复杂度：O(N log(n))，稳定排序
func mergeSort2(nums []int, left, right int) {
	// 合并 [left, mid] 和 [mid+1, right] 2个有序区间
	merge := func(nums []int, left, mid, right int) {
		// 中间数组，存放 [left, right] 区间合并后的有序数组
		temp := make([]int, 0, right-left+1)
		i, j := left, mid+1
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				temp = append(temp, nums[i])
				i++
			} else {
				temp = append(temp, nums[j])
				j++
			}
		}
		// 处理 [left, mid] 未遍历完的情况
		for ; i <= mid; i++ {
			temp = append(temp, nums[i])
		}
		// 处理 [mid+1, right] 未遍历完的情况
		for ; j <= right; j++ {
			temp = append(temp, nums[j])
		}
		// 有序数组 temp 拷贝回 nums[left, right] 区间
		for k := 0; k < len(temp); k++ {
			nums[left+k] = temp[k]
		}
	}

	// main process
	if left >= right { // 递归终止条件
		return
	}
	mid := (left + right) >> 1
	mergeSort2(nums, left, mid)
	mergeSort2(nums, mid+1, right)
	merge(nums, left, mid, right)
}
