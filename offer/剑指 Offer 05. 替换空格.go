package offer

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// 示例 1：
// 输入：s = "We are happy."
// 输出："We%20are%20happy."
// @lc: https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

// golang 字符串不可变长度
// 需要计算新字符串的长度，len(s) + 2*(空格个数)
// 时间复杂度：O(n)  2层遍历
// 时间复杂度：O(n)
func replaceSpace(s string) string {
	count, n := 0, len(s)
	// count spaces
	for _, c := range s {
		if c == ' ' {
			count++
		}
	}

	sb := make([]byte, 0, n+2*count)
	for _, c := range []byte(s) {
		if c == ' ' {
			sb = append(sb, []byte{'%', '2', '0'}...)
		} else {
			sb = append(sb, c)
		}
	}
	return string(sb)
}

// 原地修改，空间复杂度：O(1)
