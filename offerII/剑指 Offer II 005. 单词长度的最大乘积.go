package offerII

// 剑指 Offer II 005. 单词长度的最大乘积
// 给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
// 示例 1:
// 输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
// 输出: 16
// 解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
// 示例 2:
// 输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
// 输出: 4
// 解释: 这两个单词为 "ab", "cd"。
// 示例 3:
// 输入: words = ["a","aa","aaa","aaaa"]
// 输出: 0
// 解释: 不存在这样的两个单词。

// 方法一：暴力枚举
// 时间复杂度：O(n^2 * m)，m表示单词的平均长度，n表示单词的个数
// 空间复杂度：O(1)

// 方法二：位运算 + 预计算
// 时间复杂度：O(m*n + n^2)
// 空间复杂度：O(n)
func maxProduct(words []string) int {
	n := len(words)
	masks := make([]int, n)
	// O(m*n)
	for i := 0; i < n; i++ {
		bitmask := 0 // 位掩码
		for _, c := range words[i] {
			bitmask |= (1 << (c - 'a'))
		}
		masks[i] = bitmask
	}

	// O(n^2)
	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i]&masks[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}
