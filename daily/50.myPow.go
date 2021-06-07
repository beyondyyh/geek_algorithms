package daily

// 50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
// 示例 1：
// 输入：x = 2.00000, n = 10
// 输出：1024.00000
// 示例 2：
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25
// @lc: https://leetcode-cn.com/problems/powx-n/

// 方法一：分治+二分法
// 分治拆解过程：
// x^n --> 2^10: 2^5 --> 2^2 * 2
// pow(x, n):
// 		subproblem: pow(x, n/2)
// merge:
// 		if n%2==1奇数: result=subresult*subresult*x else偶数: result=subresult*subresult
// 时间复杂度：O(logN)  由于每次递归都会使得指数减少一半，因此递归的层数为log(n)
// 空间复杂度：O(logN)	递归的层数，递归函数调用使用的栈空间
func myPow(x float64, n int) float64 {
	var fastPow func(x float64, n int) float64
	fastPow = func(x float64, n int) float64 {
		// terminator
		if n == 0 {
			return 1
		}
		// current logic & drill down
		half := fastPow(x, n/2) // n/2下取整
		if n&1 == 0 {           // 偶数
			return half * half
		}
		return half * half * x // 奇数
	}

	if n >= 0 {
		return fastPow(x, n)
	}
	return 1.0 / fastPow(x, -n)
}

// 方法二：位运算，快速幂的迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func myPow2(x float64, n int) float64 {
	if x == 0 { // 0的任何次幂都是0
		return 0
	}

	var res float64 = 1
	if n < 0 { // ans = (1/x)^(-n)
		x, n = 1/x, -n
	}
	for n > 0 {
		if n&1 == 1 { // n的二进制位最低位是1，奇数，即 n%2==1
			res *= x
		}
		x *= x  // x的平方
		n >>= 1 // n右移一位，去掉二进制最低位，等价于n/2下取整
	}
	return res
}
