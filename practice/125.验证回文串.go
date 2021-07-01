package practice

import "strings"

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
// 双指针两端夹逼，遇到不是字母数字的字符时指针多移动一步，然后就是普通的判断是否回文字串
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func isPalindrome(s string) bool {
	// 是否是字母数字
	isAlphanum := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}

	// 全部转成小写
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphanum(s[left]) {
			left++
		}
		for left < right && !isAlphanum(s[right]) {
			right--
		}
		// 规避特殊case，双指针移动之后[left:right]区间全是无效字符
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// @lc code=end
