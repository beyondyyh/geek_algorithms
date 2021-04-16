package week02

// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。
// 示例:
// 输入: 38
// 输出: 2
// 解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
// @leetcode: https://leetcode-cn.com/problems/add-digits
func addDigits(num int) int {
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
