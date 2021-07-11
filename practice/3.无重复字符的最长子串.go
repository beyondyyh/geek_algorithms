package practice

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
// 方法一：暴力枚举
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)

// 方法二：滑动窗口
// 时间复杂度：O(n)
// 空间复杂度：O(k) k为所有无重复字串长度
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	// 定义窗口区间：[left : right]
	// hashmap k:字符 v:最大的索引
	m := make(map[byte]int)
	// 窗口右边界不断向右扩张
	for left, right := 0, 0; right < n; right++ {
		char := s[right]
		// 说明新进入右窗口的值有重复，此时左窗口向右移动一位，
		if i, has := m[char]; has {
			// 这里max非常关键，比如：abba，先添加a,b进map，此时left=0，再添加b，发现map中包含b，更新 left=m[b]+1=2，此时子段更新为 b，而且map中仍然包含a
			// map.get(a)=0；随后，遍历到a，发现a包含在map中，且map.get(a)=0，实际上，left此时应该不变，left始终为2，子段变成 ba才对
			left = max(left, i+1)
		}
		// 不管是否更新left，都要更新 m[char] 为当前索引
		m[char] = right
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
