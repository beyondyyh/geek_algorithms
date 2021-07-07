package practice

import "container/heap"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

// 方法一：调用库函数降序排序后，nums[k-1]就是第K大
// 时间复杂度：O(n log(n)) 快排时间复杂度
// 空间复杂度：O(n log(n)) 快排空间复杂度

// 方法二：维护一个长度为k的小顶堆，利用堆的特性每次入堆进行heapifyUp调整，堆顶元素就是第K大
// 时间复杂度：O(n log(k))
// 空间复杂度：O(k) 堆的大小
func findKthLargest1(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	h := iheap{}
	heap.Init(&h)
	// 先把前k个元素构建小顶堆
	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		// 看一下堆顶元素，如果小于等于当前nums[i]就没必要替换了，减少调整次数
		top := h.Peek()
		if nums[i] > top {
			heap.Pop(&h) // 弹出堆顶
			heap.Push(&h, nums[i])
		}
	}
	// 此时堆顶元素就是第K大
	return h.Peek()
}

// 以上 也可以每次先Push，然后在判断长度是否大于K，大于K时再Pop，这样是利用堆的自调整特性保证长度为k
// 不过调整次数比增加

type iheap []int

// 实现 Sort 接口的Len, Less, Swap方法
func (h iheap) Len() int           { return len(h) }
func (h iheap) Less(i, j int) bool { return h[i] < h[j] }
func (h iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 看一下堆顶元素，不调整
func (h iheap) Peek() int {
	return h[0]
}

// 实现 container/Heap 接口的Pop,Push方法，receiver需要是指针
func (h *iheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 弹出数组最后一个元素，同时数组长度--，堆内部通过替换堆顶元素来进行 heapifyDown 调整
func (h *iheap) Pop() interface{} {
	n := (*h).Len()
	x := (*h)[n-1]  // 数组最后一个元素
	*h = (*h)[:n-1] // size--
	return x
}

// @lc code=end
