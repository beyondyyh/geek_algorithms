package offer

// 剑指 Offer 16. 数值的整数次方
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

// 示例 1：
// 输入：x = 2.00000, n = 10
// 输出：1024.00000

// 示例 2：
// 输入：x = 2.10000, n = 3
// 输出：9.26100

// 示例 3：
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25
// @lc: https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/

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
	var fastPow func(float64, int) float64
	fastPow = func(x float64, n int) float64 {
		// terminator, 任何数的0次幂等于1
		if n == 0 {
			return 1
		}
		half := fastPow(x, n/2) // 分治，一分为二
		if n&1 == 0 {           // even，偶数
			return half * half
		}
		return half * half * x // odd，奇数
	}

	if n >= 0 {
		return fastPow(x, n)
	}
	return 1.0 / fastPow(x, -n)
}
