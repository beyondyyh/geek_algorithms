package week02

import (
	"container/heap"
	"fmt"
)

// 347. 前 K 个高频元素
// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// @leetcode: https://leetcode-cn.com/problems/top-k-frequent-elements/

// 方法一：小顶堆
// 先遍历nums用haspmap统计元素出现次数，然后维护一个长度为k的小顶堆
// 遍历「出现次数数组」：
// 如果堆的元素个数小于 k，就可以直接插入堆中；
// 如果堆的元素个数等于 k，则检查堆顶与当前出现次数的大小。如果堆顶更大，说明至少有 k个数字的出现次数比当前值大，故舍弃当前值；否则，就弹出堆顶，并将当前值插入堆中。
// 时间复杂度：O(NlogK) n为数组长度。遍历获得map为O(n)，每次pop需要heapifyDown为O(logK)，所以总共为O(NlogK)
// 空间复杂度：O(N)。哈希表的大小为 O(N)，堆的大小为 O(k)，共计为 O(N)
func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}
	fmt.Printf("frequency: %+v\n", frequency)

	// 维护长度为k的小顶堆
	h := &Iheap{}
	heap.Init(h)
	for num, count := range frequency {
		heap.Push(h, [2]int{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	fmt.Printf("heap: %+v\n", *h)

	// 最大的放最前面，所以从后往前追加
	ans := make([]int, k)
	for ; k > 0; k-- {
		ans[k-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

// ---------二维数组小顶堆的实现---------
// [2]int{元素, 出现的次数}
type Iheap [][2]int

// 实现container/heap interface{}内嵌的sort接口
func (h Iheap) Len() int           { return len(h) }
func (h Iheap) Less(i, j int) bool { return h[i][1] < h[j][1] } // 比较次数，小顶堆
func (h Iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 实现container/heap interface{}的Push，Pop接口，receiver为指针
func (h *Iheap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop 堆的删除操作，先用数组尾部元素替换数组头部元素，然后从上而下执行`heapifyDown`操作
// Pop返回数组尾部元素
func (h *Iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type freqPair struct {
	num, cnt int
}

// 方法二：利用快排的分治思想，也是最有解
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func topKFrequent2(nums []int, k int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	pairs := []freqPair{}
	for k, v := range freq {
		pairs = append(pairs, freqPair{k, v})
	}

	// 使用快排变形 O(N) 获取Pair数组中频次最大的 k 个元素（第 4 个参数是下标，因此是 k - 1）
	topkPairs := quickSortPair(pairs, 0, len(pairs)-1, k-1)
	res := make([]int, 0, k)
	for _, p := range topkPairs {
		res = append(res, p.num)
	}
	return res
}

// 快速排序，逆序排序
// 最大的 k 个元素 的下标
func quickSortPair(pairs []freqPair, lo, hi int, k int) []freqPair {
	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间[lo, hi] 执行 partition 操作，返回（选取的第一个nums[lo]作为pivot）排序以后应该在的位置
	partition := func(pairs []freqPair, lo, hi int) int {
		// 选第一个数作为轴点
		pivot := pairs[lo].cnt
		j := lo
		// 遍历剩下的元素，将大于 pivot 的元素都被交换到前面
		for i := lo + 1; i <= hi; i++ {
			if pairs[i].cnt > pivot {
				j++
				pairs[j], pairs[i] = pairs[i], pairs[j]
			}
		}
		// 在之前遍历的过程中，满足 [lo+1, j] > pivot，并且 (j, i] <= pivot
		pairs[j], pairs[lo] = pairs[lo], pairs[j]
		// 交换以后 [lo, j-1] > pivot, pairs[j] = pivot, [j+1, hi] <= pivot
		return j
	}

	// 快排切分后，j左边的数字都大于等于 pairs[j].cnt，j右边的数字出现的次数都小于等于 pairs[j].cnt
	for {
		index := partition(pairs, lo, hi)
		if index == k { // 如果 index 正好等于目标k，说明当前pairs数组中的[0,index]就是出现次数最大的 K 个元素。
			return pairs[0 : k+1]
		} else if index < k { //
			lo = index + 1
		} else {
			hi = index - 1
		}
	}
}
