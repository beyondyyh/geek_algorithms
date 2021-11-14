package offerII

import "strings"

// 剑指 Offer II 018. 有效的回文
// 给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
// 本题中，将空字符串定义为有效的 回文串 。
// 示例 1:
// 输入: s = "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串
// 示例 2:
// 输入: s = "race a car"
// 输出: false
// 解释："raceacar" 不是回文串

// 先将字符串转为小写，双指针两端加逼，只要不是数字字母就继续移动，直到相遇
// 时间复杂度：O(n)
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	// 是否为数字、字母
	isAlanum := func(c byte) bool {
		return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlanum(s[left]) {
			left++
		}
		for left < right && !isAlanum(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
