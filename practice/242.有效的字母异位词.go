package practice

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
// hashmap计数，然后判断是否相同
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var c1, c2 [26]int
	for _, c := range s {
		c1[c-'a']++
	}
	for _, c := range t {
		c2[c-'a']++
	}
	return c1 == c2
}

// @lc code=end
