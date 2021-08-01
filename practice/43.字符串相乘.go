package practice

import "strconv"

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n, res := len(num1), len(num2), "0"
	// num2 逐位与nums1 相乘
	for i := n - 1; i >= 0; i-- {
		// 保存 num2 第i位数字与 num1 相乘的结果
		var tmp string
		// 从后往前补 n-1-i 个0
		for j := 0; j < n-1-i; j++ {
			tmp += "0"
		}

		n2, carry := int(num2[i]-'0'), 0 // 表示进位
		// num2第i位的数字 n2 与 num1相乘
		for j := m - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			product := n1*n2 + carry
			product, carry = product%10, product/10
			tmp = strconv.Itoa(product) + tmp
		}
		// 将当前结果与新计算的结果求和作为新的结果
		res = addTwoStrings(res, tmp)
	}
	return res
}

func addTwoStrings(num1 string, num2 string) string {
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
