package practice

import "strconv"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
// 双指针从后往前遍历，需要考虑进位问题
func addStrings(num1 string, num2 string) string {
	var res string
	i, j, carry := len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		res = strconv.Itoa(sum) + res
	}
	// 考虑最后进位的情况
	if carry > 0 {
		res = "1" + res
	}
	return res
}

// @lc code=end
