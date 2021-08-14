package offer

// 剑指 Offer 50. 第一个只出现一次的字符
// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
// 示例:
// s = "abaccdeff"
// 返回 "b"
// s = ""
// 返回 " "
// @lc: https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

// 先遍历一次s用将(char, 出现的次数count)存入hashmap
// 再遍历一次s依次判断char的count
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	for _, c := range []byte(s) {
		m[c]++
	}
	for _, c := range []byte(s) {
		if m[c] == 1 {
			return c
		}
	}
	return ' '
}
