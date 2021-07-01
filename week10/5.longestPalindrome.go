package week10

import "math"

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
// @lc: https://leetcode-cn.com/problems/longest-palindromic-substring/

// 方法一：中心扩展法，遍历字符串 s，以当前1个字符或2个字符为中心向两边扩散，找到有效的回文串
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	// 以left和right为中心分别向两边扩散，找到有效的回文字串，返回字串的长度
	// left == right 表示以一个字符为中心扩散，得到的回文串是奇数长度
	// left != right 表示以两个字符为中心扩散，得到的回文串是偶数长度
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		// 遍历结束后，left和right 分别多走了一步，因此是 right-left-1
		return right - left - 1
	}

	// 记录得到有效回文串时的下标位置
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i)   // 以当前1字符s[i]为中心扩散
		len2 := expandAroundCenter(i, i+1) // 以当前2个字符s[i]和s[i+1]为中心扩散
		maxlen := int(math.Max(float64(len1), float64(len2)))
		if maxlen > end-start { // 得到的新串更长，更新res
			start, end = i-(maxlen-1)/2, i+maxlen/2
		}
	}
	return s[start : end+1]
}
