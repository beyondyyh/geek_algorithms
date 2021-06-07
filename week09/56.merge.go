package week09

import (
	"sort"
)

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 示例 1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// @lc: https://leetcode-cn.com/problems/merge-intervals/

// 时间复杂度：O(NlogN)，N是区间的长度；
// 空间复杂度：O(n)，最坏情况是所有的区间都没有交点的
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	// 先按起点进行排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	// 结果集
	res := make([][]int, 0, n)
	res = append(res, intervals[0])

	// 遍历区间，每次新遍历到的区间的 starti 与当前结果集中的最后一个区间的 endLast 进行比较，
	// 1. 如果 starti > endLast，说明它们没有交集，直接把该区间添加到结果集；
	// 2. 如果 starti <= endLast，说明他们有交集的，此时需要合并操作，即：更新endLast，取2个区间的end的最大值
	for _, curInterval := range intervals {
		last := res[len(res)-1] // 取结果集中的最后一个元素
		if curInterval[0] > last[1] {
			res = append(res, curInterval)
		} else {
			last[1] = max(last[1], curInterval[1])
			// res[len(res)-1] = last
		}
	}

	return res
}
