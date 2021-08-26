package offer

import (
	"container/heap"
	"sort"
)

// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

// 示例 1：
// 输入：arr = [3,2,1], k = 2
// 输出：[1,2] 或者 [2,1]

// 示例 2：
// 输入：arr = [0,1,2,1], k = 1
// 输出：[0]
// @lc: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

// 方法一：利用快排思想，通过partition找到索引为 K-1 的数，那么它左边的数就是比它小的另外 K-1 个数！不需要全部排序
// 时间复杂度：O(n)
func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 || len(arr) == 0 {
		return nil
	}
	// 最后一个参数表示索引，因此我们要找的是下标为 k-1 的那个数
	return quickSearch(arr, 0, len(arr)-1, k-1)
}

// 快排模板
// lo, hi 分别表示区间低位和高位，k 表示是要找的下标为k的数
// 返回 [0,k] 的最小k个数
func quickSearch(nums []int, lo, hi, k int) []int {
	// 分治，选取第一个元素作为轴点povit，遍历剩下的元素把小于povit的放到左边，大于或等于povit放到右边
	// 返回 povit 在排序后应该在的索引位置
	partition := func(nums []int, lo, hi int) int {
		povit := nums[lo] // 选第一个数作为轴点
		j := lo           // j表示轴点的索引下标
		for i := lo + 1; i <= hi; i++ {
			if nums[i] < povit {
				j++
				// 比pivot小的交换到左边，大于或等于povit不处理，自然就在右边
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		// 一次遍历结束后满足：[lo+1, j] < pivot，(j, i] >= pivot
		nums[lo], nums[j] = nums[j], nums[lo]
		// 交换以后满足：[lo, j-1] < pivot, nums[j] = pivot, [j+1, hi] >= pivot
		return j
	}

	// 每次快排切分之后，找到排序后下标为 povit 的元素
	// 1. 如果povit等于k：那么povit及其之前的数就是最小的k个数，直接返回
	// 2. 如果povit大于k：则继续去左半部分找
	// 3. 如果povit小于k：则继续去右半部分找
	povit := partition(nums, lo, hi)
	if povit == k {
		return nums[0 : povit+1]
	} else if povit > k {
		return quickSearch(nums, lo, povit-1, k)
	} else {
		return quickSearch(nums, povit+1, hi, k)
	}
}

// 方法二：大顶堆
func getLeastNumbers2(nums []int, k int) []int {
	h := &IMaxHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// 大顶堆
type IMaxHeap struct{ sort.IntSlice }

// 默认是小顶堆，重写 Less 方法
func (h IMaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// 注意：Push、Pop的receiver需要是指针
func (h *IMaxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// 返回最后一个元素，用于替换堆顶元素，然后执行heapifyDown自顶向下调整
func (h *IMaxHeap) Pop() interface{} {
	old := h.IntSlice
	n := h.IntSlice.Len()
	x := old[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}
