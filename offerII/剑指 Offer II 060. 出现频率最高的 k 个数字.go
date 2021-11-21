package offerII

import "container/heap"

// 剑指 Offer II 060. 出现频率最高的 k 个数字
// 给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。

// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:
// 输入: nums = [1], k = 1
// 输出: [1]
// 提示：
// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
// 进阶：所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
// @lc: https://leetcode-cn.com/problems/g5c51o/

type pair struct {
	num int // 元素
	cnt int // 出现的次数
}

// 方法一：小顶堆
// 先遍历nums用haspmap统计元素出现次数，然后维护一个长度为k的小顶堆
// 时间复杂度：O(NlogK) n为数组长度。遍历获得map为O(n)，每次pop需要heapifyDown为O(logK)，所以总共为O(NlogK)
// 空间复杂度：O(N)。哈希表的大小为 O(N)，堆的大小为 O(k)，共计为 O(N)
func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	h := iheap{}
	heap.Init(&h)
	for num, cnt := range freq {
		heap.Push(&h, pair{num: num, cnt: cnt})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	// 倒叙输出，最大达放在数组头部
	res := make([]int, k)
	for ; k > 0; k-- {
		res[k-1] = heap.Pop(&h).(pair).num
	}
	return res
}

// 小顶堆，
type iheap []pair

// type Interface interface {
// 	sort.Interface
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   // remove and return element Len() - 1.
// }
// 实现 sort.Interface 接口的 Len, Less, Swap方法
func (h iheap) Len() int           { return len(h) }
func (h iheap) Less(i, j int) bool { return h[i].cnt < h[j].cnt }
func (h iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现 Push,Pop 方法，receiver是指针
func (h *iheap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = (*h)[:n-1]
	return x
}

// 方法二：利用快排的分治思想，也是最有解
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func topKFrequent2(nums []int, k int) []int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	pairs := []pair{}
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
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
func quickSortPair(pairs []pair, lo, hi int, k int) []pair {
	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间[lo, hi] 执行 partition 操作，返回（选取的第一个nums[lo]作为pivot）排序以后应该在的位置
	partition := func(pairs []pair, lo, hi int) int {
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
