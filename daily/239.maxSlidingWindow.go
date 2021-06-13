package daily

import (
	"container/heap"
	"container/list"
	"fmt"
)

// 239. 滑动窗口最大值
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
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
// @lc: https://leetcode-cn.com/problems/sliding-window-maximum/

// 方法一：大顶堆，优先级队列 PriorityQueue
// 时间复杂度：O(n log(n))，堆的插入删除操作时间复杂度为log(n)，n为数组nums的长度，最坏情况下nums为单调递增数组，那么堆的长度会一直增加不会pop出元素，最坏达到n
// 空间复杂度：O(n)，最坏情况下nums为单调递增数组，堆的长度为n
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 总共可以形成 n-k+1 个窗口
	res := make([]int, 0, n-k+1)

	// 1.大顶堆，先把前k个元素放入堆中
	maxPQ := make(MinHeap, 0, k)
	for i := 0; i < k; i++ {
		maxPQ = append(maxPQ, kvpair{key: i, val: nums[i]})
	}
	heap.Init(&maxPQ)

	// 将第一个窗口内的最大值放入 res
	res = append(res, maxPQ.Peek().val)

	// 2.遍历后面的元素，依次入堆，并判断堆顶元素是否在当前窗口区间内，当前窗口区间下标：[i-k+1 : i] 前后都是闭区间
	for i := k; i < n; i++ {
		heap.Push(&maxPQ, kvpair{key: i, val: nums[i]})
		// 他强任他强，不在当前窗口内的即便是最大值也让其滚蛋
		for maxPQ.Peek().key <= i-k {
			heap.Pop(&maxPQ)
		}
		res = append(res, maxPQ.Peek().val)
	}
	return res
}

// 方法二：单调队列，也叫双端队列 Deque
// 时间复杂度：O(n) n是数组 nums 的长度,每一个下标恰好被放入队列一次，并且最多被弹出队列一次
// 空间复杂度：O(k) 双向队列，不断从2端弹出元素，保证了队列中最多不会有超过 k+1 个元素，因此队列使用的空间为 O(k)
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	deque := list.New()
	// 注意：队列存放的是元素索引
	// push，新加入的元素如果比队列最后一个元素大，则把它压扁（也就是Pop旧元素索引，Push新元素索引），同理可得，保持队首元素最大，保证队列的单调递减
	push := func(i int) {
		// fmt.Printf("before:%+v\n", elements(deque, nums))
		for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] { // 新进入窗口的元素大于队尾元素，干掉他
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
		// fmt.Printf("after:%+v\n", elements(deque, nums))
	}

	// 将前k个 元素的索引 放入队列
	for i := 0; i < k; i++ {
		push(i)
	}

	// 总共有 n-k+1 个窗口形成
	res := make([]int, 0, n-k+1)
	res = append(res, nums[deque.Front().Value.(int)]) // 队首元素
	// 遍历剩下的元素，依次入队，如果队首元素不在当前窗口内，则剔除她，当前窗口区间下标：[i-k+1 : i] 前后都是闭区间
	for i := k; i < n; i++ {
		push(i)                                // 依次入队新元素
		for deque.Front().Value.(int) <= i-k { // 依次剔除不在当前窗口内的队首旧元素
			deque.Remove(deque.Front())
		}
		res = append(res, nums[deque.Front().Value.(int)]) // 队首元素最大，放入结果集
	}

	return res
}

func elements(l *list.List, nums []int) []interface{} {
	eles := []interface{}{}
	for e := l.Front(); e != nil; e = e.Next() {
		key, val := e.Value.(int), nums[e.Value.(int)]
		eles = append(eles, fmt.Sprintf("%d:%d", key, val))
	}
	return eles
}
