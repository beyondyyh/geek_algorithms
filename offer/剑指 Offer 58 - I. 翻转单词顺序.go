package offer

// 剑指 Offer 58 - I. 翻转单词顺序
// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

// 示例 1：
// 输入: "the sky is blue"
// 输出: "blue is sky the"

// 示例 2：
// 输入: "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// @lc: https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

// 思路：去掉多余空格、整体翻转、翻转每个单词
func reverseWords(s string) string {
	// 1. 双指针去掉多余空格，同时转化为字节数组
	trimSpaces := func(s string) []byte {
		left, right := 0, len(s)-1
		// 去掉两端空格
		for left <= right && s[left] == ' ' {
			left++
		}
		for left <= right && s[right] == ' ' {
			right--
		}
		// 去掉中间多余空格
		sb := make([]byte, 0, right-left+1)
		for left <= right {
			if s[left] != ' ' {
				sb = append(sb, s[left]) // 不为空格则放入sb
			} else if sb[len(sb)-1] != ' ' { // sb最后一个字符不为空格则放入，此处保证了单词直接只保留1个空格
				sb = append(sb, s[left])
			}
			left++
		}
		return sb
	}

	// 2. 翻转区间 [left, right] 左闭右闭区间
	reverse := func(sb []byte, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			sb[left], sb[right] = sb[right], sb[left]
		}
	}

	// 3. 翻转每个单词
	reverseEachWord := func(sb []byte) {
		// start,end：单词的起始位置，n：字符串长度
		start, end, n := 0, 0, len(sb)
		for start < n {
			for end < n && sb[end] != ' ' {
				end++
			}
			// 此时 sb[start, end) 是一个单词，翻转之，注意：此时end指向的是空格的位置
			reverse(sb, start, end-1)
			// 更新 start,end 去找下一个单词的起始位置
			start, end = end+1, end+1
		}
	}

	// main process
	sb := trimSpaces(s)
	reverse(sb, 0, len(sb)-1)
	reverseEachWord(sb)
	return string(sb)
}
