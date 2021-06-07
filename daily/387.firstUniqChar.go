package daily

// 387. 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
// 示例：
// s = "leetcode"
// 返回 0
// s = "loveleetcode"
// 返回 2
// @lc: https://leetcode-cn.com/problems/first-unique-character-in-a-string/

// 方法一：hashmap，第一次遍历统计字符出现的次数，第二次遍历s看map中计数为1的字符，返回索引
// 时间复杂度：O(n)
// 空间复杂度：O(k) k最大为26
func firstUniqChar1(s string) int {
	counter := [26]int{}
	for _, char := range s {
		counter[char-'a']++
	}
	for i, char := range s {
		if counter[char-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	return -1
}
