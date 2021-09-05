package offer

import (
	"container/heap"
	"container/list"
)

// 剑指 Offer 59 - I. 滑动窗口的最大值
// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// 示例:
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:
//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// @lc: https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

// 定义struct用与存储当前元素的索引和值，更直观一些便于操作。
type kvpair struct {
	key, val int
}

// 方法二：大顶堆
// 思路：用大顶堆的属性保证堆顶元素就是窗口内最大，思路最简单
// 1. 先将前 k 个元素建立大顶堆；
// 2. 遍历剩下的元素依次入堆，并判断堆顶元素是否在当前窗口区间内，当前窗口区间下标：[i-k+1 : i] 前后都是闭区间，不在窗口的元素将其移除
// 最终结果：堆顶元素就是最大值，放入结果集即可。
// 时间复杂度：O(n*log(n))， n为数组nums的长度，最坏情况下nums为单调递增数组，那么堆的长度会一直增加不会pop出元素，最坏达到n
// 空间复杂度：O(n)，最坏情况下nums为单调递增数组，堆的长度为n
func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 总共可以形成 n-k+1 个窗口，保存结果集
	res := make([]int, 0, n-k+1)
	// 步骤1
	maxPQ := maxHeap{}
	for i := 0; i < k; i++ {
		maxPQ = append(maxPQ, kvpair{key: i, val: nums[i]})
	}
	heap.Init(&maxPQ)
	// 将第一个窗口最大值存入结果集
	res = append(res, maxPQ.Peek().val)

	// 步骤2
	for i := k; i < n; i++ {
		heap.Push(&maxPQ, kvpair{key: i, val: nums[i]})
		for maxPQ.Peek().key <= i-k {
			heap.Pop(&maxPQ)
		}
		res = append(res, maxPQ.Peek().val)
	}
	return res
}

// 定义大顶堆
type maxHeap []kvpair

// 实现 sort.Interfce{} 的 Len, Less, Swap 接口
func (h maxHeap) Len() int      { return len(h) }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool {
	// 先比较val，val相等时比较key
	if h[i].val == h[j].val {
		return h[i].key > h[j].key
	}
	return h[i].key > h[j].key
}

func (h maxHeap) Peek() kvpair  { return h[0] }
func (h maxHeap) IsEmpty() bool { return h.Len() == 0 }

// 实现 container/Heap 的 Push, Pop方法，receiver是指针
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(kvpair))
}

func (h *maxHeap) Pop() interface{} {
	n := (*h).Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 方法三：双端队列，单调递减队列
// 思路：总共可以形成 n-k+1 个窗口。
// 1. 从后端push时，如果nums[i]大于等于最后一个元素：则把最后一个元素移除（可以理解为：新加入的元素比较强大，把前面比他弱的都干掉了）；否则：直接入队；
// 2. 新元素入队之后，还需要看下队列头部元素是否在当前窗口内，如果不在当前窗口内即使再大也已经成为历史了，让其走人。`当前窗口区间：[i-k+1, i] 左闭右闭区间`；
// 3. i+1 >= k 要保证第一个窗口形成，经过步骤1和2之后队头元素就是最大值。
// 时间复杂度：O(n) 每一个元素恰好被放入队列一次，并且最多被弹出队列一次
// 空间复杂度：O(k) 双向队列，不断从2端弹出元素，保证了队列中最多不会有超过 k+1 个元素，因此队列使用的空间为 O(k)
func maxSlidingWindow3(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 总共可以形成 n-k+1 个窗口，保存结果集
	res := make([]int, 0, n-k+1)
	deque := list.New()
	for i := 0; i < n; i++ {
		// 步骤1
		for deque.Len() > 0 && nums[i] >= deque.Back().Value.(kvpair).val {
			deque.Remove(deque.Back())
		}
		deque.PushBack(kvpair{key: i, val: nums[i]})
		// 步骤2
		if deque.Front().Value.(kvpair).key < i-k+1 {
			deque.Remove(deque.Front())
		}
		// 步骤3
		if i+1 >= k {
			res = append(res, deque.Front().Value.(kvpair).val)
		}
	}
	return res
}
