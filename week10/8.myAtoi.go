package week10

import "math"

// 8. 字符串转换整数 (atoi)
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
// 函数 myAtoi(string s) 的算法如下：
// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// 返回整数作为最终结果。
// @lc: https://leetcode-cn.com/problems/string-to-integer-atoi/

// 整体思路：
// 1、去掉前导空格
// 2、处理正负号
// 3、识别数字，注意越界情况。
// ans = ans * 10 + digit
func myAtoi(s string) int {
	n, index := len(s), 0
	// 1、去掉前面空格
	for index < n && s[index] == ' ' {
		index++
	}
	// 如果已经遍历完成，针对极端case 如"    "
	if index == n {
		return 0
	}

	// 2、如果出现符号字符，仅第 1 个有效，并记录正负，flag=1 代表正，默认正数
	flag := 1
	if s[index] == '-' {
		index++
		flag = -1
	} else if s[index] == '+' {
		index++
	}

	res := 0
	for index < n {
		digit := int(s[index] - '0')
		if digit < 0 || digit > 9 {
			break
		}
		// 本来应该是 res * 10 + digit > Integer.MAX_VALUE
		// 但是 *10 和 + digit 都有可能越界，所有都移动到右边去就可以了。
		if res > (math.MaxInt32-digit)/10 {
			if flag == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + digit
		index++
	}
	return flag * res
}
