package week10

/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
// 思路：双指针两端夹逼，left->0..n-1, right->n-1..0 直到相遇，
// 如果 s[left] == s[right] 双指针同时移动
// 如果 s[left] != s[right] 此时s[0:left-1] 和 s[n-1:right-1] 由于已经验证过是回文串了，此时去掉一个字符可以是left或right，因此只需要看：
// s[left+1:right] 或 s[left:right-1] 是否是回文串即可
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func validPalindrome(s string) bool {
	// 验证s[left:right]区间是否是回文串
	isPalindrome := func(s string, left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}
	return true
}

// @lc code=end
