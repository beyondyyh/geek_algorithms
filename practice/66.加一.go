package practice

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
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

// @lc code=end
