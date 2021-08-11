package offer

// 剑指 Offer 58 - II. 左旋转字符串
// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
// 示例 1：
// 输入: s = "abcdefg", k = 2
// 输出: "cdefgab"
// @lc: https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

func reverseLeftWords(s string, n int) string {
	size := len(s)
	if size == 0 || n <= 0 {
		return s
	}
	n %= size
	return s[n:] + s[:n]
}
