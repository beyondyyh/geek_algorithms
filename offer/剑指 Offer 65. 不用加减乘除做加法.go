package offer

// 剑指 Offer 65. 不用加减乘除做加法
// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
// 示例:
// 输入: a = 1, b = 1
// 输出: 2
// @lc: https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

// 无进位和 与 异或运算 规律相同，进位 和 与运算 规律相同（并需左移一位）
// 时间复杂度：O(1)，最差情况 a=0x7fffffff, b=1, 需循环 32 次，使用O(1)时间；
// 空间复杂度：O(1)
func add(a int, b int) int {
	for b != 0 { // 当进位为 0 时跳出
		c := (a & b) << 1 // c = 进位
		a ^= b            // a = 非进位和
		b = c             // b = 进位
	}
	return a
}
