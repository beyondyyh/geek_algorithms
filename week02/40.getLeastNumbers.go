package week02

import (
	"container/heap"
	"sort"
)

// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
// 示例 1：
// 输入：arr = [3,2,1], k = 2
// 输出：[1,2] 或者 [2,1]
// @leetcode: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

// 方法一：从小到大排序之后取前k个
// 时间复杂度：O(NlogN) 快排思想，n为数组长度
// 空间复杂度：O(logN)
func getLeastNumbers1(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	sort.Ints(nums)
	return nums[0:k]
}

// 方法二：大顶堆，
// 维护一个size为k的大顶堆，首先将前 k 个数插入大根堆中，随后从第 k+1 个数开始遍历，如果当前遍历到的数比大根堆的堆顶的数要小，就把堆顶的数弹出，再插入当前遍历到的数。
// 这样相当于 n-k 个大的数字已经被从堆中删除了，上下的k个数就是 前k小的，最后依次出堆即可
// 时间复杂度：O(NlogK) 快排思想
// 空间复杂度：O(NlogK)
func getLeastNumbers2(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	// 创建大顶堆，先将前k个元素入堆
	h := &MaxHeap{}
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	heap.Init(h)

	// 遍历剩下的元素，如果比堆顶小则将堆顶删除，入堆该元素，内部调整堆结构
	for i := k; i < len(nums); i++ {
		if h.Peek().(int) > nums[i] {
			heap.Pop(h)           // 删除堆顶元素
			heap.Push(h, nums[i]) // 当前元素入堆，内部调整结构
		}
	}

	// 堆内元素便是前k小，无序，依次Pop便是有序
	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(int))
	}
	return ans
}

// 方法三：大顶堆，与方法二一样，只不过代码更精简
// 思路：维护一个size为k的大顶堆，遍历数组先将元素入堆（内部调整结构），只需要判断当前对的size是否大于k，如果比k大则直接Pop出堆顶最大的，依次类推
func getLeastNumbers3(nums []int, k int) []int {
	if k == 0 {
		return nil
	}

	h := &MaxHeap{}
	heap.Init(h)
	for num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 堆内元素便是前k小，无序，依次Pop便是有序
	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(int))
	}
	return ans
}
