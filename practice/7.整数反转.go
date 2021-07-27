package practice

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
// // 弹出 x 的末尾数字 digit
// digit = x % 10
// x /= 10
// // 将数字 digit 推入 rev 末尾
// rev = rev * 10 + digit
func reverse(x int) int {
	var res int
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10 // 弹出x末尾数字，去掉最低位
		x /= 10         // 舍弃末尾数字，向高位进1
		res = res*10 + digit
	}
	return res
}

// @lc code=end
