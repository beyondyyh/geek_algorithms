package week02

import "sort"

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 异位词：「两个字符串中字符出现的种类和次数均相等」
// 示例 1:
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// @leetcode https://leetcode-cn.com/problems/valid-anagram

// 方法一：暴力法，对2个字符串排序后比较是否相等即可，简单暴力
// 时间复杂度：O(NlogN)
// 空间复杂度：O(logn) 快排空间复杂度
func isAnagram1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 方法二：hashmap
// tt 是 ss 的异位词等价于「两个字符串中字符出现的种类和次数均相等」,
// 由于字符串只包含26个小写字母，因此我们可以维护一个长度为26的频次数组table，分别遍历s和t最后比较table是否相等
// 时间复杂度：O(n)
// 空间复杂度：O(26)
func isAnagram2(s, t string) bool {
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

// 方法三：方法二的进阶
// 没必要全部遍历字符串t，先遍历记录字符串s中字符出现的频次，然后遍历字符串t，减去table中对应的频次，
// 如果出现 table[i]<0，则说明 t 包含一个不在于 s 中的额外字符，直接返回 false 即可
// 时间复杂度：O(n)
// 空间复杂度：O(26)
func isAnagram3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{}
	for _, c := range s {
		cnt[c]++
	}
	for _, c := range t {
		cnt[c]--
		if cnt[c] < 0 {
			return false
		}
	}
	return true
}
