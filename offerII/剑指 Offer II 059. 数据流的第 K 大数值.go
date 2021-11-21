package offerII

import (
	"container/heap"
	"sort"
)

// 剑指 Offer II 059. 数据流的第 K 大数值
// 设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
// 请实现 KthLargest 类：
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
// 示例：
// 输入：
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// 输出：
// [null, 4, 5, 5, 8, 8]
// 解释：
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3);   // return 4
// kthLargest.add(5);   // return 5
// kthLargest.add(10);  // return 5
// kthLargest.add(9);   // return 8
// kthLargest.add(4);   // return 8
// @lc: https://leetcode-cn.com/problems/jBjn9C/

// 小顶堆，维护一个长度为k的小根堆，堆顶元素就是第k大
// 堆Pop和Push的平均时间复杂度都是：O(logK)

type KthLargest struct {
	sort.IntSlice
	k int
}

// sort.IntSlice实现了 实现 container/Heap 中sort接口的 Len, Less, Swap方法

// 实现 container/Heap 的 Push, Pop方法
// 不仅要改变元素值，还要改变数组长度，所以receiver用指针
func (k *KthLargest) Push(x interface{}) {
	k.IntSlice = append(k.IntSlice, x.(int))
}

// 弹出数组最后一个元素，用于替换堆顶，然后执行自下而上 heapifyUp
func (k *KthLargest) Pop() interface{} {
	old := k.IntSlice
	x := old[old.Len()-1]                 // 最后一个元素
	k.IntSlice = k.IntSlice[:old.Len()-1] // 长度减一
	return x
}

// func Constructor(k int, nums []int) KthLargest {
func NewKthLargest(k int, nums []int) KthLargest {
	h := KthLargest{k: k}
	heap.Init(&h)
	for _, num := range nums {
		h.Add(num)
	}
	return h
}

func (h *KthLargest) Add(val int) int {
	heap.Push(h, val)
	if h.Len() > h.k {
		heap.Pop(h)
	}
	return h.IntSlice[0]
}
