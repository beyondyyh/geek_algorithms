package week10

import (
	"regexp"
	"strings"
)

// 151. 翻转字符串里的单词
// 给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。
// 示例 1：
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// @lc: https://leetcode-cn.com/problems/reverse-words-in-a-string/

// 方法一：标准库函数
// 思路：1、先去除首位空格，2、通过空格split，3、reverse单词，4、join
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reverseWords(s string) string {
	words := regexp.MustCompile(`\s+`).Split(strings.Trim(s, " "), -1)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// 纯手写
// 思路：
//	1、先去除前后空格，且通过空格切分成字节数组，
//	2、整体reverse，
// 	3、reverse each word
func reverseWords2(s string) string {
	// 1、trim 两端空格，且通过空格切分成字节数组
	trimSpaces := func(s string) []byte {
		left, right := 0, len(s)-1
		for left <= right && s[left] == ' ' {
			left++
		}
		for left <= right && s[right] == ' ' {
			right--
		}

		sb := make([]byte, 0, right-left+1)
		// 将中间多余的空格去掉
		for left <= right {
			c := s[left]
			if c != ' ' {
				sb = append(sb, c)
			} else if sb[len(sb)-1] != ' ' {
				sb = append(sb, c)
			}
			left++
		}
		return sb
	}

	// 2、反转字符串
	reverse := func(sb []byte, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			sb[left], sb[right] = sb[right], sb[left]
		}
	}

	// 3、反转每个单词
	reverseEachWord := func(sb []byte) {
		start, end, n := 0, 0, len(sb)
		for start < n {
			for end < n && sb[end] != ' ' {
				end++
			}
			// 至此，[start,end)是一个完整的单词，反转它
			// 注意：end此时指向的是空格的位置，所以这里是end-1
			reverse(sb, start, end-1)
			// 更新 start、end 去找下一个单词
			start, end = end+1, end+1
		}
	}

	sb := trimSpaces(s)
	reverse(sb, 0, len(sb)-1)
	reverseEachWord(sb)

	return string(sb)
}
