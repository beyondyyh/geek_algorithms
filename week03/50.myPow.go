package week03

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
// 示例1：
// 输入：x = 2.00000, n = 10 输出：1024.00000
// 示例2：
// 输入：x = 2.00000, n = -2 输出：0.25000	解释：2^-2 = 1/2^2 = 1/4 = 0.25
// @leetcode: https://leetcode-cn.com/problems/powx-n

// 方法一：暴力枚举
// 初始化res=1 for i:0->n: res *= x
// 时间复杂度：O(n)
func myPow1(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		return 1 / myPow1(x, -n)
	}

	for i := 0; i < n; i++ {
		res *= x
	}
	return res
}

// 方法二：分治、递归，一分为二
// 分治拆解过程：
// x^n --> 2^10: 2^5 --> 2^2 * 2
// pow(x, n):
// 		subproblem: pow(x, n/2)
// merge:
// 		if n%2==1奇数: result=subresult*subresult*x else偶数: result=subresult*subresult
// 时间复杂度：O(logN)  由于每次递归都会使得指数减少一半，因此递归的层数为log(n)
func myPow2(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow2(x, -n)
	}
	return fastPow(x, n)
}

// fastPow 分治递归函数
func fastPow(x float64, n int) float64 {
	// teiminator 终结条件
	if n == 0 {
		return 1
	}

	// drill down & current logic，下探
	half := fastPow(x, n/2) // n/2会舍弃小数，相当于向下取整

	// merge
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
