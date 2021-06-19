package week01

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
// 示例 1：
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
// 解释：输入数组表示数字 123。
// @leetcode: https://leetcode-cn.com/problems/plus-one

// 当前位上的数只能是 9 或小于9
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++     // +1
		digits[i] %= 10 // <9的数求余10还是本身
		if digits[i] != 0 {
			return digits
		}
	}

	// 最高位可能进位，ex. 99
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
