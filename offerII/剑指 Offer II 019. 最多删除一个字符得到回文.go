package offerII

// 剑指 Offer II 019. 最多删除一个字符得到回文
// 给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
// 示例 1:
// 输入: s = "aba"
// 输出: true
// 示例 2:
// 输入: s = "abca"
// 输出: true
// 解释: 可以删除 "c" 字符 或者 "b" 字符
// 示例 3:
// 输入: s = "abc"
// 输出: false
// @lc: https://leetcode-cn.com/problems/RQku0D/

// 双指针，两端加逼
// 思路：双指针left->0..n-1，right->n-1..0，如果s[left]=s[right]则说明两端元素相同，双指针继续移动，
// 当不相等时，由于s[0, left)和s(right, n-1]已经验证是回文，因此只需要验证s[left, right]是否为回文即可，
// 此时的做法是：删除left或删除right，判断s(left, right]或s[left, right) 为回文即可。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func validPalindrome(s string) bool {
	n := len(s)
	if n < 2 {
		return true
	}

	// isPalindrome 验证s[left, right]是否是回文字符串
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

	for left, right := 0, n-1; left < right; {
		if s[left] == s[right] { // 1. 左右指针的元素相同时，继续两端夹逼
			left++
			right--
		} else { // 2. 左右指针元素不同时，判断 s(left, right] 或 s[left, right) 是否是回文串
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}
	return true
}
