package practice

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
// 先按区间的 start其实元素 排升序
// 遍历区间，设当前区间[iStart, iEnd]，结果集中最后一个区间的end为lastEnd，则：
// iStart > lastEnd 说明二者没有交集，直接将当前区间加入到结果集即可
// iStart <= lastEnd 说明二者有交集，需要更新结果集里的最后一个区间的结束位置为二者较大，lastEnd=max(lastEnd, iEnd)
// 时间复杂度：O(NlogN)，N是区间的长度；
// 空间复杂度：O(n)，最坏情况是所有的区间都没有交点的
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 先按区间的 start其实元素 排升序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res := make([][]int, 0, n)
	res = append(res, intervals[0])
	for _, curInterval := range intervals {
		last := res[len(res)-1]
		if curInterval[0] > last[1] {
			res = append(res, curInterval)
		} else {
			last[1] = max(curInterval[1], last[1])
		}
	}
	return res
}

// @lc code=end
