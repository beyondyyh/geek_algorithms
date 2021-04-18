package week03

import "strings"

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// 示例 1：
// 输入：s = "We are happy."
// @leetcode: https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof

// 方法一：直接使用库函数
// 不过真正面试时应该不允许
func replaceSpace1(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

// 方法二：遍历
// 需要计算新字符串的长度，len(s) + 2*(空格个数)
// 时间复杂度：O(n)  2层遍历，2O(n)，去掉常量
// 时间复杂度：O(n)
func replaceSpace2(s string) string {
	count, size := 0, len(s)
	// count spaces
	for _, char := range s {
		if char == ' ' {
			count++
		}
	}
	size += 2 * count // 新字符串长度

	news := make([]rune, 0, size)
	for _, char := range s {
		if char == ' ' {
			news = append(news, []rune{'%', '2', '0'}...)
		} else {
			news = append(news, char)
		}
	}
	return string(news)
}
