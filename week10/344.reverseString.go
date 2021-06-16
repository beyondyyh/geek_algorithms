package week10

// 344. 反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
// 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
// 示例 1：
// 输入：["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
// @lc: https://leetcode-cn.com/problems/reverse-string/

// 双指针搞起来
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseString(s []byte) {
	for begin, end := 0, len(s)-1; begin < end; begin, end = begin+1, end-1 {
		s[begin], s[end] = s[end], s[begin]
	}
}
