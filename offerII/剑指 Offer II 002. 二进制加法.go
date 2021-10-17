package offerII

import "strconv"

// 剑指 Offer II 002. 二进制加法
// 给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
// 输入为 非空 字符串且只包含数字 1 和 0。

// 示例 1:
// 输入: a = "11", b = "10"
// 输出: "101"

// 示例 2:
// 输入: a = "1010", b = "1011"
// 输出: "10101"

// 提示：
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。

func addBinary(a string, b string) string {
	bits := []int{}              // 表示从低到高位的值
	carry := 0                   // 表示进位
	l1, l2 := len(a)-1, len(b)-1 // a,b的长度
	for l1 >= 0 || l2 >= 0 {
		x, y := 0, 0
		if l1 >= 0 {
			x = int(a[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			y = int(b[l2] - '0')
			l2--
		}

		sum := x + y + carry      // 两数之和，再加上进位
		sum, carry = sum%2, sum/2 // 余数是该位上的值，商表示是否需要进位
		bits = append(bits, sum)  // 余数是该位上的值
	}
	if carry == 1 {
		bits = append(bits, carry)
	}

	// reverse
	res := ""
	for i := len(bits) - 1; i >= 0; i-- {
		res += strconv.Itoa(bits[i])
	}
	return res
}
