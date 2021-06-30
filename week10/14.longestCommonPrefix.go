package week10

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。
// @lc: https://leetcode-cn.com/problems/longest-common-prefix/

// 方法一：hasPrefix
// 时间复杂度：O(s) s 为所有字符串的长度之和
// 空间复杂度：O(1)
func longestCommonPrefix1(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	hasPrefix := func(s, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !hasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1] // 去掉最后一个字符，进行比较
		}
	}
	return prefix
}

// 纵向扫描
// 把所有字符串左边对齐，依次比较列上面的字符是否相同
// 时间复杂度：O(n*k) n是字符串个数，k是最短的字符串长度
// 时间复杂度：O(1)
func longestCommonPrefix2(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	prefix := strs[0]
	for i, char := range []byte(prefix) {
		for j := 1; j < n; j++ {
			// 当前 j 对应的字符串遍历完了，或者当前列上的字符不相同，则返回前面已经遍历过的相同的字符
			if i == len(strs[j]) || strs[j][i] != char {
				return prefix[:i]
			}
		}
	}
	return prefix
}
