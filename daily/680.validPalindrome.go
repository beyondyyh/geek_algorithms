package daily

// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
// 示例 1:
// 输入: "aba"
// 输出: True
// @lc: https://leetcode-cn.com/problems/valid-palindrome-ii/

// 方法一：双指针，两端夹逼
// 1. left->0..n-1, right->n-1..0，左右指针遇到的元素相等，继续向中间走；
// 2. 左右指针遇到元素不相等时，由于s[0, left) 和 s(right, n-1] 已经判断过是回文串了，
// 所以只需要判断 s[left, right] 区间是否是回文串即可，此时的做法是：删除left 或 删除right 元素，判断 s(left, right] 或 s[left, right) 是否是回文字串即可
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func validPalindrome(s string) bool {
	n := len(s)
	if n < 2 {
		return true
	}

	// isPalindrome 验证s[begin, end]是否是回文字符串
	var isPalindrome func(string, int, int) bool
	isPalindrome = func(s string, begin, end int) bool {
		for begin < end {
			if s[begin] != s[end] {
				return false
			}
			begin++
			end--
		}
		return true
	}

	left, right := 0, n-1
	for left < right {
		if s[left] == s[right] { // 1. 左右指针的元素相同时，继续两端夹逼
			left++
			right--
		} else { // 2. 左右指针元素不同时，判断 s(left, right] 或 s[left, right) 是否是回文串
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}
	return true
}
