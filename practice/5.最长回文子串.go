package practice

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// 中心扩展法，
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	// 从left和right位置开始分别向两边扩散，返回扩散的回文串窗口的长度
	// left = right: 以一个字符为中心扩散
	// right = left+1: 以2个字符为中心扩散
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		// fmt.Printf("L:%d R:%d len:%d\n", left, right, right-left-1)
		// for循环退出时，left和right都多执行了一次，所以求长度需要减去1
		// 对于 "bab" 字符串，当i=1时，
		// 以为a为中心：L:-1 R:3 len:3
		// 以为ab为中心：L:1 R:2 len:0
		return right - left - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(i, i)   // 以s[i]为中心扩散
		len2 := expandAroundCenter(i, i+1) // 以s[i]和s[i+1]为中心扩散
		maxLen := int(math.Max(float64(len1), float64(len2)))
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// @lc code=end
