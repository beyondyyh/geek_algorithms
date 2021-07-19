package practice

import "strings"

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
// 如果 s 中没有循环，那么 s+s 中必然有且只有两个 s，此时从 ss[1] 处开始寻找 s ，必然只能找到第二个，所以此时返回值为找到的索引下标，必定为 s.size()
// 如果 s 中有循环时，设循环节为 r，其长度为 l，那么 ss 中必然有 s.size()/l + 1 个 s
func repeatedSubstringPattern(s string) bool {
	newS := string((s + s)[1:])               // 去掉第一个字符
	return strings.Index(newS, s)+1 != len(s) // 去掉了第一个字符，这里要找到的索引位置要加上1
}

// @lc code=end
