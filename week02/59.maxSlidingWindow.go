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
		window := nums[i : i+k]
		ans = append(ans, max(window))
	}
	return ans
}

func max(nums []int) int {
	sort.Ints(nums) // Asc
	return nums[len(nums)-1]
}

// 方法二：大顶堆
// 时间复杂度：O(NlogK)
// 空间复杂度：
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 初始化前K的元素到堆中
	maxPQ := make(PriorityQueue, k)
	for i := 0; i < k; i++ {
		maxPQ[i] = &Item{Index: i, Value: nums[i]}
	}
	heap.Init(&maxPQ)

	// 总共有n-k+1个窗口，声名一个长度长度&容量n-k+1的slice
	ans := make([]int, n-k+1)
	// 堆顶元素即是第一个窗口最大值，先放进ans
	ans[0] = maxPQ.Peek().Value

	// 遍历将剩下的元素依次入堆
	for i := k; i < n; i++ {
		// 将新元素入堆
		item := &Item{Index: i, Value: nums[i]}
		heap.Push(&maxPQ, item)
		// 循环判断当前堆顶是否在窗口中，一般思路是遍历窗口元素与堆顶进行对比，时间复杂度为O(k)
		// 反向思维：堆顶元素已经是最大值，可以依次pop比较堆顶元素的下标是否小于窗口的左边界i-k+1，直到堆为空或者堆顶元素下标等于左边界，出栈时间复杂度O(1)
		for maxPQ.Len() > 0 && maxPQ.Peek().Index <= i-k {
			heap.Pop(&maxPQ)
		}
		// 在窗口中直接赋值即可
		if maxPQ.Len() > 0 {
			ans[i-k+1] = maxPQ.Peek().Value
		}
	}
	return ans
}
