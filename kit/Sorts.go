package kit

import (
	"container/heap"
	"sort"
)

// 快排
// 时间复杂度：O(N log(n))，最坏为O(n^2)，不稳定排序
// 基本思想：归并思想，先选择一个轴点pivot，一次遍历将小于pivot的元素放到左边，大于pivot的放到右边，递归对左右两部分进行快排
func quickSort(arr []int, start, end int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 分区操作，返回轴点索引下标
	var partition func(start, end int) int
	partition = func(start, end int) int {
		pivot := start    // 选取第一个元素作为轴点
		start = start + 1 // 遍历的游标
		for i := start; i <= end; i++ {
			if arr[i] < arr[pivot] {
				arr[i], arr[start] = arr[start], arr[i] // 比轴点小的交换到前面
				start++
			}
		}
		arr[pivot], arr[start-1] = arr[start-1], arr[pivot]
		// fmt.Printf("start: %d->%d end: %d->%d start:%d pivot: %d->%d  %+v\n", start, arr[start], end, arr[end], start, pivot, arr[pivot], arr)
		return start - 1
	}

	// recursion
	if start < end {
		pivot := partition(start, end)
		// fmt.Printf("pivot:%d->%d arr:%v\n", pivot, arr[pivot], arr)
		quickSort(arr, start, pivot-1)
		quickSort(arr, pivot+1, end)
	}
	return arr
}

// 归并排序
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
