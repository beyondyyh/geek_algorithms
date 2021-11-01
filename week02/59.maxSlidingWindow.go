package week02

import (
	"container/heap"
	"sort"
)

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
// 示例 1：
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
//
// @leetcode: https://leetcode-cn.com/problems/sliding-window-maximum

// 方法一：暴力求解，双层遍历
// 时间复杂度：O(x * k(log(k))), x为窗口个数，x=n-k+1 n为数组长度
// 空间复杂度：O(x)
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 窗口个数
	ans := make([]int, 0, n-k+1)
	for i := 0; i < n-k+1; i++ {
		window := make([]int, k)
		copy(window, nums[i:i+k]) // 需要copy一份，快排是in-place的会改变nums的结构
		ans = append(ans, max(window))
	}
	return ans
}

func max(nums []int) int {
	sort.Ints(nums) // Asc
	return nums[len(nums)-1]
}

// 方法二：大顶堆
// 时间复杂度：O(n log(n))， n为数组nums的长度，最坏情况下nums为单调递增数组，那么堆的长度会一直增加不会pop出元素，最坏达到n
// 空间复杂度：O(n)，最坏情况下nums为单调递增数组，堆的长度为n
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 总共会形成n-k+1个窗口，res存放结果集
	res := make([]int, 0, n-k+1)

	// 1. 先将第一个窗口内的元素初始化入堆
	q := make(PriorityQueue, 0, k)
	for i := 0; i < k; i++ {
		item := pair{index: i, value: nums[i]}
		q = append(q, item)
	}
	heap.Init(&q)
	// 将第一个窗口最大值存入结果集
	res = append(res, q.Peek().value)
	// 2. 遍历元素，依次入堆，并判断堆顶元素是否在当前窗口区间内，当前窗口区间下标：[i-k+1 : i] 前后都是闭区间
	for i := k; i < n; i++ {
		item := pair{index: i, value: nums[i]}
		heap.Push(&q, item)
		// 他强任他强，不在当前窗口内的即便是最大值也让其滚蛋
		for q.Peek().index <= i-k {
			heap.Pop(&q)
		}
		res = append(res, q.Peek().value)
	}
	return res
}

// ---------------------------------------------------------------
// 借助标准库 container/Heap 接口实现大顶堆（优先队列）
type pair struct {
	index int // 索引
	value int // 值
}

type PriorityQueue []pair

// 实现 sort接口的 Len,Less,Swap 方法
func (pq PriorityQueue) Len() int      { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// value相同时，比较索引
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].value == pq[j].value {
		return pq[i].index > pq[j].index
	}
	return pq[i].value > pq[j].value
}

// 实现 container/Heap 的 Push,Pop 方法，注意：receiver是指针
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := pq.Len()
	x := old[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func (pq PriorityQueue) Peek() pair {
	return pq[0]
}
