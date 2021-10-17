package offerII

import "math"

// 剑指 Offer II 001. 整数除法
// 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
// 注意：
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31, 2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1

// 示例 1：
// 输入：a = 15, b = 2
// 输出：7
// 解释：15/2 = truncate(7.5) = 7

// 示例 2：
// 输入：a = 7, b = -3
// 输出：-2
// 解释：7/-3 = truncate(-2.33333..) = -2

// 方法一：使用减法代替除法，用被除数不断减去除数，直到剩下的数据小于除数即可
// 这样的前提是：a和b都是正数或都是负数

// 将 -2147483648 转成正数会越界，但是将 2147483647 转成负数，则不会
// 所以，我们将 a 和 b 都转成负数
// 时间复杂度：O(n)，n 是最大值 2147483647 --> 10^10
// 空间复杂度：O(1)
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}

	res := 0
	for a <= b {
		a -= b
		res++
	}
	return sign * res
}

// 时间复杂度：O(1)
func divide1(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	a, b = abs(a), abs(b)

	res := 0
	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
