package offerII

// 剑指 Offer II 016. 不含重复字符的最长子字符串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 滑动窗口，整体思路：
// 1. 初始化窗口头尾指针：left, right := 0, 0
// 2. right右移，判断right指向的元素是否在 [left:right] 窗口内，
//	  不在：将该元素加入窗口，right后移，同时更新res = max(res, right-left+1)
// 	    在：将left指针后移，直到窗口内不包含该元素位置，注意：是一直后移直到不包含，而不是只移动一位。因为字串是连续的，如：abcb 这种情况
// 时间复杂度：O(n^2) 遍历一次n，判断right指向的元素是否在窗口内又需要一次遍历，最差情况整个串没有重复，时间复杂度为 n^2
// 空间复杂度：O(1)

// 方法二：滑动窗口，利用hashmap保存窗口内的元素和下标
// 时间复杂度：O(n)
// 空间复杂度：O(n) 最差情况所有字符都不重复
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	res := 0
	// 定义窗口区间：[left:right]
	// hashmap k:字符 v:最大的索引，注意此处是最大
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
