package week10

import "strings"

// @lc code=start

// 方法一：先把遍历一次字符串s把字母数字放入到 sgood 字符串中，然后利用双指针两端夹逼判断是否是回文
// 时间复杂度：O(n) 2次遍历，忽略常数
// 空间复杂度：O(n)

// 方法二：双指针两端夹逼，遇到不是字母数字的字符时指针多移动一步，然后就是普通的判断是否回文字串
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
