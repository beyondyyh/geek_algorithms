package week08

// 191. 位1的个数
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
// @lc: https://leetcode-cn.com/problems/number-of-1-bits/

// 方法一：右移最大32次，每次右移后判断二进制最低位是否为1
// 1. 使用 n&1 得到二进制最低位是否为1，2. 把n右移一位，直到 num!=0
// 时间复杂度：O(n), n为num的二进制位总长度
// 空间复杂度：O(1)
func hammingWeight1(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 == 1 { // num&1=1:最后一位为1，num&1=0:最后一位为0
			res++
		}
		num >>= 1
	}
	return res
}

// 方法二：消除二进制最低位的1，直到 num!=0
// 时间复杂度：O(k), k为num的二进制位中1的个数
// 空间复杂度：O(1)
func hammingWeight2(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1 // n & (n-1) 会清零二进制位最低位的1
	}
	return res
}
