package offer

import "container/heap"

// 剑指 Offer 41. 数据流中的中位数
// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

// 例如，
// [2,3,4] 的中位数是 3
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
// 设计一个支持以下两种操作的数据结构：
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。

// 示例 1：
// 输入：
// ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
// [[],[1],[2],[],[3],[]]
// 输出：[null,null,null,1.50000,null,2.00000]

// 示例 2：
// 输入：
// ["MedianFinder","addNum","findMedian","addNum","findMedian"]
// [[],[2],[],[3],[]]
// 输出：[null,null,2.00000,null,2.50000]
// @lc: https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/

// 方法：左手一个大顶堆 + 右手一个小顶堆
// 解题思路
// 1. 我们没有必要把所有数据进行排序。只需要保证数据左半边的数都小于右半边的数，那么根据 左半边的数的最大值 及 右半边的数的最小值 即可得到中位数。
// 2. 若输入的数据个数是奇数，比如 1、2、3、4、5。我们可以把左边的 1、2 存入一个大顶堆中，把右边的 3、4、5 存入一个小顶堆中。那么中位数就是小顶堆的 top()。
// 3. 若输入的数据个数是偶数，比如 1、2、3、4。我们可以把左边的 1、2 存入一个大顶堆中，把右边的 3、4 存入一个小顶堆中。那么中位数就是两个堆的 top() 的和再乘 0.5。
// 4. 整个过程我们需要维护两个地方：两个堆的 size() 最大只能相差 1；大顶堆的 top() 必须小于等于小顶堆的 top()。

type MedianFinder struct {
	minH iMinHeap // 小顶堆，保存较大的一半
	maxH iMaxHeap // 大顶堆，保存较小的一半
}

// func Constructor() MedianFinder {
func NewMedianFinder() *MedianFinder {
	minH, maxH := iMinHeap{}, iMaxHeap{}
	heap.Init(&minH)
	heap.Init(&maxH)
	finder := &MedianFinder{
		minH: minH,
		maxH: maxH,
	}
	return finder
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.minH.Len() != mf.maxH.Len() {
		heap.Push(&mf.minH, num)
		heap.Push(&mf.maxH, heap.Pop(&mf.minH))
	} else {
		heap.Push(&mf.maxH, num)
		heap.Push(&mf.minH, heap.Pop(&mf.maxH))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.minH.Len() == mf.maxH.Len() {
		return float64(mf.minH.Peek()+mf.maxH.Peek()) * 0.5
	}
	return float64(mf.minH.Peek())
}
