package week10

// 541. 反转字符串 II
// 给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
// 示例:
// 输入: s = "abcdefg", k = 2
// 输出: "bacdfeg"
// @lc: https://leetcode-cn.com/problems/reverse-string-ii/

// 双指针，每2k个分为一段，循环步长为2k
func reverseStr(s string, k int) string {
	n, sb := len(s), []byte(s)      // 字节数组
	for i := 0; i < n; i += 2 * k { // 每2k个为一段
		begin, end := i, min(i+k-1, n-1) // 需考虑剩余字符不足k个的情况
		for begin < end {                // 翻转 [begin, end] 区间内的字符
			sb[begin], sb[end] = sb[end], sb[begin]
			begin++
			end--
		}
	}
	return string(sb)
}
