package practice

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
// 方法一：通过取整和取余操作获取整数中对应的数字进行比较
// 对于1221：
// 1.通过 1221/1000 取得最高位1，通过 1221%1000 取得最低位1，进行比较
// 2.(1221%1000) / 10，先去掉最高位1，然后 221/10 = 22去掉最低位1，继续进行比较
func isPalindromeNum(x int) bool {
	// 边界判断
	if x < 0 {
		return false
	}

	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x > 0 {
		high := x / div // 取最高位
		low := x % 10   // 取最低位
		if high != low {
			return false
		}
		// 分别去掉最高位和最低位，继续上述比较
		x = (x % div) / 10 // x%div：去掉最高位，/10：去掉最低位
		div /= 100         // 去掉了2位，所以是100
	}
	return true // 循环完说明high和low都相等
}

// 方法二：取出后半段数字进行翻转，然后与前半段对比
// 最后根据长度奇偶性，如果是偶数的话，revertNum 和 x 相等；
// 如果是奇数的话，最中间的数字就在 revertNum 的最低位上，将它除以 10 以后应该和 x 相等。
func isPalindromeNum1(x int) bool {
	// 边界判断：x<0 或最低位是0，肯定不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x /= 10
	}
	return x == revertNum || x == revertNum/10
}

// @lc code=end
