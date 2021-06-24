package practice

import "container/list"

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

// 方法一：暴力枚举，每个窗口区间[i-k+1, i]内的最大值
// 时间复杂度：接近 O(n^2)

// 方法二：大顶堆，
// 时间复杂度：O(n*log(k))
// 空间复杂度：O(n) 最差情况数组为单调递增，堆的长度等于数组长度

// 方法一：双端队列，单调递减队列
// 时间复杂度：O(n) 每一个元素恰好被放入队列一次，并且最多被弹出队列一次
// 空间复杂度：O(k) 双向队列，不断从2端弹出元素，保证了队列中最多不会有超过 k+1 个元素，因此队列使用的空间为 O(k)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	type kvpair struct {
		key, val int
	}
	// 一共可以形成 n-k+1 个窗口
	res := make([]int, 0, n-k+1)
	// 双端队列
	deque := list.New()
	for i := 0; i < n; i++ {
		// 1. 从后面push操作
		// 从后端push时，如果nums[i]大于等于最后一个元素，则把最后一个元素移除（可以理解为：大的把小的给干掉了），否则直接入队
		for deque.Len() > 0 && nums[i] >= deque.Back().Value.(kvpair).val {
			deque.Remove(deque.Back())
		}
		// 把元素index和value同时放入队列中，便于比较
		deque.PushBack(kvpair{key: i, val: nums[i]})

		// 2. push进去之后，还需要看下队列头部元素是否在当前窗口内，如果不在当前窗口内即使再大也让其滚犊子
		// 当前窗口区间：[i-k+1, i] 左闭右闭区间
		for deque.Front().Value.(kvpair).key < i-k+1 { // 此处用 if 也是一样的，
			deque.Remove(deque.Front())
		}

		// 3. 此处要保证第一个窗口形成，队头元素就是最大值
		if i+1 >= k {
			res = append(res, deque.Front().Value.(kvpair).val)
		}
	}
	return res
}

// @lc code=end
